package models

import (
	"errors"

	"github.com/kamil-budzik/hospital-system/auth-service/db"
	"github.com/kamil-budzik/hospital-system/auth-service/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `INSERT INTO users(email, password) VALUES (?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	// TODO: might need this id later
	result, err := stmt.Exec(u.Email, hashPassword)
	if err != nil {
		return err
	}
	_ = result

	// userId, err := result.LastInsertId()
	// u.ID = userId

	return err
}

func (u *User) ValidateCredentials() error {
	query := `SELECT id, password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)
	if err != nil {
		return errors.New("Credentials invalid")
	}

	isPasswordValid := utils.CheckPasswordHash(u.Password, retrievedPassword)
	if !isPasswordValid {
		return errors.New("Credentials invalid")
	}

	return nil
}
