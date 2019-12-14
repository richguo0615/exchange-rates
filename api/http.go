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
	api.HandleFunc("/exchangeRates", PutExchangeRates).Methods("POST")

	log.Print("http listen port: 8080")
	http.ListenAndServe(":8080", r)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}