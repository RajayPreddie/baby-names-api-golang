package handler

import (
	"babyname-api/database"
	"babyname-api/models"
	"encoding/json"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// Get the database instance
	db := database.GetDB()
	log.Printf("Database instance: %v", db)

	// Check if the database is connected
	if db == nil {
		log.Println("Database connection is nil")
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}

	var babyNames []models.BabyName
	result := db.Find(&babyNames)
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
