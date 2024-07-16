package api

import (
	"babyname-api/database"
	"babyname-api/models"
	"encoding/json"
	"net/http"
)

// GetBabyNamesByEthnicity godoc
// @Summary Get baby names by ethnicity
// @Description Get baby names by ethnicity
// @Tags babynames
// @Accept  json
// @Produce  json
// @Param ethnicity query string true "Ethnicity"
// @Success 200 {array} models.BabyName
// @Failure 500 {object} models.ErrorResponse
// @Router /api/babynames/ethnicity [get]
func GetBabyNamesByEthnicity(w http.ResponseWriter, r *http.Request) {
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
