package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Yashh56/keyValueStore/api"
	"github.com/Yashh56/keyValueStore/internal/store"
)

func TestHandleSet(t *testing.T) {
	kv := store.NewKeyValueStore(100)
	handler := api.HandleSet(kv)
	reqBody := []byte(`{"key":"testKey","value":"testValue","ttl":60}`)
	req := httptest.NewRequest("POST", "/set", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}
}

func TestHandleGet(t *testing.T) {
	kv := store.NewKeyValueStore(100)
	kv.SetKeyValue("testKey", "testValue", 60)
	handler := api.HandleGet(kv)

	req := httptest.NewRequest("GET", "/get?key=testKey", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	var resp struct {
		Key   string `json:"Key"`
		Value string `json:"value"`
	}
	json.NewDecoder(w.Body).Decode(&resp)
	if resp.Value != "testValue" {
		t.Errorf("Expected value %s, got %s", "testValue", resp.Value)
	}
}

func TestHandleBatchSet(t *testing.T) {
	kv := store.NewKeyValueStore(100)
	handler := api.HandleBatchSet(kv)
	reqBody := []byte(`{"items":{"key1":"value1","key2":"value2"},"ttl":60}`)
	req := httptest.NewRequest("POST", "/batchset", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}
}

func TestHandleBatchGet(t *testing.T) {
	kv := store.NewKeyValueStore(100)
	kv.SetKeyValue("key1", "value1", 60)
	kv.SetKeyValue("key2", "value2", 60)
	handler := api.HandleBatchGet(kv)

	reqBody := []byte(`{"keys":["key1","key2"]}`)
	req := httptest.NewRequest("POST", "/batchget", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	var resp map[string]string
	json.NewDecoder(w.Body).Decode(&resp)

	if resp["key1"] != "value1" || resp["key2"] != "value2" {
		t.Errorf("Unexpected response: %v", resp)
	}
}
