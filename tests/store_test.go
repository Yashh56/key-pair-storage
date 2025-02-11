package tests

import (
	"testing"
	"time"

	"github.com/Yashh56/keyValueStore/internal/store"
)

func TestKeyValueStore_SetGetKeyValue(t *testing.T) {
	store := store.NewKeyValueStore(10)
	store.SetKeyValue("foo", "bar", 0)

	val, found := store.GetKeyValue("foo")
	if !found || val != "bar" {
		t.Errorf("Expected 'bar', got '%s'", val)
	}
}

func TestKeyValueStore_DeleteKeyValue(t *testing.T) {
	store := store.NewKeyValueStore(10)
	store.SetKeyValue("foo", "bar", 0)
	deleted := store.DeleteKeyValue("foo")
	if !deleted {
		t.Errorf("Expected key 'foo' to be deleted")
	}

	_, found := store.GetKeyValue("foo")
	if found {
		t.Errorf("Expected 'foo' to be absent after deletion")
	}
}

func TestKeyValueStore_SetBatch(t *testing.T) {
	store := store.NewKeyValueStore(10)
	items := map[string]string{"k1": "v1", "k2": "v2"}
	store.SetBatch(items, 0)

	for k, v := range items {
		val, found := store.GetKeyValue(k)
		if !found || val != v {
			t.Errorf("Expected %s for key %s, got %s", v, k, val)
		}
	}
}

func TestKeyValueStore_GetBatch(t *testing.T) {
	store := store.NewKeyValueStore(10)
	items := map[string]string{"k1": "v1", "k2": "v2"}
	store.SetBatch(items, 0)

	keys := []string{"k1", "k2", "k3"}
	results := store.GetBatch(keys)

	if len(results) != 2 {
		t.Errorf("Expected 2 keys in batch result, got %d", len(results))
	}
}

func TestKeyValueStore_DeleteBatch(t *testing.T) {
	store := store.NewKeyValueStore(10)
	items := map[string]string{"k1": "v1", "k2": "v2"}
	store.SetBatch(items, 0)
	store.DeleteBatch([]string{"k1", "k2"})

	for k := range items {
		_, found := store.GetKeyValue(k)
		if found {
			t.Errorf("Expected key %s to be deleted", k)
		}
	}
}

func TestKeyValueStore_TTLExpiration(t *testing.T) {
	store := store.NewKeyValueStore(10)
	store.SetKeyValue("temp", "data", 1) // 1 second TTL
	time.Sleep(2 * time.Second)          // Wait for TTL to expire

	_, found := store.GetKeyValue("temp")
	if found {
		t.Errorf("Expected 'temp' key to expire but it still exists")
	}
}
