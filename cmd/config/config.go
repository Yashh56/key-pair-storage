package config

import (
	"log"

	badger "github.com/dgraph-io/badger/v4"
)

func Config() {

	var DB, err = badger.Open(badger.DefaultOptions("/tmp/badger"))

	if err != nil {
		log.Fatal(err)
	}

	defer DB.Close()
}
