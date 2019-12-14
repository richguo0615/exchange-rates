package api

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func NewHandler() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/exchangeRates", GetExchangeRates).Methods("GET")

	log.Print("http listen port: 8080")
	http.ListenAndServe(":8080", r)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}