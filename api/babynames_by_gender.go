package api

import (
	"babyname-api/database"
	"babyname-api/models"
	"encoding/json"
	"net/http"
)

// GetBabyNamesByGender godoc
// @Summary Get baby names by gender
// @Description Get baby names by gender
// @Tags babynames
// @Accept  json
// @Produce  json
// @Param gender query string true "Gender"
// @Success 200 {array} models.BabyName
// @Failure 500 {object} models.ErrorResponse
// @Router /api/babynames/gender [get]
func GetBabyNamesByGender(w http.ResponseWriter, r *http.Request) {
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
