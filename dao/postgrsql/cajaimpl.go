package postgrsql

import (
	"awesomeProject/models"
	"database/sql"
	"errors"
	"time"
)

type CajaImpl struct{}

//Allows create a new cash box
func (dao CajaImpl) Create (cashBox *models.Caja) (models.Caja, error) {

	var newCashBox models.Caja

	open, err := dao.Get_Cash_Box()

	if open.Id_caja != 0 {return open, nil}

	query := "INSERT INTO caja (inicio, fin, horaInicio, horaFin, cierreReal, cierreFiscal) " +
			 "VALUES ($1, $2, $3, $4, $5, $6) " +
			 "RETURNING id_caja"

	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return newCashBox, err}
	defer stmt.Close()

	row := stmt.QueryRow(cashBox.Inicio, 0, time.Now(), time.Time{}, 0, 0)
	err = row.Scan(&cashBox.Id_caja); if err != nil {return newCashBox, err}

	newCashBox, err = dao.Get_By_Id(cashBox.Id_caja); if err != nil {return newCashBox, err}

	return newCashBox, nil
}
//Return the open cash box
func (dao CajaImpl) Get_Cash_Box() (models.Caja, error) {

	var cashBox models.Caja

	query := "SELECT * FROM caja WHERE fin = 0"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return cashBox, err}
	defer stmt.Close()

	row := stmt.QueryRow()
	err = row.Scan(&cashBox.Id_caja, &cashBox.Inicio, &cashBox.Fin, &cashBox.HoraInicio, &cashBox.HoraFin, &cashBox.CierreReal, &cashBox.CierreFiscal)
	if err != nil && err != sql.ErrNoRows {return cashBox, err}

	return cashBox, nil
}

//Returns a cash box given an id
func (dao CajaImpl) Get_By_Id(id int) (models.Caja, error) {

	var cashBox models.Caja

	query := "SELECT * FROM caja WHERE id_caja = $1"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return cashBox, err}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(&cashBox.Id_caja, &cashBox.Inicio, &cashBox.Fin, &cashBox.HoraInicio, &cashBox.HoraFin, &cashBox.CierreReal, &cashBox.CierreFiscal)
	if err != nil && err != sql.ErrNoRows {return cashBox, err}

	return cashBox, err
}

//Returns all cash boxes
func (dao CajaImpl) Get_All() ([]models.Caja, error) {

	cashBoxes := make([]models.Caja, 0)
	query := "SELECT * FROM caja WHERE fin != 0"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return cashBoxes, err}
	defer stmt.Close()

	rows, err := stmt.Query(); if err != nil {return cashBoxes, err}
	for rows.Next() {

		var row models.Caja

		err := rows.Scan(&row.Id_caja, &row.Inicio, &row.Fin, &row.HoraInicio, &row.HoraFin, &row.CierreReal, &row.CierreFiscal); if err != nil {return cashBoxes, err}
		cashBoxes = append(cashBoxes, row)
	}

	return cashBoxes, nil
}

//Close cash box
func (dao CajaImpl) Close_Cash_Box (cashBox *models.Caja) (models.Caja, error) {

	query := "UPDATE caja SET fin = $1, horaFin = $2, cierreReal = $3, cierreFiscal = $4 WHERE id_caja = $5"
	db := getConnection()
	defer db.Close()

	var cashBoxAux models.Caja

	stmt, err := db.Prepare(query); if err != nil {return cashBoxAux, err}
	defer stmt.Close()

	row, err := stmt.Exec(cashBox.Fin, time.Now(), cashBox.CierreReal, cashBox.CierreFiscal, cashBox.Id_caja); if err != nil {return cashBoxAux, err}
	i, _ := row.RowsAffected(); if i != 1 {return cashBoxAux, errors.New("Error, se esperaba una fila afectada")}

	return dao.Get_By_Id(cashBox.Id_caja)
}

//Returns cash boxes by date
func (dao CajaImpl) Get_By_Date(startDate time.Time, finalDate time.Time) ([]models.Caja, error) {

	query := "SELECT * FROM caja where horaInicio BETWEEN $1 AND $2"
	db := getConnection()
	defer db.Close()

	var cashBoxes []models.Caja
	stmt, err := db.Prepare(query); if err != nil {return cashBoxes, err}
	defer stmt.Close()

	rows, err := stmt.Query(startDate, finalDate); if err != nil {return cashBoxes, err}

	for rows.Next() {

		var caja models.Caja

		err := rows.Scan(&caja.Id_caja, &caja.Inicio, &caja.Fin, &caja.HoraInicio, &caja.HoraFin, &caja.CierreReal, &caja.CierreFiscal); if err != nil {return cashBoxes, err}
		cashBoxes = append(cashBoxes, caja)
	}

	return cashBoxes, nil
}