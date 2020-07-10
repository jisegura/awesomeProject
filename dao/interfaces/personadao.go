package interfaces

import "awesomeProject/models"

type EmpleadoDAO interface {
	Create(empleado *models.Empleado) error
	Get_All() ([]models.Empleado, error)
	Get_By_Id(id int) (models.Empleado, error)
	Delete(id int) error
	Update_Name(empleado *models.Empleado) error
	Unsubscribe(id int) error



}
