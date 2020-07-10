package interfaces

import "awesomeProject/models"

type RolDAO interface {

	Create(rol models.Rol) (models.Rol, error)
	Get_By_Id (id int) (models.Rol, error)
	Get_All()([]models.Rol, error)
	Get_Attr(id int, attrName string) (string, error)
	Update(id int, attrName string, attr string) (models.Rol, error)
}
