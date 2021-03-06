package postgrsql

import (
	"awesomeProject/models"
	"database/sql"
	"errors"
	"time"
)

type CajaImpl struct{}

//INSERT
func (dao CajaImpl) Create(caja *models.Caja) (models.Caja, error) {

	var newCaja models.Caja

	cAbierta, err := GetCajaAbierta()

	if cAbierta.Id_caja != 0 {
		return cAbierta, nil
	}

	query := "INSERT INTO caja (inicio, fin, horaInicio, horaFin) VALUES ($1, $2, $3, $4) RETURNING id_caja"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return newCaja, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(caja.Inicio, 0, time.Now(), time.Time{})
	err = row.Scan(&caja.Id_caja)
	if err != nil {
		return newCaja, err
	}

	newCaja, err = GetById(caja.Id_caja)
	if err != nil {
		return newCaja, err
	}

	return newCaja, nil
}

func (dao CajaImpl) GetCaja() (models.Caja, error) {

	return GetCajaAbierta()
}

func GetCajaAbierta() (models.Caja, error) {

	var caja models.Caja
	query := "SELECT * FROM caja WHERE fin = 0"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return caja, err
	}
	defer stmt.Close()

	row := stmt.QueryRow()
	err = row.Scan(&caja.Id_caja, &caja.Inicio, &caja.Fin, &caja.HoraInicio, &caja.HoraFin)
	if err != nil && err != sql.ErrNoRows {
		return caja, err
	}

	return caja, nil
}

func GetById(id int) (models.Caja, error) {

	var caja models.Caja

	query := "SELECT * FROM caja WHERE id_caja = $1"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return caja, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(&caja.Id_caja, &caja.Inicio, &caja.Fin, &caja.HoraInicio, &caja.HoraFin)
	if err != nil && err != sql.ErrNoRows {
		return caja, err
	}

	return caja, err
}

//SELECT ALL
func (dao CajaImpl) GetAll() ([]models.Caja, error) {

	cajas := make([]models.Caja, 0)
	query := "SELECT * FROM caja WHERE fin != 0"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return cajas, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return cajas, err
	}

	for rows.Next() {
		var row models.Caja
		err := rows.Scan(&row.Id_caja, &row.Inicio, &row.Fin, &row.HoraInicio, &row.HoraFin)
		if err != nil {
			return cajas, err
		}
		cajas = append(cajas, row)
	}

	return cajas, nil
}

//UPADTE CIERRE CAJA
func (dao CajaImpl) CierreCaja(caja *models.Caja) (models.Caja, error) {

	query := "UPDATE caja SET fin = $1, horaFin = $2 WHERE id_caja = $3"
	db := getConnection()
	defer db.Close()

	var c models.Caja

	stmt, err := db.Prepare(query)
	if err != nil {
		return c, err
	}
	defer stmt.Close()

	row, err := stmt.Exec(caja.Fin, time.Now(), caja.Id_caja)
	if err != nil {
		return c, err
	}

	i, _ := row.RowsAffected()
	if i != 1 {
		return c, errors.New("Error, se esperaba una fila afectada")
	}

	c, err = GetById(caja.Id_caja)
	if err != nil {
		return c, err
	}

	return c, nil
}
