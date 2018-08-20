package interfaces

import "awesomeProject/models"

type CategoriaDAO interface {
	Create(categoria *models.Categoria) error
	Update(categoria *models.Categoria) error
	Delete(id int) error
	GetAll() ([]models.Categoria, error)
	GetById(id int) (models.Categoria, error)
}
