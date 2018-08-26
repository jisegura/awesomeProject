package postgrsql

import (
	"awesomeProject/models"
	"errors"
	"fmt"
	_ "github.com/teepark/pqinterval"
	"time"
)

type FacturaImpl struct{}

//PUBLIC
func (dao FacturaImpl) Create(factura *models.Factura) error {
	return InsertFactura(factura)
}

func (dao FacturaImpl) CreateCliente(factura *models.Factura) (models.Factura, error) {

	newFactura, err := InsertCliente(factura)
	return newFactura, err
}

func (dao FacturaImpl) CreateOtro(factura *models.Factura) (models.Factura, error) {

	newFactura, err := InsertOtros(factura)
	return newFactura, err
}

func (dao FacturaImpl) GetAll(id int) ([]models.Factura, error) {

	facturas, err := GetAllFacturas(id)
	return facturas, err
}

func (dao FacturaImpl) GetFacturasCliente(id int) ([]models.Factura, error) {

	facturas, err := GetAllClientes(id)
	return facturas, err
}

func (dao FacturaImpl) GetFacturasOtro(id int) ([]models.Factura, error) {

	facturas, err := GetAllOtros(id)
	return facturas, err
}

func (dao FacturaImpl) GetById(id int) (models.Factura, error) {

	factura, err := GetRetiroById(id)
	return factura, err
}

func (dao FacturaImpl) GetByIdCliente(id int) (models.Factura, error) {

	factura, err := GetClienteById(id)
	return factura, err
}

func (dao FacturaImpl) GetByIdOtros(id int) (models.Factura, error) {

	factura, err := GetOtrosById(id)
	return factura, err
}

func (dao FacturaImpl) Update(factura *models.Factura) error {
	return UpdateComentario(factura)
}

func (dao FacturaImpl) GetFacturasEliminadas() ([]models.Factura, error) {

	facturas, err := GetFacturasEliminadas()
	return facturas, err
}

func (dao FacturaImpl) GetFacturasById(id int) ([]int, error) {

	var ids []int

	query := "SELECT id_factura FROM factura WHERE id_caja = $1"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return ids, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return ids, err
	}

	for rows.Next() {
		var row models.Factura

		err = rows.Scan(&row.Id_factura)
		if err != nil {
			return ids, err
		}

		ids = append(ids, row.Id_factura)
	}

	return ids, nil
}

func (dao FacturaImpl) GetAllFacturasById(id int) ([]models.Factura, error) {

	query := "SELECT g.id_factura, id_caja, id_empleado, fecha, precio, comentarioBaja, descuento, formaDePago, comentario FROM otros o RIGHT JOIN" +
		"(SELECT f.id_factura, id_caja, id_empleado, fecha, precio, comentarioBaja, descuento, formaDePago FROM cliente c RIGHT JOIN " +
		"(SELECT * FROM factura WHERE id_caja = $1 ORDER BY fecha DESC) f ON c.id_factura = f.id_factura) g " +
		"ON o.id_factura = g.id_factura"

	db := getConnection()
	defer db.Close()

	var facturas []models.Factura

	stmt, err := db.Prepare(query)
	if err != nil {
		return facturas, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return facturas, err
	}

	for rows.Next() {
		var row models.Factura

		err = rows.Scan(&row.Id_factura, &row.Id_caja, &row.Id_empleado, &row.Fecha, &row.Precio, &row.ComentarioBaja, &row.Descuento, &row.FormaDePago, &row.Comentario)
		if err != nil {
			return facturas, err
		}

		facturas = append(facturas, row)
	}

	return facturas, nil
}

////////////////////////////////////////////////////////////
//PRIVATE

//INSERT RETIRO
func InsertFactura(factura *models.Factura) error {

	query := "INSERT INTO factura (id_caja, id_empleado, fecha, precio, comentarioBaja) VALUES ($1, $2, $3, $4, $5) RETURNING id_factura"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	factura.Fecha = time.Now()
	row := stmt.QueryRow(factura.Id_caja, factura.Id_empleado, factura.Fecha, factura.Precio, factura.ComentarioBaja)
	err = row.Scan(&factura.Id_factura)
	if err != nil {
		return err
	}

	return nil
}

