package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

	validPayload := `{"email":"test@example.com","password":"password123"}`
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(validPayload)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d but got %d", http.StatusOK, w.Code)
	}

	var responseBody map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &responseBody); err != nil {
		t.Fatalf("Failed to parse response body: %s", err)
	}

	if responseBody["message"] != "Login successful" {
		t.Errorf("Expected message 'Login successful' but got '%v'", responseBody["message"])
	}

	tokenString, ok := responseBody["token"].(string)
	if !ok || tokenString == "" {
		t.Fatal("Expected a valid token in the response")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("desnuts"), nil
	})
	if err != nil || !token.Valid {
		t.Errorf("Token is invalid: %s", err)
	}

	invalidPayload := `{"email":"test@example.com","password":"wrongpassword"}`
	req, _ = http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(invalidPayload)))
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status %d but got %d", http.StatusUnauthorized, w.Code)
	}

	expectedError := `{"message":"Could not authenticate user"}`
	if w.Body.String() != expectedError {
		t.Errorf("Expected body %s but got %s", expectedError, w.Body.String())
	}
}
