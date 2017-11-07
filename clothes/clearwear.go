package clothes

import (
	"sync"

	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
)

func ClearWears(db *bolt.DB, ids []string) error {
	errs := make(chan error, len(ids))

	var wg sync.WaitGroup
	for _, id := range ids {
		wg.Add(1)
		go func(itemID string) {
			errs <- db.Batch(func(tx *bolt.Tx) error {
				return tx.Bucket(bucketWears).Delete([]byte(itemID))
			})
			wg.Done()
		}(id)
	}

	wg.Wait()
	close(errs)

	for err := range errs {
		if err != nil {
			return errors.Wrap(err, "failed to clear wear-counts")
		}
	}

	return nil
}
