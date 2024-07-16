package main

import (
	"babyname-api/config"
	"babyname-api/database"
	"babyname-api/models"
	"encoding/json"
	"net/http"
)

func getBabyNamesByEthnicity(w http.ResponseWriter, r *http.Request) {
	ethnicity := r.URL.Query().Get("ethnicity")
	var babyNames []models.BabyName
	result := database.DB.Where("ethnicity = ?", ethnicity).Find(&babyNames)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(babyNames)
}

func main() {
	config.Init()
	http.HandleFunc("/", getBabyNamesByEthnicity)
	http.ListenAndServe(":8080", nil)
}
