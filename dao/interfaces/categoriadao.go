package interfaces

import "awesomeProject/models"

type CategoriaDAO interface {
	Create(categoria *models.Categoria) error
	Update(categoria *models.Categoria) error
	Delete(id int) error
	Get_All() ([]models.Categoria, error)
	Get_By_Id(id int) (models.Categoria, error)
}