//INSERT CLIENTE
func InsertCliente(factura *models.Factura) (models.Factura, error) {

	query := "WITH X AS" +
		"(INSERT INTO factura (id_caja, id_empleado, fecha, precio, comentarioBaja) VALUES ($1, $2, $3, $4, $5) RETURNING id_factura)" +
		"INSERT INTO cliente (id_factura, descuento, formaDePago) VALUES ((SELECT id_factura FROM X), $6, $7) RETURNING id_factura"

	db := getConnection()
	defer db.Close()

	var newFactura models.Factura

	stmt, err := db.Prepare(query)
	if err != nil {
		return newFactura, err
	}
	defer stmt.Close()

	factura.Fecha = time.Now()
	row := stmt.QueryRow(factura.Id_caja, factura.Id_empleado, factura.Fecha, factura.Precio, factura.ComentarioBaja, factura.Descuento, factura.FormaDePago)
	err = row.Scan(&factura.Id_factura)
	if err != nil {
		return newFactura, err
	}

	err = InsertRenglones(factura.Renglones, factura.Id_factura)
	if err != nil {
		return newFactura, err
	}

	newFactura, err = GetClienteById(factura.Id_factura)
	if err != nil {
		return newFactura, err
	}

	return newFactura, nil
}

//INSERT OTRO
func InsertOtros(factura *models.Factura) (models.Factura, error) {

	query := "WITH X AS" +
		"(INSERT INTO factura (id_caja, id_empleado, fecha, precio, comentarioBaja) VALUES ($1, $2, $3, $4, $5) RETURNING id_factura)" +
		"INSERT INTO otros (id_factura, comentario) VALUES ((SELECT id_factura FROM X), $6) RETURNING id_factura"

	db := getConnection()
	defer db.Close()

	var newFactura models.Factura

	stmt, err := db.Prepare(query)
	if err != nil {
		return newFactura, err
	}
	defer stmt.Close()

	factura.Fecha = time.Now()
	row := stmt.QueryRow(factura.Id_caja, factura.Id_empleado, factura.Fecha, factura.Precio, factura.ComentarioBaja, factura.Comentario)
	err = row.Scan(&factura.Id_factura)
	if err != nil {
		return newFactura, err
	}

	newFactura, err = GetOtrosById(factura.Id_factura)
	if err != nil {
		return newFactura, err
	}

	return newFactura, nil
}

//SELECT ALL RETIROS
func GetAllFacturas(id int) ([]models.Factura, error) {

	facturas := make([]models.Factura, 0)
	query := "SELECT * FROM factura " +
		"WHERE id_factura NOT IN (SELECT id_factura FROM cliente) " +
		"AND id_factura NOT IN (SELECT id_factura FROM otros) " +
		"AND id_caja = $1"

	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return facturas, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return facturas, err
	}

	for rows.Next() {
		var row models.Factura
		err := rows.Scan(&row.Id_factura, &row.Id_caja, &row.Id_empleado, &row.Fecha, &row.Precio, &row.ComentarioBaja)
		if err != nil {
			return facturas, err
		}

		facturas = append(facturas, row)
	}

	return facturas, nil
}

//SELECT  FROM CLIENTE FOR CAJA
func GetAllClientes(id int) ([]models.Factura, error) {

	facturas := make([]models.Factura, 0)
	query := "SELECT c.id_factura, id_caja, id_empleado, fecha, precio, comentarioBaja, descuento, formaDePago FROM factura f INNER JOIN Cliente c ON f.id_factura = c.id_factura WHERE id_caja = $1"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return facturas, err
	}

	rows, err := stmt.Query(id)
	if err != nil {
		return facturas, err
	}

	for rows.Next() {
		var row models.Factura
		err := rows.Scan(&row.Id_factura, &row.Id_caja, &row.Id_empleado, &row.Fecha, &row.Precio, &row.ComentarioBaja, &row.Descuento, &row.FormaDePago)
		if err != nil {
			return facturas, err
		}

		row.Renglones, err = GetAll(row.Id_factura)
		if err != nil {
			return facturas, err
		}
		facturas = append(facturas, row)
	}

	return facturas, nil
}

//SELECT ALL FROM OTRO FOR CAJA
func GetAllOtros(id int) ([]models.Factura, error) {

	facturas := make([]models.Factura, 0)
	query := "SELECT o.id_factura, id_caja, id_empleado, fecha, precio, comentarioBaja, comentario FROM factura f INNER JOIN otros o ON f.id_factura = o.id_factura WHERE id_caja =$1"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return facturas, err
	}

	rows, err := stmt.Query(id)
	if err != nil {
		return facturas, err
	}

	for rows.Next() {
		var row models.Factura
		err := rows.Scan(&row.Id_factura, &row.Id_caja, &row.Id_empleado, &row.Fecha, &row.Precio, &row.ComentarioBaja, &row.Comentario)
		if err != nil {
			return facturas, err
		}

		facturas = append(facturas, row)

	}

	return facturas, nil
}

//SELECT BY ID
func GetRetiroById(id int) (models.Factura, error) {

	var factura models.Factura

	query := "SELECT * FROM factura WHERE id_factura = $1"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return factura, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(&factura.Id_factura, &factura.Id_caja, &factura.Id_empleado, &factura.Fecha, &factura.Precio, &factura.ComentarioBaja)
	if err != nil {
		return factura, err
	}

	return factura, nil
}

