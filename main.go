package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/richguo0615/exchange-rates/models"
	"github.com/richguo0615/exchange-rates/storage"
	"github.com/richguo0615/exchange-rates/utils"
	"log"
)

func main() {
	db, err := storage.Init()
	if err != nil {
		log.Fatal(err)
	}

	sto := storage.NewStorage(db)

	ex := models.NewExchangeRate("美元", "USD", 29.6000, "")
	exJson := utils.ToJson(ex)

	err := sto.ExchangeRateBucket.Save("USD", exJson)
	if err != nil {
		log.Fatal(err)
	}

	err = sto.ExchangeRate.ForEach(func(k, v []byte) error {
		fmt.Println("k: ", string(k), ", v: ", string(v))
		return nil
	})

	defer db.Close()
}
