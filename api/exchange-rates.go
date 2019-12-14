package api

import (
	"encoding/json"
	"github.com/richguo0615/exchange-rates/models"
	"github.com/richguo0615/exchange-rates/storage"
	"github.com/richguo0615/exchange-rates/utils"
	"log"
	"net/http"
)

func GetExchangeRates(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	data, err := storage.Sto.ExchangeRateBucket.ForEach()
	if err != nil {
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func PutExchangeRates(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	r.ParseForm()

	data := r.FormValue("Data")
	rateList := make([]*models.ExchangeRate, 0)

	utils.FromJson(data, &rateList)

	err := storage.Sto.ExchangeRateBucket.Update(rateList)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
