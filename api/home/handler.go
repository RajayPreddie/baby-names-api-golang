package home

import (
	"babyname-api/config"
	"babyname-api/database"
	"babyname-api/models"
	"encoding/json"
	"log"

	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// Log the database state
	config.Init()
	log.Printf("Database connection status: %v", database.DB)

	// Check if the database is connected
	if database.DB == nil {
		log.Println("Database connection is nil")
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	// Attempt to obtain the result
	var babyNames []models.BabyName
	result := database.DB.Find(&babyNames)
	if result.Error != nil {
		log.Printf("Error retrieving baby names: %v", result.Error)
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(babyNames); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
