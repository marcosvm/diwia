package clothes

import (
	"fmt"

	"github.com/boltdb/bolt"
)

func AddItem(db *bolt.DB, id, desc string) error {
	return db.Update(func(tx *bolt.Tx) error {
		bk, err := tx.CreateBucketIfNotExists(bucketClothes)
		if err != nil {
			return fmt.Errorf("Failed to create bucket: %v", err)
		}

		if err := bk.Put([]byte(id), []byte(desc)); err != nil {
			return fmt.Errorf("Failed to insert '%s': %v", desc, err)
		}
		return nil
	})
}
