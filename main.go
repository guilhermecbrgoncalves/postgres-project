package main

import (
	"log"
	"net/http"
	"postgres-project/database"
	"postgres-project/router"

	_ "github.com/lib/pq"
)

const (
	dbHost     = "localhost"
	dbPort     = 5432
	dbUser     = "your_username"
	dbName     = "your_database"
	dbPassword = "your_password"
)

func main() {
	// Initialize the database
	db := database.InitDB()
	defer db.Close()

	// Initialize the router
	r := router.InitRouter(db)

	// Start the HTTP server
	port := ":8080"
	log.Printf("Listening on: %s", port)
	log.Fatal(http.ListenAndServe(port, r))
}
