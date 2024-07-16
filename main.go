package main

import (
	"log"
	"net/http"
	"os"

	"babyname-api/api"
	"babyname-api/database"
	"github.com/joho/godotenv"

	_ "babyname-api/docs" // Import generated docs
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Baby Names API
// @version 1.0
// @description This is a sample server for baby names API.
// @host localhost:8080
// @BasePath /
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	database.Connect()

	http.HandleFunc("/api/babynames", api.GetBabyNames)
	http.HandleFunc("/api/babynames/year", api.GetBabyNamesByYear)
	http.HandleFunc("/api/babynames/gender", api.GetBabyNamesByGender)
	http.HandleFunc("/api/babynames/ethnicity", api.GetBabyNamesByEthnicity)
	http.Handle("/swagger/", httpSwagger.WrapHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
