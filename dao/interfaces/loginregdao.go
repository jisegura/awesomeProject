package interfaces

type UserDAO interface {

	Login(user string, pass string) (error, bool)
	Logout(id int) (error, bool)
}