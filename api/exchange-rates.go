package api

import (
	"encoding/json"
	"github.com/richguo0615/exchange-rates/storage"
	"log"
	"net/http"
)

func GetExchangeRates(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	data, err := storage.Sto.ExchangeRateBucket.ForEach()
	if err != nil {
		log.Fatal(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
