package storage

import (
	"errors"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

type ExchangeRateBucket struct {
	name string
	db   *bolt.DB
}

func NewExchangeRateBucket(name string, db *bolt.DB) *ExchangeRateBucket {
	return &ExchangeRateBucket{
		name: name,
		db:   db,
	}
}

func (b *ExchangeRateBucket) Save(key string, value string) error {
	err := b.db.Update(func(tx *bolt.Tx) error {

		bucket := tx.Bucket([]byte(b.name))
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
		bucket := tx.Bucket([]byte(b.name))
		if bucket != nil {
			err := bucket.ForEach(func(k, v []byte) error {
				fmt.Println(string(k), string(v))
				return nil
			})
			if err != nil {
				log.Fatal(err)
				return err
			}
		} else {
			err := errors.New(fmt.Sprint("can not find db: ", b.name))
			return err
		}
		return  nil
	})
	return err
}
