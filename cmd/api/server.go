package api

import (
	"fmt"
	"net/http"

	"github.com/Yashh56/keyValueStore/cmd/internal/store"
)

func Server() {
	kv := store.NewKeyValueStore()

	http.HandleFunc("/set", HandleSet(kv))
	http.HandleFunc("/get", HandleGet(kv))
	var port = 8080
	address := fmt.Sprintf(":%d", port)
	fmt.Printf("Starting server on localhost%s\n", address)
	err := http.ListenAndServe(address, nil)

	if err != nil {
		fmt.Sprintf("Error %s\n", err)
	}
}
