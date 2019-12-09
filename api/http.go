package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewHandler() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/exchangeRates", GetExchangeRates).Methods("GET")

	http.ListenAndServe(":8080", r)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}