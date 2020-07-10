package models

import (
	"github.com/lib/pq"
	"time"
)

type Login_Registration struct {

	Id_login 		 int
	Id_persona 		 int
	Hora_Conexion    time.Time
	Hora_Desconexion pq.NullTime
}
