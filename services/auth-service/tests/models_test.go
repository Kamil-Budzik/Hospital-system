package tests

import (
	"database/sql"
	"os"
	"testing"

	"github.com/kamil-budzik/hospital-system/auth-service/db"
	"github.com/kamil-budzik/hospital-system/auth-service/models"
	"github.com/kamil-budzik/hospital-system/auth-service/utils"
	_ "github.com/mattn/go-sqlite3"
)

func setupTestDB() {
	dbFile := "test.db"
	_ = os.Remove(dbFile)

	var err error
	db.DB, err = sql.Open("sqlite3", dbFile)
	if err != nil {
		panic(err)
	}

	_, err = db.DB.Exec(`
		CREATE TABLE users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		)
	`)
	if err != nil {
		panic(err)
	}
}

func teardownTestDB() {
	_ = db.DB.Close()
	_ = os.Remove("test.db")
}

func TestUser_Save(t *testing.T) {
	setupTestDB()
	defer teardownTestDB()

	user := &models.User{
		Email:    "test@example.com",
		Password: "password123",
	}

	err := user.Save()
	if err != nil {
		t.Fatalf("Failed to save user: %s", err)
	}

	row := db.DB.QueryRow("SELECT email, password FROM users WHERE email = ?", user.Email)

	var email, hashedPassword string
	err = row.Scan(&email, &hashedPassword)
	if err != nil {
		t.Fatalf("Failed to retrieve user: %s", err)
	}

	if email != user.Email {
		t.Errorf("Expected email %s, got %s", user.Email, email)
	}

	if utils.CheckPasswordHash(user.Password, hashedPassword) != true {
		t.Errorf("Password was not correctly hashed")
	}
}

func TestUser_ValidateCredentials(t *testing.T) {
	setupTestDB()
	defer teardownTestDB()

	hashedPassword, _ := utils.HashPassword("password123")
	_, err := db.DB.Exec("INSERT INTO users (email, password) VALUES (?, ?)", "test@example.com", hashedPassword)
	if err != nil {
		t.Fatalf("Failed to insert user: %s", err)
	}

	user := &models.User{
		Email:    "test@example.com",
		Password: "password123",
	}
	err = user.ValidateCredentials()
	if err != nil {
		t.Errorf("Expected credentials to be valid, but got error: %s", err)
	}

	user.Password = "wrongpassword"
	err = user.ValidateCredentials()
	if err == nil {
		t.Errorf("Expected credentials to be invalid, but got no error")
	}
}
