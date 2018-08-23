package models

import "database/sql"

type Empleado struct {
	Id_empleado int
	FirstName   string
	LastName    sql.NullString
}
