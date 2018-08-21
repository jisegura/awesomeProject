package interfaces

import "awesomeProject/models"

type CajaDAO interface {
	Create(caja *models.Caja) (models.Caja, error)
	GetAll() ([]models.Caja, error)
	GetCaja() (models.Caja, error)
	CierreCaja(caja *models.Caja) error
}
