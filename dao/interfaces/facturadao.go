package interfaces

import (
	"awesomeProject/models"
)

type FacturaDAO interface {
	Create(factura *models.Factura) error
	CreateCliente(factura *models.Factura) (models.Factura, error)
	CreateOtro(factura *models.Factura) (models.Factura, error)

	Update(factura *models.Factura) error

	GetFacturasEliminadas() ([]models.Factura, error)
	GetFacturasById(id int) ([]int, error)
	GetLastFacturas(id int) ([]models.Factura, error)
	GetAllFacturasById(id int) ([]models.Factura, error)

	GetAll(id int) ([]models.Factura, error)
	GetById(id int) (models.Factura, error)

	GetFacturasCliente(id int) ([]models.Factura, error)
	GetByIdCliente(id int) (models.Factura, error)

	GetFacturasOtro(id int) ([]models.Factura, error)
	GetByIdOtros(id int) (models.Factura, error)

	GetIngresos(id int, formaDePago int) (float64, error)
}
