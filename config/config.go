package config

import (
	"babyname-api/database"
	"github.com/joho/godotenv"
	"log"
)

func Init() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Connect to the database
	database.Connect()
}
