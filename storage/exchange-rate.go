package storage

import (
	"github.com/boltdb/bolt"
)

type ExchangeRateBucket struct {
	name string
	db   *bolt.DB
}

func NewExchangeRateBucket(name string, db *bolt.DB) *ExchangeRateBucket {
	return &ExchangeRateBucket{
		name: "",
		db:   db,
	}
}

func (b *ExchangeRateBucket) Save(key string, value string) error {
	err := b.db.Update(func(tx *bolt.Tx) error {

		bucket := tx.Bucket([]byte(ExchangeRate))
		if bucket != nil {
			err := bucket.Put([]byte(key), []byte(value))
			if err != nil {
				return err
			}
		}

		return nil
	})
	return err
}

func (b *ExchangeRateBucket) ForEach() error {
	err := b.db.View(func(tx *bolt.Tx) error {
		tx.Bucket([]byte(ExchangeRate))
	})
	return err
}
