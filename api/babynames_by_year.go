package api

import (
	"babyname-api/database"
	"babyname-api/models"
	"encoding/json"
	"net/http"
)

// GetBabyNamesByYear godoc
// @Summary Get baby names by birth year
// @Description Get baby names by birth year
// @Tags babynames
// @Accept  json
// @Produce  json
// @Param year query int true "Birth Year"
// @Success 200 {array} models.BabyName
// @Failure 500 {object} models.ErrorResponse
// @Router /api/babynames/year [get]
func GetBabyNamesByYear(w http.ResponseWriter, r *http.Request) {
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
