package main

import (
	"homeassignment/database"
	"homeassignment/router"
	"log"
)

func main() {
	// Setup db connection
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to db :%v", err)
	}
	defer database.CloseDB(db)
	container := router.Build(db)
	router := router.RegisterRoutes(container)

	// Set the port for the server to listen on
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
