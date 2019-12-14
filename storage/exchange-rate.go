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
		"AUD", "CNY", "EUR", "GBP", "HKD", "JPY", "SGD", "USD",
	}

	defaultData := map[string]string{
		keys[0]: utils.ToJson(models.NewExchangeRate("澳幣", keys[0], 19.6100, "", 1)),
		keys[1]: utils.ToJson(models.NewExchangeRate("人民幣", keys[1], 4.0500, "", 2)),
		keys[2]: utils.ToJson(models.NewExchangeRate("歐元", keys[2], 31.6400, "", 3)),
		keys[3]: utils.ToJson(models.NewExchangeRate("英鎊", keys[3], 36.8900, "", 4)),
		keys[4]: utils.ToJson(models.NewExchangeRate("港幣", keys[4], 3.5900, "", 5)),
		keys[5]: utils.ToJson(models.NewExchangeRate("日圓", keys[5], 0.2600, "", 6)),
		keys[6]: utils.ToJson(models.NewExchangeRate("新加坡幣", keys[6], 20.9100, "", 7)),
		keys[7]: utils.ToJson(models.NewExchangeRate("美金", keys[7], 29.6000, "", 8)),
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

func (b *ExchangeRateBucket) ForEach() (data []*models.ExchangeRate, err error) {
	data = make([]*models.ExchangeRate, 0)
	err = b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(b.name))
		if bucket != nil {
			err := bucket.ForEach(func(k, v []byte) error {
				//fmt.Println(string(k), string(v))

				exRate := &models.ExchangeRate{}
				utils.FromJson(string(v), exRate)

				data = append(data, exRate)
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
		return nil
	})
	return data, err
}

func (b *ExchangeRateBucket) Update(data []*models.ExchangeRate) (err error) {
	err = b.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(b.name))
		if bucket != nil {

			for _, rate := range data {
				if val := bucket.Get([]byte(rate.Currency)); val != nil {
					exRate := models.ExchangeRate{}
					utils.FromJson(string(val), &exRate)

					if rate.Rate > 0 {
						exRate.Rate = rate.Rate
					}

					if rate.Rate > 0 {
						exRate.Rank = rate.Rank
					}

					err := bucket.Put([]byte(rate.Currency), []byte(utils.ToJson(exRate)))
					if err != nil {
						return err
					}
				}
			}
		} else {
			err := errors.New(fmt.Sprint("can not find db: ", b.name))
			return err
		}
		return nil
	})
	return err
}
