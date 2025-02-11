package cache

import (
	"container/list"
	"sync"
	"time"
)

type LRUCache struct {
	capacity int
	cache    map[string]*list.Element
	ll       *list.List
	mu       sync.Mutex
}

type entry struct {
	key    string
	value  string
	expiry int64
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[string]*list.Element),
		ll:       list.New(),
	}
}

func (c *LRUCache) Set(key, value string, ttlSeconds int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	expiry := time.Now().Unix() + int64(ttlSeconds)

	if elem, found := c.cache[key]; found {
		elem.Value.(*entry).value = value
		elem.Value.(*entry).expiry = expiry
		c.ll.MoveToFront(elem)
		return
	}

	if c.ll.Len() >= c.capacity {
		lastElem := c.ll.Back()
		if lastElem != nil {
			delete(c.cache, lastElem.Value.(*entry).key)
			c.ll.Remove(lastElem)
		}
	}

	newElem := c.ll.PushFront(&entry{key, value, expiry})
	c.cache[key] = newElem
}

func (c *LRUCache) Get(key string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, found := c.cache[key]; found {

		if time.Now().Unix() > elem.Value.(*entry).expiry {
			delete(c.cache, key)
			c.ll.Remove(elem)
			return "", false
		}
		c.ll.MoveToFront(elem)
		return elem.Value.(*entry).value, true
	}
	return "", false
}

func (c *LRUCache) Delete(key string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if elem, found := c.cache[key]; found {
		delete(c.cache, key)
		c.ll.Remove(elem)
		return true
	}
	return false
}

func (c *LRUCache) StartTTLWorker() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		for key, elem := range c.cache {
			if time.Now().Unix() > elem.Value.(*entry).expiry {
				delete(c.cache, key)
				c.ll.Remove(elem)
			}
		}
		c.mu.Unlock()
	}
}
