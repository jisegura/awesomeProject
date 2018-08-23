package postgrsql

import (
	"awesomeProject/models"
	"errors"
	"log"
)

type CategoriaImpl struct{}

//INSERT
func (dao CategoriaImpl) Create(categoria *models.Categoria) error {

	query := "INSERT INTO categoria (nombre) VALUES ($1) RETURNING id_categoria"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	row := stmt.QueryRow(categoria.Nombre)
	err = row.Scan(&categoria.Id_categoria)
	if err != nil {
		return err
	}

	return nil
}

//SELECT ALL
func (dao CategoriaImpl) GetAll() ([]models.Categoria, error) {

	categorias := make([]models.Categoria, 0)
	query := "SELECT * FROM categoria"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return categorias, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return categorias, err
	}

	for rows.Next() {
		var row models.Categoria
		err := rows.Scan(&row.Id_categoria, &row.Nombre)
		if err != nil {
			return categorias, err
		}
		categorias = append(categorias, row)
	}

	return categorias, nil
}

//SELECT BY ID
func (dao CategoriaImpl) GetById(id int) (models.Categoria, error) {

	var p models.Categoria

	query := "SELECT * FROM categoria WHERE id_categoria = $1"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return p, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(&p.Id_categoria, &p.Nombre)
	if err != nil {
		return p, err
	}

	return p, nil
}

//DELETE
func (dao CategoriaImpl) Delete(id int) error {

	query := "DELETE FROM categoria WHERE id_categoria = $1"
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
func (dao CategoriaImpl) Update(categoria *models.Categoria) error {

	query := "UPDATE categoria SET nombre = $1 WHERE id_categoria = $2"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	r, err := stmt.Exec(categoria.Nombre, categoria.Id_categoria)
	if err != nil {
		log.Fatal(err)
	}

	i, _ := r.RowsAffected()
	if i != 1 {
		return errors.New("Error, se esperaba una fila afectada")
	}

	return nil
}
