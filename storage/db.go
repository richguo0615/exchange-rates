package storage

import (
	"github.com/boltdb/bolt"
	"log"
)

//Buckets
const ExchangeRate = "ExchangeRate"

var Sto *Storage

type Storage struct {
	ExchangeRateBucket *ExchangeRateBucket
}

func Init() (*bolt.DB, error) {
	db, err := bolt.Open("storage/database.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}

func NewStorage(db *bolt.DB) *Storage {

	updateErr := db.Update(func(tx *bolt.Tx) error {

		if err := CreateBucket(tx, ExchangeRate); err != nil {
			log.Fatal(err)
		}

		return nil
	})

	if updateErr != nil {
		log.Fatal(updateErr)
	}

	s := &Storage{
		ExchangeRateBucket: NewExchangeRateBucket(ExchangeRate, db),
	}

	return s
}

func CreateBucket(tx *bolt.Tx, bucket string) error {
	if exBk := tx.Bucket([]byte(bucket)); exBk == nil {
		_, err := tx.CreateBucket([]byte(bucket))
		if err != nil {
			return err
		}
	}
	return nil
}
