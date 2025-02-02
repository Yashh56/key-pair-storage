package store

import (
	"sync"
)

type KeyValueStore struct {
	keyValue map[string]string
	mu       sync.RWMutex
}

func NewKeyValueStore() *KeyValueStore {
	return &KeyValueStore{
		keyValue: make(map[string]string),
	}
}

func (kv *KeyValueStore) SetKeyValue(key, value string) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	SaveToDisk(key, value)
	kv.keyValue[key] = value
}

func (kv *KeyValueStore) GetKeyValue(key string) (string, bool) {
	kv.mu.RLock()
	defer kv.mu.RUnlock()

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
