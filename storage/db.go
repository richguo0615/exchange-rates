package storage

import (
	"github.com/asdine/storm"
	"github.com/boltdb/bolt"
	"log"
)

//Buckets
const ExchangeRate = "ExchangeRate"

type Storage struct {
	ExchangeRateBucket *ExchangeRateBucket
}

func Init() (*bolt.DB, error) {
	storm.Open("database.db")
	db, err := bolt.Open("database.db", 0600, nil)
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
		ExchangeRateBucket: NewExchangeRateBucket(db),
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
