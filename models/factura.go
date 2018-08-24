package models

import (
	"database/sql"
	"time"
)

type Factura struct {
	Id_factura     int
	Id_caja        int
	Id_empleado    int
	Fecha          time.Time
	Precio         float64
	ComentarioBaja string
	Descuento      sql.NullInt64
	Comentario     sql.NullString
	FormaDePago    sql.NullInt64
	Renglones      []Renglon `json:"Renglones"`
}
