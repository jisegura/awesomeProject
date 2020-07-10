package postgrsql

import (
	"awesomeProject/models"
	"encoding/json"
	"errors"
)

type RolImpl struct {}

//Allows to add a new rol
func (dao RolImpl) Create(rol models.Rol) (models.Rol, error) {

	var newRol models.Rol

	query := "INSERT INTO Rol (Nombre, Descripcion)" +
			 "VALUES ($1, $2) " +
			 "RETURNING Id_rol"

	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return newRol, err}
	defer stmt.Close()

	row := stmt.QueryRow(&rol.Nombre, &rol.Descripcion)
	err = row.Scan(&rol.Id_rol); if err != nil {return newRol, err}

	newRol, err = dao.Get_By_Id(rol.Id_rol); if err != nil {return newRol, err}

	return newRol, nil
}

//Returns a role given an id
func (dao RolImpl) Get_By_Id (id int) (models.Rol, error) {

	var role models.Rol

	query := "SELECT * FROM Rol WHERE Id_rol = $1"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return role, err}
	defer stmt.Close()

	row := stmt.QueryRow(id); if err != nil {return role, err}
	err = row.Scan(&role.Id_rol, &role.Nombre, &role.Descripcion); if err != nil {return role, err}

	return role, nil
}

//Returns all roles
func (dao RolImpl) Get_All () ([]models.Rol, error) {

	roles := make([]models.Rol, 0)
	var role models.Rol

	query := "SELECT * FROM Rol"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return roles, err}
	defer stmt.Close()

	rows, err := stmt.Query(); if err != nil {return roles, err}

	for rows.Next() {

		err := rows.Scan(&role.Id_rol, &role.Nombre, &role.Descripcion); if err != nil {return roles, err}
		roles = append(roles, role)
	}
	json.Marshal(roles)

	return roles, err
}

//Returns an attribute of the role
func (dao RolImpl) Get_Attr (id int, attrName string) (string, error) {

	var attr string

	query := "SELECT " + attrName + " FROM Rol WHERE Id_rol = $1"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return attr, err}
	defer stmt.Close()

	row := stmt.QueryRow(&id)

	err = row.Scan(&attr); if err != nil {return attr, err}

	return attr, err
}

//Allows updating an attribute
func (dao RolImpl) Update (id int, attrName string, attr string) (models.Rol, error) {

	var role models.Rol

	query := "UPDATE Rol SET " + attrName + " = $1 WHERE Id_rol = $2"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return role, err}
	defer stmt.Close()

	row, err := stmt.Exec(attr, id); if err != nil {return role, err}
	i , _ := row.RowsAffected(); if i != 1 {return role, errors.New("Error, se esperaba una fila afectada")}

	return dao.Get_By_Id(id)
}

