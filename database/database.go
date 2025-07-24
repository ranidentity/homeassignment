package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB establishes a connection to the PostgreSQL database.
func ConnectDB() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: Could not load .env file: %v. Assuming environment variables are set directly.\n", err)
	}

	// Retrieve database connection details from environment variables.
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	// Open a database connection.
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error opening GORM database connection: %w", err)
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, fmt.Errorf("error getting underlying sql.DB from GORM: %w", err)
	}

	// Set connection pool settings (optional but recommended for performance).
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	// Ping the database to verify the connection is alive.
	err = sqlDB.Ping()
	if err != nil {
		sqlDB.Close() // Close the connection if ping fails
		return nil, fmt.Errorf("error connecting to the database (ping failed): %w", err)
	}

	log.Println("Successfully connected to the PostgreSQL database!")
	return gormDB, nil
}

// CloseDB closes the database connection.
func CloseDB(gormDB *gorm.DB) {
	if gormDB != nil {
		sqlDB, err := gormDB.DB()
		if err != nil {
			log.Printf("Error getting underlying sql.DB for closing: %v\n", err)
			return
		}
		err = sqlDB.Close()
		if err != nil {
			log.Printf("Error closing GORM database connection: %v\n", err)
		} else {
			log.Println("GORM database connection closed.")
		}
	}
}
