package main

import (
	"log"
	"net/http"

	"ShoppingCalculator/api"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", api.HealthHandler)
	mux.HandleFunc("POST /api/checkout", api.CheckoutHandler)
	mux.HandleFunc("GET /api/suggestions", api.SuggestionsHandler)

	fileServer := http.FileServer(http.Dir("web"))
	mux.Handle("/", fileServer)

	addr := ":8080"
	log.Printf("server listening on http://localhost%s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
