package config

import (
	"babyname-api/database"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Init() {
	// Load environment variables from .env file only in local development
	if os.Getenv("VERCEL_ENV") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}

	// Connect to the database
	database.Connect()
}
