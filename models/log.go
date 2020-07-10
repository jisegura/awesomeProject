package models

import "time"

type Log struct {

	Id_Login       	 int
	Hora_Conexion  	 time.Time
	Hora_Desconexion time.Time
}
