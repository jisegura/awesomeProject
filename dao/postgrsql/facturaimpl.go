package postgrsql

import (
	"awesomeProject/models"
	"errors"
	_ "github.com/teepark/pqinterval"
	"strings"
	"time"
)

type FacturaImpl struct{}

//Allows to enter a new invoice
func (dao FacturaImpl) Create(invoice *models.Factura) error {

	query := "INSERT INTO factura (id_caja, id_empleado, fecha, precio, comentarioBaja) " +
			 "VALUES ($1, $2, $3, $4, $5) " +
			 "RETURNING id_factura"

	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return err}
	defer stmt.Close()

	row := stmt.QueryRow(invoice.Id_caja, invoice.Id_empleado, time.Now(), invoice.Precio, invoice.ComentarioBaja)
	err = row.Scan(&invoice.Id_factura); if err != nil {return err}

	return nil
}

//Allows to enter a new client
func (dao FacturaImpl) Create_Customer(invoice *models.Factura) (models.Factura, error) {

	var newInvoice models.Factura

	query := "WITH X AS" +
		"(INSERT INTO factura (id_caja, id_empleado, fecha, precio, comentarioBaja) VALUES ($1, $2, $3, $4, $5) RETURNING id_factura)" +
		"INSERT INTO cliente (id_factura, descuento, formaDePago) VALUES ((SELECT id_factura FROM X), $6, $7) RETURNING id_factura"

	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return newInvoice, err}
	defer stmt.Close()

	invoice.Fecha = time.Now()
	row := stmt.QueryRow(invoice.Id_caja, invoice.Id_empleado, invoice.Fecha, invoice.Precio, invoice.ComentarioBaja, invoice.Descuento, invoice.FormaDePago)
	err = row.Scan(&invoice.Id_factura); if err != nil {

		if strings.Contains(err.Error(), "fk_empleado") {
			return newInvoice, errors.New("Error, no hay empleado seleccionado")
		}
		return newInvoice, errors.New("Error, no hay caja abierta")
	}

	err = InsertRenglones(invoice.Renglones, invoice.Id_factura); if err != nil {return newInvoice, err}

	newInvoice, err = dao.Get_By_Customer_Id(invoice.Id_factura); if err != nil {return newInvoice, err}

	return newInvoice, nil
}

//Allows to enter a new invoice (type 'otros')
func (dao FacturaImpl) Create_Other(invoice *models.Factura) (models.Factura, error) {

	var newInvoice models.Factura

	query := "WITH X AS" +
		"(INSERT INTO factura (id_caja, id_empleado, fecha, precio, comentarioBaja) VALUES ($1, $2, $3, $4, $5) RETURNING id_factura)" +
		"INSERT INTO otros (id_factura, comentario) VALUES ((SELECT id_factura FROM X), $6) RETURNING id_factura"

	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return newInvoice, err}
	defer stmt.Close()

	invoice.Fecha = time.Now()
	row := stmt.QueryRow(invoice.Id_caja, invoice.Id_empleado, invoice.Fecha, invoice.Precio, invoice.ComentarioBaja, invoice.Comentario)
	err = row.Scan(&invoice.Id_factura); if err != nil {return newInvoice, err}

	_ , err = dao.Get_Other_Invoices(invoice.Id_factura); if err != nil {return newInvoice, err}

	return newInvoice, nil
}

//Returns all invoices
func (dao FacturaImpl) Get_All(id int) ([]models.Factura, error) {

	invoice := make([]models.Factura, 0)

	query := "SELECT * FROM factura " +
		"WHERE id_factura NOT IN (SELECT id_factura FROM cliente) " +
		"AND id_factura NOT IN (SELECT id_factura FROM otros) " +
		"AND id_caja = $1"

	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return invoice, err}
	defer stmt.Close()

	rows, err := stmt.Query(id); if err != nil {return invoice, err}

	for rows.Next() {
		var row models.Factura

		err := rows.Scan(&row.Id_factura, &row.Id_caja, &row.Id_empleado, &row.Fecha, &row.Precio, &row.ComentarioBaja); if err != nil {return invoice, err}
		invoice = append(invoice, row)
	}

	return invoice, nil
}

