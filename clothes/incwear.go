package clothes

import (
	"fmt"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
)

func IncWears(db *bolt.DB, id string) error {
	err := db.Update(func(tx *bolt.Tx) error {
		bk, err := tx.CreateBucketIfNotExists([]byte(bucketWears))
		if err != nil {
			return errors.Wrapf(ErrNoBucket, "failed to get 'wears' bucket")
		}

		var count int

		bs := bk.Get([]byte(id))
		if bs == nil {
			fmt.Printf("No previous wears found for '%s'\n", id)
		} else {
			count, err = strconv.Atoi(string(bs))
			if err != nil {
				return errors.Wrapf(ErrWearNaN, "invalid wear value for '%s'", id)
			}
		}

		count = count + 1
		countStr := strconv.Itoa(count)

		return bk.Put([]byte(id), []byte(countStr))
	})
	return err
}
