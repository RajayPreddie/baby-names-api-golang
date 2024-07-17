package home

import (
	"babyname-api/database"
	"babyname-api/models"
	"encoding/json"
	"log"
	"os"

	"net/http"

	"github.com/joho/godotenv"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	var babyNames []models.BabyName
	// Conditionally load .env file only during local development
	if os.Getenv("VERCEL_ENV") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}

	// Check if the database is connected
	if database.DB == nil {
		database.Connect()
	}
	result := database.DB.Find(&babyNames)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(babyNames)
}
