package models

import "database/sql"

type Renglon struct {
	Id_renglon  int
	Id_producto sql.NullInt64
	Id_factura  int
	Cantidad    int
	Precio      float64
	Descuento   float64
}
