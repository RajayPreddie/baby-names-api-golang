package database

import (
	"log"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB   *gorm.DB
	once sync.Once
)

// Connect establishes a database connection if not already connected
func Connect() {
	once.Do(func() {
		dsn := os.Getenv("DB_DSN")
		var err error
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		log.Println("Connected to the database successfully.")
	})
}
