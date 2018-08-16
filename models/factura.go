package models

import "time"

type Factura struct {
	Id_factura  int
	Id_caja     int
	Precio      float64
	Fecha       time.Time
	Descuento   float64
	Activa      bool
}
