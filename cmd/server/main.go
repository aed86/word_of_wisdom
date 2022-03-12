package main

import (
	"log"
	"net/http"

	"github.com/aed86/proof_of_work/internal/server/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/get_challenge", handlers.GetChallenge).Methods("GET")
	r.HandleFunc("/get_quote", handlers.GetQuote).Methods("GET")

	srv := &http.Server{
		Addr:    ":8000",
		Handler: r,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
