package interfaces

import "awesomeProject/models"

type UserDAO interface {
	AddUser(user models.User, id int) error
	Login(user models.User) (error, bool)
	Logout(id int) error
}
