package main

import (
	"log"
	"net/http"
	"os"

	"babyname-api/api"
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
	http.HandleFunc("/", api.Handler)
	http.HandleFunc("/api/babynames", api.GetBabyNamesHandler)
	http.HandleFunc("/api/babynames/year", api.GetBabyNamesByYear)
	http.HandleFunc("/api/babynames/gender", api.GetBabyNamesByGender)
	http.HandleFunc("/api/babynames/ethnicity", api.GetBabyNamesByEthnicity)

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
