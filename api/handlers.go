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
			TTL   int    `json:'ttl'`
		}

		err := json.NewDecoder(r.Body).Decode(&req)

		if err != nil {
			http.Error(w, "Invalid request Body", http.StatusBadRequest)
			return
		}

		kv.SetKeyValue(req.Key, req.Value, req.TTL)
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
