package database

import (
	"log"

	"github.com/boltdb/bolt"
)

func OpenRead() *bolt.DB {
	return openDB(true)
}

func OpenWrite() *bolt.DB {
	return openDB(false)
}

func openDB(readOnly bool) *bolt.DB {
	db, err := bolt.Open("diwia.db", 0600, &bolt.Options{ReadOnly: readOnly})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
