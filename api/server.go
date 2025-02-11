package api

import (
	"fmt"
	"net/http"

	"github.com/Yashh56/keyValueStore/internal/store"
)

func Server() {
	kv := store.NewKeyValueStore(100)

	http.HandleFunc("/set", HandleSet(kv))
	http.HandleFunc("/get", HandleGet(kv))
	http.HandleFunc("/batchSet", HandleBatchSet(kv))
	http.HandleFunc("/batchGet", HandleBatchGet(kv))
	var port = 8080
	address := fmt.Sprintf(":%d", port)
	fmt.Printf("Starting server on localhost%s\n", address)
	err := http.ListenAndServe(address, nil)

	if err != nil {
		fmt.Sprintf("Error %s\n", err)
	}
}
