package api

import (
	"babyname-api/config"
	"babyname-api/database"
	"babyname-api/models"
	"encoding/json"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func init() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Initialize configurations
	config.Init()
}

func GetBabyNames(w http.ResponseWriter, r *http.Request) {
	var babyNames []models.BabyName
	result := database.DB.Find(&babyNames)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(babyNames)
}
