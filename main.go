package main

import (
	"homeassignment/database"
	"homeassignment/router"
	"log"

	"github.com/go-playground/validator/v10"
)

// @title homeassignment
// @version 1.0
// @host localhost:8080
// @BasePath /
func main() {
	// Setup db connection
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to db :%v", err)
	}
	defer database.CloseDB(db)
	validate := validator.New()
	// Register the custom validation function for "decimal2" tag here in main.go

	container := router.Build(db, validate)
	router := router.RegisterRoutes(container)

	// Set the port for the server to listen on
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
