package postgrsql

import (
	"awesomeProject/models"
	"errors"
	"log"
)

type ProductoImpl struct{}

//INSERT
func (dao ProductoImpl) Create(producto *models.Producto) error {

	query := "INSERT INTO producto (id_categoria, nombre, precio, imagen) VALUES ($1, $2, $3, $4) RETURNING id_producto"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	row := stmt.QueryRow(producto.Id_categoria, producto.Nombre, producto.Precio, producto.Imagen)
	row.Scan(&producto.Id_producto)
	return nil
}

//SELECT ALL
func (dao ProductoImpl) GetAll() ([]models.Producto, error) {

	productos := make([]models.Producto, 0)
	query := "SELECT * FROM producto"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return productos, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return productos, err
	}

	for rows.Next() {
		var row models.Producto
		err := rows.Scan(&row.Id_producto, &row.Id_categoria, &row.Nombre, &row.Precio, &row.Imagen)
		if err != nil {
			return productos, err
		}
		productos = append(productos, row)
	}

	return productos, nil
}

//SELECT BY ID
func (dao ProductoImpl) GetById(id int) (models.Producto, error) {

	var p models.Producto

	query := "SELECT * FROM producto WHERE id_producto = $1"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return p, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(&p.Id_producto, &p.Id_categoria, &p.Nombre, &p.Precio, &p.Imagen)
	if err != nil {
		return p, err
	}

	return p, nil
}

func GetNombreById(id int64) (string, error) {

	query := "SELECT nombre FROM producto WHERE id_producto = $1"
	db := getConnection()
	defer db.Close()

	var nombre string

	stmt, err := db.Prepare(query)
	if err != nil {
		return nombre, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(&nombre)
	if err != nil {
		return nombre, err
	}

	return nombre, nil
}

//DELETE
func (dao ProductoImpl) Delete(id int) error {

	query := "DELETE FROM producto WHERE id_producto = $1"
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
func (dao ProductoImpl) Update(producto *models.Producto) error {

	query := "UPDATE producto SET id_categoria = $1, nombre = $2, precio = $3, imagen = $4 WHERE id_producto = $5"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	row, err := stmt.Exec(producto.Id_categoria, producto.Nombre, producto.Precio, producto.Imagen, producto.Id_producto)
	if err != nil {
		log.Fatal(err)
	}

	i, _ := row.RowsAffected()
	if i != 1 {
		return errors.New("Error, se esperaba una fila afectada")
	}

	return nil
}
