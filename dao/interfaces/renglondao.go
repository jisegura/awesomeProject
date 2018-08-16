package interfaces

import "awesomeProject/models"

type RenglonDAO interface {
	Create(renglon *models.Renglon) error
	//	Update(renglon *models. Renglon) error
	//	Delete(id int) error
	GetAll()([]models. Renglon, error)
}
