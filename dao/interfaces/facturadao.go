package interfaces

import (
	"awesomeProject/models"
)

type FacturaDAO interface {
	Create(factura *models.Factura) error
	CreateCliente(factura *models.Factura) error
	CreateOtro(factura *models.Factura) error

	Update(factura *models.Factura) error

	GetFacturasEliminadas() ([]models.Factura, error)

	GetAll(id int) ([]models.Factura, error)
	GetById(id int) (models.Factura, error)

	GetFacturasCliente(id int) ([]models.Factura, error)
	GetByIdCliente(id int) (models.Factura, error)

	GetFacturasOtro(id int) ([]models.Factura, error)
	GetByIdOtros(id int) (models.Factura, error)
}
