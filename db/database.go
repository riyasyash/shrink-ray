package db

import "database/sql"

var database *sql.DB

func GetDatabase() (*sql.DB, error) {
	if database != nil {
		return database, nil
	}
	db, err := sql.Open("sqlite3", "./shrinkray.db")
	return db, err
}
