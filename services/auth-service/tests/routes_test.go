package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kamil-budzik/hospital-system/auth-service/db"
	"github.com/kamil-budzik/hospital-system/auth-service/routes"
	"github.com/kamil-budzik/hospital-system/auth-service/utils"
	_ "github.com/mattn/go-sqlite3"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	routes.RegisterRoutes(router)
	return router
}

func TestSignupRoute(t *testing.T) {
	setupTestDB()
	defer teardownTestDB()

	router := setupRouter()

	payload := `{"email":"test@example.com","password":"password123"}`
	req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status %d but got %d", http.StatusCreated, w.Code)
	}

	expected := `{"message":"User created"}`
	if w.Body.String() != expected {
		t.Errorf("Expected body %s but got %s", expected, w.Body.String())
	}

	row := db.DB.QueryRow("SELECT email FROM users WHERE email = ?", "test@example.com")
	var email string
	err := row.Scan(&email)
	if err != nil {
		t.Fatalf("User not found in database: %s", err)
	}
	if email != "test@example.com" {
		t.Errorf("Expected email %s but got %s", "test@example.com", email)
	}
}

func TestLoginRoute(t *testing.T) {
	setupTestDB()
	defer teardownTestDB()

	hashedPassword, _ := utils.HashPassword("password123")
	_, err := db.DB.Exec("INSERT INTO users (email, password) VALUES (?, ?)", "test@example.com", hashedPassword)
	if err != nil {
		t.Fatalf("Failed to insert user: %s", err)
	}

	router := setupRouter()

	// Test valid login
	validPayload := `{"email":"test@example.com","password":"password123"}`
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(validPayload)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d but got %d", http.StatusOK, w.Code)
	}

	expected := `{"message":"Login successful","token":"OUDAB"}`
	if w.Body.String() != expected {
		t.Errorf("Expected body %s but got %s", expected, w.Body.String())
	}

	// Test invalid login
	invalidPayload := `{"email":"test@example.com","password":"wrongpassword"}`
	req, _ = http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(invalidPayload)))
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status %d but got %d", http.StatusUnauthorized, w.Code)
	}

	expected = `{"message":"Could not authenticate user"}`
	if w.Body.String() != expected {
		t.Errorf("Expected body %s but got %s", expected, w.Body.String())
	}
}
