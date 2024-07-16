package api

import (
	"babyname-api/config"
	"babyname-api/database"
	"babyname-api/models"
	"encoding/json"
	"net/http"
)

func init() {
	config.Init()
}
func GetBabyNamesByYearHandler(w http.ResponseWriter, r *http.Request) {
	year := r.URL.Query().Get("year")
	var babyNames []models.BabyName
	result := database.DB.Where("birth_year = ?", year).Find(&babyNames)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(babyNames)
}
