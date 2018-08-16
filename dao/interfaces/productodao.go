package interfaces

import "awesomeProject/models"

type ProductoDAO interface {
	Create(producto *models.Producto) error
	Update(producto *models.Producto) error
	Delete(id int) error
	GetById(id int) (models.Producto, error)
	GetAll()([]models.Producto, error)
}