//Returns all client invoices
func (dao FacturaImpl) Get_Customer_Invoices(id int) ([]models.Factura, error) {

	invoices:= make([]models.Factura, 0)

	query := "SELECT c.id_factura, id_caja, id_empleado, fecha, precio, comentarioBaja, descuento, formaDePago FROM factura f INNER JOIN Cliente c ON f.id_factura = c.id_factura WHERE id_caja = $1"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return invoices, err}

	rows, err := stmt.Query(id); if err != nil {return invoices, err}

	for rows.Next() {
		var row models.Factura

		err := rows.Scan(&row.Id_factura, &row.Id_caja, &row.Id_empleado, &row.Fecha, &row.Precio, &row.ComentarioBaja, &row.Descuento, &row.FormaDePago); if err != nil {return invoices, err}
		row.Renglones, err = GetAll(row.Id_factura); if err != nil {return invoices, err}
		invoices = append(invoices, row)
	}

	return invoices, nil
}

//Returns an invoice given an id
func (dao FacturaImpl) Get_Other_Invoices(id int) ([]models.Factura, error) {

	invoices := make([]models.Factura, 0)

	query := "SELECT o.id_factura, id_caja, id_empleado, fecha, precio, comentarioBaja, comentario FROM factura f INNER JOIN otros o ON f.id_factura = o.id_factura WHERE id_caja =$1"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return invoices, err}

	rows, err := stmt.Query(id); if err != nil {return invoices, err}

	for rows.Next() {
		var row models.Factura

		err := rows.Scan(&row.Id_factura, &row.Id_caja, &row.Id_empleado, &row.Fecha, &row.Precio, &row.ComentarioBaja, &row.Comentario); if err != nil {return invoices, err}
		invoices = append(invoices, row)
	}

	return invoices, nil
}

//Returns an invoices given an id (type 'recibo')
func (dao FacturaImpl) Get_By_Id(id int) (models.Factura, error) {

	var invoice models.Factura

	query := "SELECT * FROM factura WHERE id_factura = $1"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return invoice, err}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(&invoice.Id_factura, &invoice.Id_caja, &invoice.Id_empleado, &invoice.Fecha, &invoice.Precio, &invoice.ComentarioBaja); if err != nil {return invoice, err}

	return invoice, nil
}

//Returns a client given an id
func (dao FacturaImpl) Get_By_Customer_Id(id int) (models.Factura, error) {

	var invoice models.Factura

	query := "SELECT c.id_factura, id_caja, id_empleado, fecha, precio, comentarioBaja, descuento, formaDePago FROM Factura f INNER JOIN Cliente c ON f.id_factura = c.id_factura WHERE c.id_factura = $1"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return invoice, err}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(&invoice.Id_factura, &invoice.Id_caja, &invoice.Id_empleado, &invoice.Fecha, &invoice.Precio, &invoice.ComentarioBaja, &invoice.Descuento, &invoice.FormaDePago); if err != nil {return invoice, err}

	invoice.Renglones, err = GetAll(invoice.Id_factura); if err != nil {return invoice, err}

	return invoice, nil
}

//Returns an invoice given an id (type 'Otros')
func (dao FacturaImpl) Get_By_Other_Id(id int) (models.Factura, error) {

	var invoice models.Factura

	query := "SELECT o.id_factura, id_caja, id_empleado, fecha, precio, comentarioBaja, comentario FROM factura f INNER JOIN otros o ON f.id_factura = o.id_factura WHERE o.id_factura = $1"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return invoice, err}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(&invoice.Id_factura, &invoice.Id_caja, &invoice.Id_empleado, &invoice.Fecha, &invoice.Precio, &invoice.ComentarioBaja, &invoice.Comentario); if err != nil {return invoice, err}

	return invoice, nil
}

