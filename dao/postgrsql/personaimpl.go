package postgrsql

import (
	"awesomeProject/models"
	"database/sql"
	"errors"
	"github.com/lib/pq"
	"time"
)

type EmpleadoImpl struct{}

//Allows an employee to enter
func (dao EmpleadoImpl) Create(empleado *models.Empleado) error {

	query := "INSERT INTO empleado (firstname, lastname) VALUES ($1, $2) RETURNING id_empleado"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return err}
	defer stmt.Close()

	row := stmt.QueryRow(empleado.FirstName, empleado.LastName)
	err = row.Scan(&empleado.Id_empleado); if err != nil {return err}

	return nil
}

//Returns all employees
func (dao EmpleadoImpl) Get_All() ([]models.Empleado, error) {

	employees := make([]models.Empleado, 0)
	var row models.Empleado

	query := "SELECT * FROM empleado"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return employees, err}
	defer stmt.Close()

	rows, err := stmt.Query(); if err != nil {return employees, err}

	for rows.Next() {

		err := rows.Scan(&row.Id_empleado, &row.FirstName, &row.LastName, &row.FechaBaja, &row.Id_login); if err != nil {return employees, err}

		if !row.FechaBaja.Valid {
			employees = append(employees, row)
		}
	}

	return employees, nil
}

//Returns an employee given an id
func (dao EmpleadoImpl) Get_By_Id(id int) (models.Empleado, error) {

	var employee models.Empleado

	query := "SELECT * FROM empleado WHERE id_empleado = $1"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return employee, err}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(&employee.Id_empleado, &employee.FirstName, &employee.LastName, &employee.FechaBaja, &employee.Id_login); if err != nil {return employee, err}

	return employee, nil
}

//Delete an employee given their id
func (dao EmpleadoImpl) Delete(id int) error {

	query := "DELETE FROM empleado WHERE id_empleado = $1"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return err}
	defer stmt.Close()

	row, err := stmt.Exec(id); if err != nil {return err}

	i, _ := row.RowsAffected()
	if i != 1 {
		return errors.New("Error, se esperaba una fila afectada")
	}

	return nil
}

//Update an employee's first and last name
func (dao EmpleadoImpl) Update_Name(empleado *models.Empleado) error {

	query := "UPDATE empleado SET firstname = $1, lastname = $2 WHERE id_empleado= $3"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return err}
	defer stmt.Close()

	row, err := stmt.Exec(empleado.FirstName, empleado.LastName, empleado.Id_empleado); if err != nil {return err}

	i, _ := row.RowsAffected()
	if i != 1 {
		return errors.New("Error, se esperaba una fila afectada")
	}

	return nil
}

//Unsubscribe an employee
func (dao EmpleadoImpl) Unsubscribe(id int) error {

	query := "UPDATE empleado SET fechaBaja = $1 WHERE id_empleado = $2"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return err}
	defer stmt.Close()

	row, err := stmt.Exec(time.Now(), id); if err != nil {return err}

	i, _ := row.RowsAffected()
	if i != 1 {
		return errors.New("Error, se esperaba una fila afectada")
	}

	return nil
}

//Returns the name of an employee
func Get_Nombre(id int64) (string, error) {

	var name string

	query := "SELECT firstname FROM empleado WHERE id_empleado = $1"
	db := getConnection()
	defer db.Close()
	
	stmt, err := db.Prepare(query); if err != nil{return name, err}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(&name); if err != nil {return name, err}

	return name, nil
}

//Returns if an employee is active, this is, does not have a discharge date
func Is_Valid (id int) (bool, error) {

	var unsubscribeDate pq.NullTime

	query := "SELECT fechaBaja FROM empleado WHERE id_empleado = $1"
	db := getConnection()
	defer db.Close()
	
	stmt, err := db.Prepare(query); if err != nil {return false, err}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(&unsubscribeDate); if err != nil {return false, err}

	return !unsubscribeDate.Valid, nil
}

//Returns if an employee already has a user
func Has_User (id int) (bool, error) {

	var id_login sql.NullInt64

	query := "SELECT id_login from Empleado WHERE id_empleado = $1"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return false, err}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(&id_login); if err != nil {
		if err == sql.ErrNoRows {return false, nil}
		return false, err
	}

	return true, nil
}

//Set the id_login to an employee
func Set_Id_Login(id_empleado int, id_login int) error {

	query := "UPDATE empleado SET id_login = $1 WHERE id_empleado = $2"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query);
	if err != nil {return err}
	defer stmt.Close()

	row, err := stmt.Exec(id_login, id_empleado); if err != nil {return err}

	i, _ := row.RowsAffected()
	if i != 1 {
		return errors.New("Error, se esperaba una fila afectada")
	}

	return nil
}