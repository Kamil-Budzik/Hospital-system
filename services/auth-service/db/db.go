package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() error {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		return fmt.Errorf("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	err = createTable()
	if err != nil {
		return fmt.Errorf("Failed to create users table")
	}

	return nil
}

func createTable() error {
	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		)
	`

	_, err := DB.Exec(createUsersTable)

	return err
}
