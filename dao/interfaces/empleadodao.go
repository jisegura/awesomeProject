package interfaces

import "awesomeProject/models"

type EmpleadoDAO interface {
	Create(empleado *models.Empleado) error
	Update(empleado *models.Empleado) error
	Delete(id int) error
	GetById(id int) (models.Empleado, error)
	GetAll() ([]models.Empleado, error)
}
