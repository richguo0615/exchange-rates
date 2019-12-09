package main

import (
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

	err = sto.ExchangeRateBucket.Save("USD", exJson)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = sto.ExchangeRateBucket.ForEach()
	if err != nil {
		log.Fatal(err)
		return
	}

	defer db.Close()
}
