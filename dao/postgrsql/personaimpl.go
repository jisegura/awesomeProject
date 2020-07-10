package postgrsql

import (
	"awesomeProject/models"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type PersonaImpl struct{}

//Allows an person to enter
func (dao PersonaImpl) Create(persona *models.Persona) error {

	query := "INSERT INTO Persona (Id_rol, Nombre, Apellido, Telefono, Mail, Direccion, Fecha_alta, Fecha_baja, Usuario, Contrasenia) " +
			 "VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) " +
			 "RETURNING Id_persona"

	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return err}
	defer stmt.Close()

	row := stmt.QueryRow(persona.Id_rol, persona.Nombre, persona.Apellido, persona.Telefono, persona.Mail, persona.Direccion, time.Now(), pq.NullTime{}, persona.Usuario, persona.Contrasenia)
	err = row.Scan(&persona.Id_persona); if err != nil {return err}

	return nil
}

//Returns all people
func (dao PersonaImpl) Get_All() ([]models.Persona, error) {

	people := make([]models.Persona, 0)
	var row models.Persona

	query := "SELECT * FROM Persona"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return people, err}
	defer stmt.Close()

	rows, err := stmt.Query(); if err != nil {return people, err}

	for rows.Next() {

		err := rows.Scan(&row.Id_persona, &row.Id_rol, &row.Nombre, &row.Apellido, &row.Telefono, &row.Mail, &row.Direccion, &row.Fecha_alta, &row.Fecha_baja, &row.Usuario, &row.Contrasenia); if err != nil {return people, err}
		if !row.Fecha_baja.Valid {
			people = append(people, row)
		}
	}

	return people, nil
}

//Returns a person given an id
func (dao PersonaImpl) Get_By_Id(id int) (models.Persona, error) {

	var person models.Persona

	query := "SELECT * FROM Persona WHERE Id_persona = $1"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return person, err}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(&person.Id_persona, &person.Id_rol, &person.Nombre, &person.Apellido, &person.Telefono, &person.Mail, &person.Direccion, &person.Fecha_alta, &person.Fecha_baja, &person.Usuario, &person.Contrasenia); if err != nil {return person, err}

	return person, nil
}

//Delete a person given their id
func (dao PersonaImpl) Delete(id int) error {

	query := "DELETE FROM Persona WHERE Id_persona = $1"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return err}
	defer stmt.Close()

	row, err := stmt.Exec(id); if err != nil {return err}

	i, _ := row.RowsAffected()
	if i != 1 {return errors.New("Error, se esperaba una fila afectada")}

	return nil
}

//Update an person's attribute
func (dao PersonaImpl) Update_Attribute(id int, attrName string, attr string) error {

	query := "UPDATE Persona SET " + attrName + " = $1 WHERE Id_persona= $2"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return err}
	defer stmt.Close()

	row, err := stmt.Exec(attr, id); if err != nil {return err}

	i, _ := row.RowsAffected()
	if i != 1 {return errors.New("Error, se esperaba una fila afectada")}

	return nil

}

//Returns a attribute of a person
func (dao PersonaImpl) Get_Attr(id int, attrName string) (sql.NullString, error) {

	var attribute sql.NullString

	query := "SELECT "+ attrName + " FROM Persona WHERE id_persona = $1"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil{return attribute, err}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(&attribute); if err != nil {return attribute, err}

	return attribute, nil
}

//Returns the role of a person
func (dao PersonaImpl) Get_Role(id int) (int, error) {

	var role int

	query := "SELECT Id_rol FROM Persona WHERE id_persona = $1"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil{return role, err}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(&role); if err != nil {return role, err}

	return role, nil
}

//Unsubscribe an person
func (dao PersonaImpl) Unsubscribe(id int) error {

	query := "UPDATE Persona SET Fecha_baja = $1 WHERE Id_Persona = $2"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return err}
	defer stmt.Close()

	row, err := stmt.Exec(time.Now(), id); if err != nil {return err}

	i, _ := row.RowsAffected(); if i != 1 {return errors.New("Error, se esperaba una fila afectada")}

	return nil
}

//Returns if a person is active, this is, does not have a discharge date
func is_Valid (id int) (bool, error) {

	var unsubscribeDate pq.NullTime

	query := "SELECT Fecha_baja " +
			 "FROM Persona " +
			 "WHERE id_persona = $1"

	db := getConnection()
	defer db.Close()
	
	stmt, err := db.Prepare(query); if err != nil {return false, err}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(&unsubscribeDate); if err != nil {return false, err}

	return !unsubscribeDate.Valid, nil
}

//------------------------------------------------------------------------------------------------------//
//------------------------VERIFICATION-TO-ACCESS-------------------------------------------------------//
//----------------------------------------------------------------------------------------------------//

