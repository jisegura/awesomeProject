package models

import (
	"time"
)

type Factura struct {
	Id_factura     int
	Id_caja        int
	Id_empleado    int
	Fecha          time.Time
	Precio         float64
	ComentarioBaja string
	Descuento      int
	Comentario     string
	FormaDePago    int
	Renglones      []Renglon `json:"Renglones"`
}
