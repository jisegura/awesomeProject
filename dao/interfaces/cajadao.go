package interfaces

import "awesomeProject/models"

type CajaDAO interface {
	Create(caja *models.Caja) error
	//Update(caja *models.Caja) error
	GetAll() ([]models.Caja, error)
	//GetById(int)(*models.Caja, error)
	CierreCaja(caja *models.Caja) error
}
