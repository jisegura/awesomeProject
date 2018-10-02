package postgrsql

import (
	"awesomeProject/models"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type LoginImpl struct{}

func (dao LoginImpl) AddLogin(login models.Login) error {

	query := "INSERT INTO Login (usuario, password) VALUES ($1, $2) RETURNING id_login"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	password, err := HashPassword(login.Password)
	if err != nil {
		return err
	}

	row := stmt.QueryRow(login.Usuario, password)
	err = row.Scan(&login.Id_login)
	if err != nil {
		return err
	}

	return nil
}

func (dao LoginImpl) Login(login models.Login) (bool, error) {

	query := "SELECT password FROM login WHERE id_login = $1"
	db := getConnection()
	defer db.Close()

	var correct bool
	var password string

	stmt, err := db.Prepare(query)
	if err != nil {
		return correct, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(login.Id_login)
	err = row.Scan(&password)
	if err != nil {
		return correct, err
	}

	return ComparePassword(password, login.Password), err
}

func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), 10 /*cost*/)
	if err != nil {
		return "", err
	}

	// Encode the entire thing as base64 and return
	hashBase64 := base64.StdEncoding.EncodeToString(hashedBytes)

	return hashBase64, nil
}

func ComparePassword(hashBase64, testPassword string) bool {

	// Decode the real hashed and salted password so we can
	//	// split out the salt
	hashBytes, err := base64.StdEncoding.DecodeString(hashBase64)
	if err != nil {
		fmt.Println("Error, we were given invalid base64 string", err)
		return false
	}

	err = bcrypt.CompareHashAndPassword(hashBytes, []byte(testPassword))
	return err == nil
}
