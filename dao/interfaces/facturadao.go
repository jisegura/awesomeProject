package interfaces

import "awesomeProject/models"

type FacturaDAO interface {
	Create(factura *models.Factura) error
	//	Update(factura *models.Factura) error
	//	Delete(id int) error
	GetAll()([]models.Factura, error)
}