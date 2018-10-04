package postgrsql

import (
	"awesomeProject/models"
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type LoginImpl struct{}

func AddLogin(login models.Login, id int) error {

	activo, err := InsertActivo(id)
	if err != nil {
		return err
	}
	if activo {
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

		err = UpdateLogin(id, login.Id_login)
		if err != nil {
			return err
		}

		return nil
	}

	err = errors.New("Error, usuario inválido")
	return err
}

func Login(login models.Login) (bool, error) {

	var ok = false

	activo, err := Activo(login.Id_login)
	if err != nil {
		return ok, err
	}
	if activo {

		existe, err := existeUsuario(login.Usuario)
		if err != nil {
			return ok, err
		}
		if existe {

			query := "SELECT password FROM login WHERE id_login = $1"
			db := getConnection()
			defer db.Close()

			var password string

			stmt, err := db.Prepare(query)
			if err != nil {
				return ok, err
			}
			defer stmt.Close()

			row := stmt.QueryRow(login.Id_login)
			err = row.Scan(&password)
			if err != nil {
				return ok, err
			}

			if ComparePassword(password, login.Password) {
				return true, nil
			}
			err = errors.New("Error, contraseña incorrecta")
			return false, err
		}

		return false, err
	}

	err = errors.New("Error, usuario inválido")
	return false, err
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
