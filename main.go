package main

import (
	"fmt"
	"homeassignment/database"
	"homeassignment/router"
	"log"
	"net/http"
)

func main() {
	// Setup db connection
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to db :%v", err)
	}
	defer database.CloseDB(db)

	router.RegisterRoutes()

	// Set the port for the server to listen on
	port := ":8080"
	fmt.Printf("Server starting on port %s\n", port)
	fmt.Println("Access the API at http://localhost:8080")

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(port, nil))
}
