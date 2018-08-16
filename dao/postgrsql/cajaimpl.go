package postgrsql

import (
	"awesomeProject/models"
)

type CajaImpl struct {}

//INSERT
func (dao CajaImpl)Create(caja *models.Caja) error {

	query := "INSERT INTO caja (inicio, fin) VALUES ($1, $2) RETURNING id_caja"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	row := stmt.QueryRow(caja.Inicio, caja.Fin)
	row.Scan(&caja.Id_caja)
	return nil
}

//SELECT ALL
func (dao CajaImpl) GetAll()([]models.Caja, error) {

	cajas := make([]models.Caja, 0)
	query := "SELECT * FROM caja"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return cajas, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return cajas, err
	}

	for rows.Next() {
		var row models.Caja
		err := rows.Scan(&row.Id_caja, &row.Inicio, &row.Fin)
		if err != nil {
			return cajas, err
		}
		cajas = append(cajas, row)
	}

	return cajas, nil
}
