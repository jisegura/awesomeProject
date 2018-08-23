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
	ComentarioBaja sql.NullString
	Descuento      sql.NullFloat64
	Comentario     sql.NullString
	FormaDePago    sql.NullInt64
	Renglones      []Renglon `json:"Renglones"`
}
