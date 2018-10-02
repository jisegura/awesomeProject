package models

import (
	"database/sql"
	"github.com/lib/pq"
)

type Empleado struct {
	Id_empleado int
	Id_login    sql.NullInt64
	FechaBaja   pq.NullTime
	FirstName   string
	LastName    string
}
