package api

import (
	"log"
	"net/http"
	"os"

	"babyname-api/config"
	"babyname-api/handlers"

	_ "babyname-api/docs"
	"github.com/joho/godotenv"
	"github.com/swaggo/http-swagger"
)

func init() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Initialize configurations
	config.Init()
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// Set up API routes
	http.HandleFunc("/api/babynames", handlers.GetBabyNames)
	http.HandleFunc("/api/babynames/year", handlers.GetBabyNamesByYear)
	http.HandleFunc("/api/babynames/gender", handlers.GetBabyNamesByGender)
	http.HandleFunc("/api/babynames/ethnicity", handlers.GetBabyNamesByEthnicity)
	http.Handle("/swagger/", httpSwagger.WrapHandler)

	// Serve the request
	w.Write([]byte("Hello, World!"))
}

func main() {
	// Get the port from environment variables or use a default value
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the server
	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
