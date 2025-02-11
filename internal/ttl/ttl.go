package ttl

import (
	"sync"
	"time"

	"github.com/Yashh56/keyValueStore/internal/store"
)

type TTLManager struct {
	ttlMap map[string]time.Time
	mu     sync.Mutex
	store  *store.KeyValueStore
}

func NewTTLManager(store *store.KeyValueStore) *TTLManager {
	ttlManager := &TTLManager{
		ttlMap: make(map[string]time.Time),
		store:  store,
	}

	go ttlManager.startTTLWorker()
	return ttlManager
}

func (t *TTLManager) SetTTL(key string, ttlSeconds int) {
	if ttlSeconds > 0 {
		t.mu.Lock()
		t.ttlMap[key] = time.Now().Add(time.Duration(ttlSeconds) * time.Second)
		t.mu.Unlock()
	}
}

func (t *TTLManager) IsExpired(key string) bool {
	t.mu.Lock()
	defer t.mu.Unlock()

	expiry, exists := t.ttlMap[key]
	return exists && time.Now().After(expiry)
}

func (t *TTLManager) startTTLWorker() {
	for {
		time.Sleep(5 * time.Second)
		t.mu.Lock()
		now := time.Now()

		for key, expiry := range t.ttlMap {
			if now.After(expiry) {
				delete(t.ttlMap, key)
				t.store.DeleteKeyValue(key)
			}
		}
		t.mu.Unlock()
	}
}
