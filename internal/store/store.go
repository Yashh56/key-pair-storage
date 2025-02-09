package store

import (
	"sync"
)

type KeyValueStore struct {
	keyValue   map[string]string
	mu         sync.RWMutex
	ttlManager *TTLManager
}

func NewKeyValueStore() *KeyValueStore {

	store := &KeyValueStore{
		keyValue: make(map[string]string),
	}
	store.ttlManager = NewTTLManager(store)

	return store
}

func (kv *KeyValueStore) SetKeyValue(key, value string, ttlSeconds int) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	kv.keyValue[key] = value
	SaveToDisk(key, value)
	kv.ttlManager.SetTTL(key, ttlSeconds)
}

func (kv *KeyValueStore) GetKeyValue(key string) (string, bool) {
	kv.mu.RLock()
	defer kv.mu.RUnlock()

	if kv.ttlManager.IsExpired(key) {
		kv.mu.RUnlock()
		kv.DeleteKeyValue(key)
		kv.mu.RLock()
		return "", false
	}

	val, ok := LoadFromDisk(key)

	return val, ok

}

func (kv *KeyValueStore) DeleteKeyValue(key string) bool {

	kv.mu.Lock()
	defer kv.mu.Unlock()

	_, ok := kv.keyValue[key]

	if ok {
		delete(kv.keyValue, key)
		DeleteFromDisk(key)
	}

	return ok
}
