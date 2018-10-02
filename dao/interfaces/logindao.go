package interfaces

import "awesomeProject/models"

type LoginDAO interface {
	AddLogin(login models.Login) error
	Login(login models.Login) (bool, error)
}
