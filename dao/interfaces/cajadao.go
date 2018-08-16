package interfaces

import "awesomeProject/models"

type CajaDAO interface {
	Create(caja *models.Caja) error
	//	Update(caja *models.Caja) error
	//	Delete(id int) error
	GetAll()([]models.Caja, error)
}
