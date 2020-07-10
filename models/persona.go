package models

import (
	"database/sql"
	"github.com/lib/pq"
	"time"
)

type Persona struct {

	Id_persona  int
	Id_rol      int

	Nombre	    string
	Apellido    sql.NullString
	Telefono    sql.NullString
	Mail        sql.NullString
	Direccion   sql.NullString

	Fecha_alta  time.Time
	Fecha_baja  pq.NullTime

	Usuario     sql.NullString
	Contrasenia sql.NullString
}