//Allows to update the comment of an invoice
func (dao FacturaImpl) Update_Comment(invoice *models.Factura) error {

	query := "UPDATE factura SET comentarioBaja = $1 WHERE id_factura = $2"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return err}
	defer stmt.Close()

	row, err := stmt.Exec(invoice.ComentarioBaja, invoice.Id_factura); if err != nil {return err}

	i, _ := row.RowsAffected()
	if i != 1 {
		return errors.New("Error, se esperaba una fila afectada")
	}

	return nil
}

//Allows to view deleted invoices
func (dao FacturaImpl) Get_Deleted_Invoices() ([]models.Factura, error) {

	var invoices []models.Factura

	query := "SELECT g.id_factura, id_caja, id_empleado, fecha, precio, comentarioBaja, descuento, formaDePago, comentario FROM otros o RIGHT JOIN" +
		"(SELECT f.id_factura, id_caja, id_empleado, fecha, precio, comentarioBaja, descuento, formaDePago FROM cliente c RIGHT JOIN " +
		"(SELECT * FROM factura WHERE comentarioBaja  NOT LIKE '') f ON c.id_factura = f.id_factura) g " +
		"ON o.id_factura = g.id_factura"

	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return invoices, err}
	defer stmt.Close()

	rows, err := db.Query(query); if err != nil {return invoices, err}

	for rows.Next() {
		var row models.Factura

		err := rows.Scan(&row.Id_factura, &row.Id_caja, &row.Id_empleado, &row.Fecha, &row.Precio, &row.ComentarioBaja, &row.Descuento, &row.FormaDePago, &row.Comentario); if err != nil {return invoices, err}
		row.Renglones, err = GetAll(row.Id_factura); if err != nil {return invoices, err}
		invoices = append(invoices, row)
	}

	return invoices, nil
}

//Returns id of the invoices given an cash box
func (dao FacturaImpl) Get_Invoice_Id_By_Cash_Box(id int) ([]int, error) {

	var ids []int

	query := "SELECT id_factura FROM factura WHERE id_caja = $1"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return ids, err}
	defer stmt.Close()

	rows, err := stmt.Query(id); if err != nil {return ids, err}

	for rows.Next() {
		var row models.Factura

		err = rows.Scan(&row.Id_factura); if err != nil{return ids, err}
		ids = append(ids, row.Id_factura)
	}

	return ids, nil
}

//Returns all invoces given an cash box
func (dao FacturaImpl) Get_All_Ivloices_By_Cash_Box(id int) ([]models.Factura, error) {

	query := "SELECT g.id_factura, id_caja, id_empleado, fecha, precio, comentarioBaja, descuento, formaDePago, comentario FROM otros o RIGHT JOIN" +
		     "(SELECT f.id_factura, id_caja, id_empleado, fecha, precio, comentarioBaja, descuento, formaDePago FROM cliente c RIGHT JOIN " +
		     "(SELECT * FROM factura WHERE id_caja = $1 ORDER BY fecha DESC) f ON c.id_factura = f.id_factura) g " +
		     "ON o.id_factura = g.id_factura"

	db := getConnection()
	defer db.Close()

	var invoices []models.Factura

	stmt, err := db.Prepare(query); if err != nil {return invoices, err}
	defer stmt.Close()

	rows, err := stmt.Query(id); if err != nil {return invoices, err}

	for rows.Next() {
		var row models.Factura

		err = rows.Scan(&row.Id_factura, &row.Id_caja, &row.Id_empleado, &row.Fecha, &row.Precio, &row.ComentarioBaja, &row.Descuento, &row.FormaDePago, &row.Comentario); if err != nil {return invoices, err}
		row.Renglones, err = GetAll(row.Id_factura); if err != nil {return invoices, err}
		invoices = append(invoices, row)
	}

	return invoices, nil
}

