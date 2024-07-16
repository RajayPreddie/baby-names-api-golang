package main

import (
	"babyname-api/config"
	"babyname-api/database"
	"babyname-api/models"
	"encoding/json"
	"net/http"
)

func getBabyNamesByGender(w http.ResponseWriter, r *http.Request) {
	gender := r.URL.Query().Get("gender")
	var babyNames []models.BabyName
	result := database.DB.Where("gender = ?", gender).Find(&babyNames)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(babyNames)
}

func main() {
	config.Init()
	http.HandleFunc("/", getBabyNamesByGender)
	http.ListenAndServe(":8080", nil)
}
