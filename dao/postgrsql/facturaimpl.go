package postgrsql

import (
	"awesomeProject/models"
)

type FacturaImpl struct {}

//INSERT
func (dao FacturaImpl)Create(factura *models.Factura) error {

	query := "INSERT INTO factura (id_caja, descuento, precio, activa, fecha) VALUES ($1, $2, $3, $4, $5) RETURNING id_factura"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	row := stmt.QueryRow(factura.Id_caja, factura.Descuento, factura.Precio, factura.Activa, factura.Fecha)
	row.Scan(&factura.Id_factura)
	return nil
}

//SELECT ALL
func (dao FacturaImpl) GetAll()([]models.Factura, error) {

	facturas := make([]models.Factura, 0)
	query := "SELECT * FROM factura"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return facturas, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return facturas, err
	}

	for rows.Next() {
		var row models.Factura
		err := rows.Scan(&row.Id_factura, &row.Id_caja, &row.Descuento, &row.Fecha, &row.Precio, &row.Activa)
		if err != nil {
			return facturas, err
		}
		facturas = append(facturas, row)
	}

	return facturas, nil
}