//SELECT BY ID CLIENTE
func GetClienteById(id int) (models.Factura, error) {

	var factura models.Factura

	query := "SELECT c.id_factura, id_caja, id_empleado, fecha, precio, comentarioBaja, descuento, formaDePago FROM Factura f INNER JOIN Cliente c ON f.id_factura = c.id_factura WHERE c.id_factura = $1"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return factura, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(&factura.Id_factura, &factura.Id_caja, &factura.Id_empleado, &factura.Fecha, &factura.Precio, &factura.ComentarioBaja, &factura.Descuento, &factura.FormaDePago)
	if err != nil {
		return factura, err
	}

	factura.Renglones, err = GetAll(factura.Id_factura)
	if err != nil {
		return factura, err
	}

	return factura, nil
}

//SELECT BY ID OTROS
func GetOtrosById(id int) (models.Factura, error) {

	var factura models.Factura

	query := "SELECT o.id_factura, id_caja, id_empleado, fecha, precio, comentarioBaja, comentario FROM factura f INNER JOIN otros o ON f.id_factura = o.id_factura WHERE o.id_factura = $1"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return factura, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(&factura.Id_factura, &factura.Id_caja, &factura.Id_empleado, &factura.Fecha, &factura.Precio, &factura.ComentarioBaja, &factura.Comentario)
	if err != nil {
		return factura, err
	}

	return factura, nil
}

//UPDATE COMENTARIO
func UpdateComentario(factura *models.Factura) error {

	query := "UPDATE factura SET comentarioBaja = $1 WHERE id_factura = $2"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	fmt.Print(factura)
	row, err := stmt.Exec(factura.ComentarioBaja, factura.Id_factura)
	if err != nil {
		return err
	}

	i, _ := row.RowsAffected()
	if i != 1 {
		return errors.New("Error, se esperaba una fila afectada")
	}

	return nil
}

func GetFacturasEliminadas() ([]models.Factura, error) {

	var facturas []models.Factura
	query := "SELECT g.id_factura, id_caja, id_empleado, fecha, precio, comentarioBaja, descuento, formaDePago, comentario FROM otros o RIGHT JOIN" +
		"(SELECT f.id_factura, id_caja, id_empleado, fecha, precio, comentarioBaja, descuento, formaDePago FROM cliente c RIGHT JOIN " +
		"(SELECT * FROM factura WHERE comentarioBaja  NOT LIKE '') f ON c.id_factura = f.id_factura) g " +
		"ON o.id_factura = g.id_factura"

	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return facturas, err
	}
	defer stmt.Close()

	rows, err := db.Query(query)
	if err != nil {
		return facturas, err
	}

	for rows.Next() {
		var row models.Factura
		err := rows.Scan(&row.Id_factura, &row.Id_caja, &row.Id_empleado, &row.Fecha, &row.Precio, &row.ComentarioBaja, &row.Descuento, &row.FormaDePago, &row.Comentario)
		if err != nil {
			return facturas, err
		}

		row.Renglones, err = GetAll(row.Id_factura)
		if err != nil {
			return facturas, err
		}

		facturas = append(facturas, row)

	}

	return facturas, nil
}

func (dao FacturaImpl) GetLastFacturas(id int) ([]models.Factura, error) {

	query := "SELECT g.id_factura, id_caja, id_empleado, fecha, precio, comentarioBaja, descuento, formaDePago, comentario FROM otros o RIGHT JOIN" +
		"(SELECT f.id_factura, id_caja, id_empleado, fecha, precio, comentarioBaja, descuento, formaDePago FROM cliente c RIGHT JOIN " +
		"(SELECT * FROM factura WHERE id_caja = $1  AND comentarioBaja LIKE '' ORDER BY fecha DESC LIMIT 5) f ON c.id_factura = f.id_factura) g " +
		"ON o.id_factura = g.id_factura"

	db := getConnection()
	defer db.Close()

	var facturas []models.Factura
	stmt, err := db.Prepare(query)
	if err != nil {
		return facturas, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return facturas, err
	}

	for rows.Next() {
		var factura models.Factura

		err = rows.Scan(&factura.Id_factura, &factura.Id_caja, &factura.Id_empleado, &factura.Fecha, &factura.Precio, &factura.ComentarioBaja, &factura.Descuento, &factura.FormaDePago, &factura.Comentario)
		if err != nil {
			return facturas, err
		}

		factura.Renglones, err = GetAll(factura.Id_factura)
		if err != nil {
			return facturas, err
		}

		facturas = append(facturas, factura)
	}

	return facturas, nil
}
