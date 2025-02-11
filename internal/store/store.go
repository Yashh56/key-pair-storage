package store

type KeyValueStore struct {
	lru *LRUCache
}

func NewKeyValueStore(maxEntries int) *KeyValueStore {
	cache := NewLRUCache(maxEntries)
	go cache.startTTLWorker()
	return &KeyValueStore{lru: cache}
}

func (kv *KeyValueStore) SetKeyValue(key, value string, ttlSeconds int) {
	kv.lru.Set(key, value, ttlSeconds)
	SaveToDisk(key, value, ttlSeconds)
}

func (kv *KeyValueStore) GetKeyValue(key string) (string, bool) {
	return kv.lru.Get(key)
}

func (kv *KeyValueStore) DeleteKeyValue(key string) bool {
	del1 := kv.lru.Delete(key)
	del2 := DeleteFromDisk(key)
	return del1 == del2
}

func (kv *KeyValueStore) SetBatch(items map[string]string, ttlSeconds int) {
	for key, value := range items {
		kv.lru.Set(key, value, ttlSeconds)
		SaveToDisk(key, value, ttlSeconds)
	}
}

func (kv *KeyValueStore) GetBatch(keys []string) map[string]string {
	results := make(map[string]string)

	for _, key := range keys {
		if val, found := kv.lru.Get(key); found {
			results[key] = val
		}
	}
	return results
}

func (kv *KeyValueStore) DeleteBatch(keys []string) {
	for _, key := range keys {
		kv.lru.Delete(key)
		DeleteFromDisk(key)
	}
}
