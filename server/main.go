package main

import (
	"log"
	"net/http"
	"os"

	"ShoppingCalculator/api"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", api.HealthHandler)
	mux.HandleFunc("POST /api/checkout", api.CheckoutHandler)
	mux.HandleFunc("GET /api/suggestions", api.SuggestionsHandler)

	fileServer := http.FileServer(http.Dir("web"))
	mux.Handle("/", fileServer)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := ":" + port
	log.Printf("server listening on http://localhost%s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
