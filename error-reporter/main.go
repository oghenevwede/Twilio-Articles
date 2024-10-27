package main

import (
	"error-reporter/handlers"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Access the environment variable
	sendgridAPIKey := os.Getenv("SENDGRID_API_KEY")
	if sendgridAPIKey == "" {
		log.Fatal("SENDGRID_API_KEY is not set")
	} else {
		log.Println("SENDGRID_API_KEY loaded successfully")
		log.Println("Using SendGrid API Key:", sendgridAPIKey) // Debugging only, remove after testing
	}

	http.HandleFunc("/errors", handlers.HandleErrorReport)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
