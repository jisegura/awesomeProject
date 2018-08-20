package postgrsql

import (
	"awesomeProject/models"
	"errors"
	"log"
)

type EmpleadoImpl struct{}

func (dao EmpleadoImpl) Create(empleado *models.Empleado) error {
	query := "INSERT INTO empleado (firstname, lastname) VALUES ($1, $2) RETURNING id_empleado"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	row := stmt.QueryRow(empleado.FirstName, empleado.LastName)
	row.Scan(&empleado.Id_empleado)
	return nil
}

func (dao EmpleadoImpl) GetAll() ([]models.Empleado, error) {
	empleados := make([]models.Empleado, 0)
	query := "SELECT * FROM empleado"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return empleados, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return empleados, err
	}

	for rows.Next() {
		var row models.Empleado
		err := rows.Scan(&row.Id_empleado, &row.FirstName, &row.LastName)
		if err != nil {
			return empleados, err
		}
		empleados = append(empleados, row)
	}
	return empleados, nil
}

//SELECT BY ID
func (dao EmpleadoImpl) GetById(id int) (models.Empleado, error) {

	var p models.Empleado

	query := "SELECT * FROM empleado WHERE id_empleado = $1"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return p, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(&p.Id_empleado, &p.FirstName, &p.LastName)
	if err != nil {
		return p, err
	}

	return p, nil
}

//DELETE
func (dao EmpleadoImpl) Delete(id int) error {

	query := "DELETE FROM empleado WHERE id_empleado = $1"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	i, _ := r.RowsAffected()
	if i != 1 {
		return errors.New("Error, se esperaba una fila afectada")
	}

	return nil
}

//UPDATE
func (dao EmpleadoImpl) Update(empleado *models.Empleado) error {

	query := "UPDATE empleado SET firstname = $1, lastname = $2 WHERE id_empleado= $3"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	r, err := stmt.Exec(empleado.FirstName, empleado.LastName, empleado.Id_empleado)
	if err != nil {
		log.Fatal(err)
	}

	i, _ := r.RowsAffected()
	if i != 1 {
		return errors.New("Error, se esperaba una fila afectada")
	}

	return nil
}
