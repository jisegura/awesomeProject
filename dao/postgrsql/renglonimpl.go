package postgrsql

import "awesomeProject/models"

type RenglonImpl struct {}

//INSERT
func (dao RenglonImpl)Create(renglon *models.Renglon) error {

	query := "INSERT INTO renglon (id_factura, id_producto, cantidad, precio, descuento) VALUES ($1, $2, $3, $4, $5) RETURNING id_renglon"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	row := stmt.QueryRow(renglon.Id_factura, renglon.Id_producto, renglon.Cantidad, renglon.Precio, renglon.Descuento)
	row.Scan(&renglon.Id_renglon)
	return nil
}

//SELECT ALL
func (dao RenglonImpl) GetAll()([]models.Renglon, error) {

	renglones := make([]models.Renglon, 0)
	query := "SELECT * FROM renglon"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return renglones, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return renglones, err
	}

	for rows.Next() {
		var row models.Renglon
		err := rows.Scan(&row.Id_factura, &row.Id_producto, &row.Cantidad, &row.Precio, &row.Descuento)
		if err != nil {
			return renglones, err
		}
		renglones = append(renglones, row)
	}

	return renglones, nil
}