//Allows login a user
func (dao PersonaImpl) Login (user string, pass string) (error, bool) {

	var id int

	ok, err := dao.Verify_Data(user, pass); if err != nil {return err, false}

	if ok {

		query := "INSERT INTO Registro_Login (id_persona, hora_login, hora_logout) " +
			"VALUES ($1, $2, $3) " +
			"RETURNING Id_login"

		db := getConnection()
		defer db.Close()

		stmt, err := db.Prepare(query); if err != nil {return err, false}
		defer stmt.Close()

		id, err = dao.Get_User_Id(user); if err != nil {return err, false}

		row := stmt.QueryRow(id, time.Now(), pq.NullTime{})
		err = row.Scan(&id); if err != nil {return err, false}
	}

	return nil, true
}

// Allows deloge a user
func (dao PersonaImpl) Logout () error {

	query := "UPDATE Registro_Login " +
		"SET hora_logout = $1 "

	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return err}
	defer stmt.Close()

	row, err := stmt.Exec(time.Now()); if err != nil {return err}

	i, _ := row.RowsAffected()
	if i != 1 {
		return errors.New("Error, se esperaba una fila afectada")
	}

	err = Delete_Reg(); if err != nil {return err}

	return nil
}

//Returns if there is a logged in user
func (dao PersonaImpl) User_Is_Logged_In () (int, error) {

	var id_persona int

	query := "SELECT Id_persona FROM Registro_Login"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return id_persona, err}
	defer stmt.Close()

	row := stmt.QueryRow()
	err = row.Scan(&id_persona); if err != nil {
		if err == sql.ErrNoRows {
			return  id_persona, nil
		}
	}

	return id_persona, err
}

//Allows removing the user from the login table
func Delete_Reg () error {

	query := "DELETE FROM Registro_Login"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return err}
	defer stmt.Close()

	row, err := stmt.Exec(); if err != nil {return err}

	i, _ := row.RowsAffected(); if i != 1 {return errors.New("Error, se esperaba una fila afectada")}

	return nil
}

func HashPassword(password string) (string, error) {

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), 10); if err != nil {return "", err}

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

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {

	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomString(n int) (string, error) {

	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	bytes, err := GenerateRandomBytes(n)
	if err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}

	hashBase64, err := HashPassword(string(bytes))
	log.Print("pass ", string(bytes), "  hash: ", hashBase64)
	otro, err := HashPassword("cocacola")
	log.Print("otro: ", otro)

	return hashBase64, nil
}

//
func (dao PersonaImpl) Change_Password (user string, pass string, newPass string) error {

	id, err := dao.Get_User_Id(user); if err != nil {return err}

	err = verify_Pass(id, pass); if err != nil {return err}

	err = dao.Update_Attribute(id, "contrasenia", newPass); if err != nil {return err}

	return nil
}

//Returns if it is a user's first login
func (dao PersonaImpl)Is_First_Login (user string) (error, bool) {

	var pass sql.NullString

	query := "SELECT Contrasenia FROM Persona " +
			 "WHERE Usuario SIMILAR TO $1"

	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err, false
	}
	defer stmt.Close()

	row := stmt.QueryRow(&user)
	err = row.Scan(&pass)
	if err != nil {
		return err, false
	}

	return nil, !pass.Valid
}

// Allows to verify the data sent by the user
func (dao PersonaImpl) Verify_Data(user string, pass string) (bool, error) {

	id, err := dao.Get_User_Id(user); if err != nil {return false, errors.New("Error, el usuario no existe")}

	active, err := is_Valid(id); if err != nil {return false, err}

	if active {

		err := verify_Pass(id, pass); if err != nil {return false, err}
		return true, nil
		}
	return false, errors.New("Error, el usuario no está activo")

}

//Returns a person given an user
func (dao PersonaImpl) Get_User_Id (user string) (int, error) {

	var id int

	query := "SELECT Id_persona FROM Persona " +
		"WHERE Usuario SIMILAR TO $1"

	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return id, err}
	defer stmt.Close()

	row := stmt.QueryRow(&user)
	err = row.Scan(&id); if err != nil {return id, err}

	return id, nil
}

//Verify the pass sent fro the user
func verify_Pass (id int, pass string) error {

	query := "SELECT contrasenia " +
			 "FROM Persona " +
			 "WHERE id_persona = $1"

	db := getConnection()
	defer db.Close()

	var password string

	stmt, err := db.Prepare(query); if err != nil {return err}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(&password); if err != nil {return err}

	if ComparePassword(password, pass) {return nil}

	err = errors.New("Error, contraseña incorrecta")
	return err
}

func (dao PersonaImpl) Reset_Password (user string) (error) {

	query := "UPDATE Persona SET Contrasenia = $1" +
		" WHERE Usuario SIMILAR TO $2"

	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return err}
	defer stmt.Close()

	row, err := stmt.Exec(sql.NullString{}, user); if err != nil {return err}

	i, _ := row.RowsAffected(); if i != 1 {return errors.New("Error, se esperaba una fila afectada")}

	return nil
}

func (dao PersonaImpl) Set_Password(user string, pass string) error {

	id, err := dao.Get_User_Id(user); if err != nil {return err}

	password, err := HashPassword(pass); if err != nil {return err}

	return dao.Update_Attribute(id, "contrasenia", password)
}