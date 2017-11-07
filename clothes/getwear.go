package clothes

import (
	"fmt"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
)

var bucketWears = []byte("wears")

var ErrNotFound = fmt.Errorf("key not found")
var ErrWearNaN = fmt.Errorf("wear count is not a number")

func GetWears(db *bolt.DB, id string) (int, error) {
	var count int
	err := db.View(func(tx *bolt.Tx) error {
		bk := tx.Bucket([]byte(bucketWears))
		if bk == nil {
			return errors.Wrapf(ErrNoBucket, "failed to get 'wears' bucket")
		}

		bs := bk.Get([]byte(id))
		if bs == nil {
			return errors.Wrapf(ErrNotFound, "failed to find wears for 's'", id)
		}

		var err error
		count, err = strconv.Atoi(string(bs))
		if err != nil {
			return errors.Wrapf(ErrWearNaN, "invalid wear value for '%s'", id)
		}

		return nil
	})
	return count, err
}
