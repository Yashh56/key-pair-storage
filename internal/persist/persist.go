package persist

import (
	"log"
	"time"

	"github.com/dgraph-io/badger/v4"
)

func SaveToDisk(key string, value string, ttlSeconds int) bool {
	opts := badger.DefaultOptions("badgerdb").WithLoggingLevel(badger.ERROR)
	db, err := badger.Open(opts)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry([]byte(key), []byte(value))
		if ttlSeconds > 0 {
			e.WithTTL(time.Duration(ttlSeconds) * time.Second)
		}
		return txn.SetEntry(e)
	})

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true

}

func LoadFromDisk(key string) (string, bool) {

	opts := badger.DefaultOptions("badgerdb").WithLoggingLevel(badger.ERROR)
	db, err := badger.Open(opts)
	value := "null"
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}

		return item.Value(func(val []byte) error {
			value = string(val)
			return nil
		})
	})

	if err != nil {
		log.Fatal(err)
	}

	if value != "" {
		return value, true
	}

	return "No Value Found for this Key.", false
}

func DeleteFromDisk(key string) bool {
	opts := badger.DefaultOptions("badgerdb").WithLoggingLevel(badger.ERROR)
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	err = db.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(key))
	})
	if err != nil {
		log.Fatal(err)
	}
	return true
}
