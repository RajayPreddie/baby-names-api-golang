package main

import (
	"log"
	"net/http"
	"os"

	"babyname-api/api/get_baby_names"
	"babyname-api/api/get_baby_names_by_ethnicity"
	"babyname-api/api/get_baby_names_by_gender"
	"babyname-api/api/get_baby_names_by_year"
	"babyname-api/api/home"

	"babyname-api/config"

	_ "babyname-api/docs"
	"github.com/joho/godotenv"
	"github.com/swaggo/http-swagger"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Initialize configurations
	config.Init()

	// Set up API routes
	http.HandleFunc("/", home.Handler)
	http.HandleFunc("/api/babynames", get_baby_names.Handler)
	http.HandleFunc("/api/babynames/year", get_baby_names_by_year.Handler)
	http.HandleFunc("/api/babynames/gender", get_baby_names_by_gender.Handler)
	http.HandleFunc("/api/babynames/ethnicity", get_baby_names_by_ethnicity.Handler)

	// Serve the Swagger UI at /swagger/
	http.Handle("/swagger/", httpSwagger.WrapHandler)

	// Get the port from environment variables or use a default value
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the server
	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
