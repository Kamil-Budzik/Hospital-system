package db_test

import (
	"os"
	"testing"

	"github.com/kamil-budzik/hospital-system/auth-service/db"
)

func TestInitDB(t *testing.T) {
	err := db.InitDB()
	if err != nil {
		t.Fatalf("Failed to initialize DB: %s", err)
	}
	defer os.Remove("api.db")

	_, err = db.DB.Exec("SELECT * FROM users")
	if err != nil {
		t.Errorf("Users table not created: %s", err)
	}
}
