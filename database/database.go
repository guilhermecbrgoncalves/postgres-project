package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	dbHost     = "localhost"
	dbPort     = 5432
	dbUser     = "your_username"
	dbName     = "your_database"
	dbPassword = "your_password"
	dbNonUser  = "db_non_used"
)

func InitDB() *sql.DB {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbName, dbPassword))
	if err != nil {
		panic(err)
	}
	return db
}
