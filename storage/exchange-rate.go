package storage

import (
	"errors"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/richguo0615/exchange-rates/models"
	"github.com/richguo0615/exchange-rates/utils"
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

func (b *ExchangeRateBucket) SetDefault() error {

	keys := []string{
		"USD", "CNY", "JPY", "SGD", "HKD", "EUR", "GBP",
	}

	defaultData := map[string]string{
		keys[0]: utils.ToJson(models.NewExchangeRate("美元", "USD", 29.6000, "")),
		keys[1]: utils.ToJson(models.NewExchangeRate("人民幣", "CNY", 4.1110, "")),
		keys[2]: utils.ToJson(models.NewExchangeRate("日幣", "JPY", 0.2625, "")),
		keys[3]: utils.ToJson(models.NewExchangeRate("新加坡", "SGD", 21.1070, "")),
		keys[4]: utils.ToJson(models.NewExchangeRate("新加坡", "HKD", 3.6270, "")),
		keys[5]: utils.ToJson(models.NewExchangeRate("新加坡", "EUR", 31.7960, "")),
		keys[6]: utils.ToJson(models.NewExchangeRate("新加坡", "GBP", 36.9660, "")),
	}

	for _, key := range keys {
		err := b.Save(key, defaultData[key])
		if err != nil {
			return err
		}
	}
	return nil
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