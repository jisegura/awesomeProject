package postgrsql

import (
	"awesomeProject/models"
	"encoding/json"
	"fmt"
)

type RenglonImpl struct{}

type Renglones struct {
}

//INSERT
func InsertRenglones(renglones []models.Renglon, id int) error {

	for i := range renglones {
		renglones[i].Id_factura = id
		err := create(&renglones[i])
		if err != nil {
			return err
		}
	}

	return nil
}

//CREATE
func create(renglon *models.Renglon) error {

	query := "INSERT INTO renglon (id_producto, id_factura, cantidad, precio, descuento) VALUES ($1, $2, $3, $4, $5) RETURNING id_renglon"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	row := stmt.QueryRow(renglon.Id_producto, renglon.Id_factura, renglon.Cantidad, renglon.Precio, renglon.Descuento)
	row.Scan(&renglon.Id_renglon)

	return nil
}

//SELECT ALL
func GetAll(id int) ([]models.Renglon, error) {

	renglones := make([]models.Renglon, 0)
	query := "SELECT * FROM renglon WHERE id_factura = $1"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return renglones, err
	}

	rows, err := stmt.Query(id)
	if err != nil {
		return renglones, err
	}

	for rows.Next() {
		var row models.Renglon
		fmt.Print(row)
		err := rows.Scan(&row.Id_factura, &row.Id_renglon, &row.Id_producto, &row.Cantidad, &row.Precio, &row.Descuento)
		if err != nil {
			return renglones, err
		}

		renglones = append(renglones, row)
	}
	json.Marshal(renglones)
	return renglones, nil
}
