package main

import (
	"github.com/richguo0615/exchange-rates/storage"
	"log"
)

func main() {
	db, err := storage.Init()
	if err != nil {
		log.Fatal(err)
	}

	sto := storage.NewStorage(db)

	err = sto.ExchangeRateBucket.SetDefault()
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
