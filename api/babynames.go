package api

import (
	"babyname-api/database"
	"babyname-api/models"
	"encoding/json"
	"net/http"
)

// GetBabyNames godoc
// @Summary Get all baby names
// @Description Get all baby names
// @Tags babynames
// @Accept  json
// @Produce  json
// @Success 200 {array} models.BabyName
// @Failure 500 {object} models.ErrorResponse
// @Router /api/babynames [get]
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
