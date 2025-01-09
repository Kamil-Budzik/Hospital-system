package tests

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kamil-budzik/hospital-system/auth-service/utils"
)

var secretKey = []byte("desnuts")

func TestValidateToken(t *testing.T) {
	createTestToken := func(secret []byte, claims jwt.MapClaims) string {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, _ := token.SignedString(secret)
		return tokenString
	}

	validToken := createTestToken(secretKey, jwt.MapClaims{
		"sub": 12345,
		"exp": time.Now().Add(time.Hour).Unix(),
	})
	expiredToken := createTestToken(secretKey, jwt.MapClaims{
		"sub": 12345,
		"exp": time.Now().Add(-time.Hour).Unix(),
	})

	tests := []struct {
		name        string
		authHeader  string
		expectError bool
	}{
		{"Valid token", "Bearer " + validToken, false},
		{"Expired token", "Bearer " + expiredToken, true},
		{"Missing Bearer prefix", validToken, true},
		{"Invalid token format", "Bearer invalid.token.here", true},
		{"Missing token", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := utils.ValidateToken(tt.authHeader)
			if (err != nil) != tt.expectError {
				t.Errorf("expected error: %v, got: %v", tt.expectError, err)
			}
		})
	}
}
