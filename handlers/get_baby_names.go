package handlers

import (
	"babyname-api/database"
	"babyname-api/models"
	"encoding/json"
	"net/http"
)

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
