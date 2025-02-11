package api

import (
	"encoding/json"
	"net/http"

	"github.com/Yashh56/keyValueStore/internal/store"
)

func HandleSet(kv *store.KeyValueStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Key   string `json:"key"`
			Value string `json:"value"`
			TTL   *int   `json:'ttl'`
		}

		err := json.NewDecoder(r.Body).Decode(&req)

		if err != nil {
			http.Error(w, "Invalid request Body", http.StatusBadRequest)
			return
		}

		ttl := 0
		if req.TTL != nil {
			ttl = *req.TTL
		}

		kv.SetKeyValue(req.Key, req.Value, ttl)
		w.WriteHeader(http.StatusOK)
	}
}

func HandleGet(kv *store.KeyValueStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")

		if key == "" {
			http.Error(w, "Key Parameter is missing", http.StatusBadRequest)
			return
		}

		val, ok := kv.GetKeyValue(key)
		if !ok {
			http.Error(w, "Key not found", http.StatusNotFound)
			return
		}
		resp := struct {
			Key   string `json:"Key"`
			Value string `json:"value"`
		}{Key: key, Value: val}

		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}

func HandleDelete(kv *store.KeyValueStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")

		if key == "" {
			http.Error(w, "Key Parameter is missing", http.StatusBadRequest)
			return
		}
		ok := kv.DeleteKeyValue(key)
		if !ok {
			http.Error(w, "Key not found", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func HandleBatchSet(kv *store.KeyValueStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Items map[string]string `json:"items"` // Fixed JSON tag
			TTL   int               `json:"ttl"`
		}

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		kv.SetBatch(req.Items, req.TTL) // Use req.Items instead of req.items
		w.WriteHeader(http.StatusOK)
	}
}

func HandleBatchGet(kv *store.KeyValueStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Keys []string `json:"keys"` // Fixed JSON tag
		}

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if len(req.Keys) == 0 {
			http.Error(w, "Enter the keys", http.StatusBadRequest)
			return
		}

		val := kv.GetBatch(req.Keys)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(val)
	}
}
