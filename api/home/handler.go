package home

import (
	"babyname-api/database"
	"babyname-api/models"
	"encoding/json"
	"log"

	"net/http"

	"github.com/joho/godotenv"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	var babyNames []models.BabyName
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Initialize configurations
	database.Connect()
	// Check if the database is connected
	if database.DB == nil {
		log.Fatalf("Database connection is nil")
	}
	result := database.DB.Find(&babyNames)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(babyNames)
}
