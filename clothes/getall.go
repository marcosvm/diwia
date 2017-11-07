package clothes

import (
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
)

var bucketClothes = []byte("clothes")

var ErrNoBucket = fmt.Errorf("failed to find bucket")

func GetAll(db *bolt.DB) ([]Piece, error) {
	var pieces []Piece
	err := db.View(func(tx *bolt.Tx) error {
		bk := tx.Bucket([]byte(bucketClothes))
		if bk == nil {
			return errors.Wrapf(ErrNoBucket, "failed to get 'clothes' bucket")
		}

		c := bk.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			pieces = append(pieces, Piece{
				ID:          string(k),
				Description: string(v),
			})
		}

		return nil
	})
	return pieces, err
}