//Return the last 5 invoices
func (dao FacturaImpl) Get_Last_Invoices(id int) ([]models.Factura, error) {

	var invoices []models.Factura

	query := "SELECT g.id_factura, id_caja, id_empleado, fecha, precio, comentarioBaja, descuento, formaDePago, comentario FROM otros o RIGHT JOIN" +
		     "(SELECT f.id_factura, id_caja, id_empleado, fecha, precio, comentarioBaja, descuento, formaDePago FROM cliente c RIGHT JOIN " +
		     "(SELECT * FROM factura WHERE id_caja = $1 ORDER BY fecha DESC LIMIT 5) f ON c.id_factura = f.id_factura) g " +
		     "ON o.id_factura = g.id_factura"

	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return invoices, err}
	defer stmt.Close()

	rows, err := stmt.Query(id); if err != nil {return invoices, err}

	for rows.Next() {
		var invoice models.Factura

		err = rows.Scan(&invoice.Id_factura, &invoice.Id_caja, &invoice.Id_empleado, &invoice.Fecha, &invoice.Precio, &invoice.ComentarioBaja, &invoice.Descuento, &invoice.FormaDePago, &invoice.Comentario); if err != nil {return invoices, err}
		invoice.Renglones, err = GetAll(invoice.Id_factura); if err != nil {return invoices, err}
		invoices = append(invoices, invoice)
	}

	return invoices, nil
}

//Returns the income of a payment method
func (dao FacturaImpl) Get_Incomes(id int, formaDePago int) (float64, error) {

	var total float64

	query := "SELECT COALESCE(SUM(precio), 0) FROM " +
			 "(SELECT formadepago, precio FROM cliente c INNER JOIN (SELECT id_factura, precio FROM factura WHERE id_caja = $1 AND comentarioBaja SIMILAR TO '') f ON " +
			 "c.id_factura = f.id_factura) f " +
			 "WHERE f.formadepago = $2"

	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return total, err}
	defer stmt.Close()

	row := stmt.QueryRow(id, formaDePago)
	err = row.Scan(&total); if err != nil {return total, err}

	return total, nil
}

//Returns the type of an invoice
func GetTipo(factura models.Factura) string {

	if len(factura.Renglones) == 0 {
		if !factura.Comentario.Valid {
			return "Retiro"
		}
		return "Gastos"
	}
	return "Cliente"
}

//1 efectivo
//2 debito
//3 credito

//Returns the payment method of an invoice
func GetFormaDePago(factura models.Factura) string {

	if factura.FormaDePago.Valid {
		if factura.FormaDePago.Int64 == 1 {
			return "Efectivo"
		} else {
			if factura.FormaDePago.Int64 == 2 {
				return "Débito"
			}
		}
	}
	return "Crédito"
}

//Returns the total number of extractions of a cash box
func GetTotalRetiros(id int) (float64, error) {

	var total float64

	query := "SELECT COALESCE(SUM(precio), 0) FROM factura WHERE id_caja = $1 AND comentarioBaja SIMILAR TO '' " +
		     "AND id_factura NOT IN (SELECT id_factura FROM otros) AND " +
		     "id_factura NOT IN (SELECT id_factura FROM cliente)"

	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return total, err}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(&total); if err != nil {return total, err}

	return total, nil
}

//Returns the total number of expenses of a cash box
func GetTotalGastos(id int) (float64, error) {

	var total float64

	query := "SELECT COALESCE(SUM(precio), 0) FROM otros o INNER JOIN " +
		     "(SELECT id_factura, precio FROM factura WHERE id_caja = $1  AND comentarioBaja SIMILAR TO '') f " +
		     "ON o.id_factura = f.id_factura"

	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return total, err}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(&total); if err != nil {return total, err}

	return total, nil
}