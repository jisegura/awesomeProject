package models

import "time"

type Factura struct {
	Id_factura       int
	Id_caja          int
	Id_empleado      int
	Fecha            time.Time
	Precio           float64
	ComentarioBaja   string
	Descuento        float64
	Comentario       string
	DescuentoCliente float64
	FormaDePago      int
	Renglones        []Renglon `json:"Renglones"`
}
