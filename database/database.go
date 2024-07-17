package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
)

// Connect establishes a database connection if not already connected
func Connect() {

	// Get the database DSN from the environment variables
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN environment variable is not set")
	}

	log.Printf("Connecting to database")

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt:                              false, // Disable prepared statement caching
		Logger:                                   logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction:                   true,
		DisableAutomaticPing:                     true,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Connected to the database successfully.")

}

func GetDB() *gorm.DB {
	if DB == nil {
		log.Fatal("Database connection is not initialized")
	}
	return DB
}
