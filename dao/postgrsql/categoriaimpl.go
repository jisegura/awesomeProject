package postgrsql

import (
	"awesomeProject/models"
	"errors"
)

type CategoriaImpl struct{}

//Allows to create a new category
func (dao CategoriaImpl) Create(category *models.Categoria) error {

	query := "INSERT INTO categoria (nombre) VALUES ($1) RETURNING id_categoria"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return err}
	defer stmt.Close()

	row := stmt.QueryRow(category.Nombre)
	err = row.Scan(&category.Id_categoria); if err != nil {return err}

	return nil
}

// Returns all categories
func (dao CategoriaImpl) Get_All() ([]models.Categoria, error) {

	categories := make([]models.Categoria, 0)
	query := "SELECT * FROM categoria"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return categories, err}
	defer stmt.Close()

	rows, err := stmt.Query(); if err != nil {return categories, err}

	for rows.Next() {

		var row models.Categoria

		err := rows.Scan(&row.Id_categoria, &row.Nombre); if err != nil {return categories, err}
		categories = append(categories, row)
	}

	return categories, nil
}

//Returns a category given an id
func (dao CategoriaImpl) Get_By_Id(id int) (models.Categoria, error) {

	var category models.Categoria

	query := "SELECT * FROM categoria WHERE id_categoria = $1"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return category, err}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(&category.Id_categoria, &category.Nombre); if err != nil {return category, err}

	return category, nil
}

//Delete a category given its id
func (dao CategoriaImpl) Delete(id int) error {

	query := "DELETE FROM categoria WHERE id_categoria = $1"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return err}
	defer stmt.Close()

	r, err := stmt.Exec(id); if err != nil {return err}

	i, _ := r.RowsAffected()
	if i != 1 {
		return errors.New("Error, se esperaba una fila afectada")
	}

	return nil
}

// Update a category name
func (dao CategoriaImpl) Update(categoria *models.Categoria) error {

	query := "UPDATE categoria SET nombre = $1 WHERE id_categoria = $2"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return err}
	defer stmt.Close()

	row, err := stmt.Exec(categoria.Nombre, categoria.Id_categoria); if err != nil {return err}
	i, _ := row.RowsAffected()
	if i != 1 {
		return errors.New("Error, se esperaba una fila afectada")
	}

	return nil
}