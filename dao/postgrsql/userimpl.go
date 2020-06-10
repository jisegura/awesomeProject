package postgrsql

import (
	"awesomeProject/models"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserImpl struct{}

//Allows add a new user
func (dao UserImpl) AddUser(user models.User, id int) error {

	if user.Empleado {
		valid, err := Is_Valid(id); if err != nil {
			return errors.New("Error, usario inválido")
		}
		if valid {
			exists, err := Has_User(id); if err != nil {
				return err
			}; if exists {
				return errors.New("Error, el empleado ya tiene usuario registrado")
			}
		} else {
			return errors.New("Error, el usuario pertenece a un empleado dado de baja")}
	}

	query := "INSERT INTO Usuario (usuario, password) VALUES ($1, $2) RETURNING id_login"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {
		return err
	}
	defer stmt.Close()

	password, err := HashPassword(user.Password); if err != nil {
		return err
	}

	row := stmt.QueryRow(user.Usuario, password)
	err = row.Scan(&user.Id_login); if err != nil {
		return err
	}

	if user.Empleado {

		err := Set_Id_Login(id, user.Id_login);
		if err != nil {
			return err
		}
	}

	return nil
}

//Allows login a user
func (dao UserImpl) Login (user models.User) (error, bool) {

	ok, err := verify_Data(user); if err != nil {
		return err, false
	}

	if ok {

		query := "INSERT INTO Login (id_login, hora_conexion, hora_desconexion) VALUES ($1, $2, $3) RETURNING id_login"

		db := getConnection()
		defer db.Close()

		stmt, err := db.Prepare(query); if err != nil {
			return err, false
		}
		defer stmt.Close()

		row := stmt.QueryRow(user.Id_login, time.Now(), time.Time{})
		err = row.Scan(&user.Id_login); if err != nil {
			return err, false
		}
	}
	return nil, true
}

// Allows deloge a user
func (dao UserImpl) Logout (id int)(error) {

	query := "UPDATE Login SET hora_desconexion = $1 WHERE Id_Login = $2"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {
		return err
	}
	defer stmt.Close()

	row, err := stmt.Exec(time.Now(), id); if err != nil {
		return err
	}

	i, _ := row.RowsAffected()
	if i != 1 {
		return errors.New("Error, se esperaba una fila afectada")
	}

	return nil
}

// Allows to verify the data sent by the user
func verify_Data(user models.User) (bool, error) {

	exist, err := user_Exists(user.Usuario); if err != nil {
		return false, err
	}

	if exist {

		active, err := user_Is_Active(user.Id_login)
		if err != nil {
			return false, err
		}
		if active {

			err := verify_Pass(user); if err != nil {
				return false, err
			}
			return true, nil
		}
		return false, errors.New("Error, el usuario no esta activo")
	}
	return false, errors.New("Error, el usuario no existe")
}

//Verify the pass sent fro the user
func verify_Pass (user models.User) error {

	query := "SELECT password FROM Usuario WHERE id_login = $1"
	db := getConnection()
	defer db.Close()

	var password string

	stmt, err := db.Prepare(query);
	if err != nil {
		return err
	}
	defer stmt.Close()

	row := stmt.QueryRow(user.Id_login)
	err = row.Scan(&password)
	if err != nil {
		return err
	}

	if ComparePassword(password, user.Password) {
		return nil
	}

	err = errors.New("Error, contraseña incorrecta")
	return err
}

//Returns if there is a logged in user
func user_Is_Logged_In () (bool, error) {

	var horaDesconexion pq.NullTime

	query := "SELECT hora_desconexion FROM Login"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {
		return false, err
	}
	defer stmt.Close()

	row := stmt.QueryRow()
	err = row.Scan(&horaDesconexion); if err != nil {
		return false, err
	}

	return !horaDesconexion.Valid, nil
}

//Returns if a user is active
func user_Is_Active (id int) (bool, error) {

	var active bool

	query := "SELECT activo FROM Usuario WHERE id_login = $1"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {
		return false, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(&active); if err != nil {
		return false, err
	}

	return active, nil
}

//Returns if username already exists
func user_Exists(nombre string) (bool, error) {

	query := "SELECT id_login FROM Usuario WHERE usuario SIMILAR TO $1"
	db := getConnection()
	defer db.Close()

	var id_login sql.NullInt64

	stmt, err := db.Prepare(query); if err != nil {
		return false, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(nombre)
	err = row.Scan(&id_login); if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), 10 /*cost*/); if err != nil {
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
