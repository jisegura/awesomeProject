package interfaces

import "awesomeProject/models"

type UserDAO interface {
	Login(user models.User) (bool, error)
	AddUser(user models.User) error
}
