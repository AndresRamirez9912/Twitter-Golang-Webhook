package main

import (
	"log"
	"net/http"
	"twitter-webhook/src/database"
	"twitter-webhook/src/handler"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env Variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env File")
	}

	// Start Router
	router := chi.NewRouter()

	// Handlers
	router.Get("/login", handler.LogIn)
	router.Get("/callback", handler.Authorize)
	router.Get("/webhook", handler.WebhookHandler)

	// Create Dynamo Table
	err = database.CreateTableDynamodb()
	if err != nil {
		log.Fatal(err)
	}

	// Start API
	log.Println("Starting Webhook at port 3000")
	err = http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatal("Error starting webhopok at port 3000")
	}
}
