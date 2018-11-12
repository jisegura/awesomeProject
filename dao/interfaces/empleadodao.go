package interfaces

import "awesomeProject/models"

type EmpleadoDAO interface {
	Create(empleado *models.Empleado) error
	UpdateNombre(empleado *models.Empleado) error
	UpdateBaja(id int) error
	Delete(id int) error
	GetById(id int) (models.Empleado, error)
	GetAll() ([]models.Empleado, error)
	AddLogin(login models.Login, id int) error
	Login(login models.Login) (bool, error)
}
