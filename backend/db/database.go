package db

import (
	"os"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var database *sql.DB

func GetDatabase() (*sql.DB, error) {
	if database != nil {
		return database, nil
	}
	dbString, exists := os.LookupEnv("DB_CONNECTION_STRING")
	if !exists {
		return nil, fmt.Errorf("connection string not specified")
	}
	db, err := sql.Open("postgres", dbString)
	return db, err
}
