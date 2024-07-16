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

// GetBabyNameByID godoc
// @Summary Get baby name by ID
// @Description Get baby name by ID
// @Tags babynames
// @Accept  json
// @Produce  json
// @Param id query string true "BabyName ID"
// @Success 200 {object} models.BabyName
// @Failure 404 {object} models.ErrorResponse
// @Router /api/babynames/id [get]
func GetBabyNameByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var babyName models.BabyName
	result := database.DB.First(&babyName, "id = ?", id)
	if result.Error != nil {
		http.Error(w, "Record not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(babyName)
}

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
