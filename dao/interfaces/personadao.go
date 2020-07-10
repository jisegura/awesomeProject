package interfaces

import (
	"awesomeProject/models"
	"database/sql"
)

type PersonaDAO interface {

	Create(empleado *models.Persona) error
	Get_All() ([]models.Persona, error)
	Get_By_Id(id int) (models.Persona, error)
	Delete(id int) error

	Update_Attribute(id int, attrName string, attr string) error
	Get_Attr(id int, attrName string) (sql.NullString, error)
	Get_Role(id int) (int, error)
	Get_User_Id (user string) (int, error)

	Unsubscribe(id int) error

	Set_Password(user string, pass string) error
	Change_Password(user string, pass string, newPass string) error
	Reset_Password(user string) error
	Is_First_Login(user string) (error, bool)


	Login(user string, pass string) (error, bool)
	Logout() error
	User_Is_Logged_In () (int, error)
	Verify_Data(user string, pass string) (bool, error)
}
