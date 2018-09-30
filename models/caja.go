package models

import "time"

type Caja struct {
	Id_caja    int
	Inicio     float64
	Fin        float64
	HoraInicio time.Time
	HoraFin    time.Time
	CierreReal float64
}
