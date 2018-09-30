package interfaces

import (
	"awesomeProject/models"
	"time"
)

type CajaDAO interface {
	Create(caja *models.Caja) (models.Caja, error)
	GetAll() ([]models.Caja, error)
	GetCaja() (models.Caja, error)
	CierreCaja(caja *models.Caja) (models.Caja, error)
	GetCajasByFechas(fechaIncio time.Time, fechaFin time.Time) ([]models.Caja, error)
}
