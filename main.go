package main

import (
	"log"
	"net/http"
	"twitter-webhook/src/handler"

	"github.com/go-chi/chi"
)

func main() {
	// Start Router
	router := chi.NewRouter()

	// Handlers
	router.Get("/", handler.IndexHandler)
	router.Get("/webhook", handler.WebhookHandler)

	// Start Webhook
	log.Println("Starting Webhook at port 3000")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatal("Error starting webhopok at port 3000")
	}
}
