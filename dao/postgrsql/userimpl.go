package postgrsql

import "awesomeProject/models"

type UserImpl struct{}

func (dao UserImpl) Login(user models.User) (bool, error) {

	query := "SELECT contraseña FROM usuario WHERE nombre SIMILAR TO '$1'"
	db := getConnection()
	defer db.Close()

	var correcto bool

	stmt, err := db.Prepare(query)
	if err != nil {
		return correcto, err
	}
	defer stmt.Close()

	var contraseñaCorrecta string
	row := stmt.QueryRow(user.Nombre)
	err = row.Scan(&contraseñaCorrecta)
	if err != nil {
		return err
	}

	if contraseñaCorrecta == user.Contraseña {
		return true, nil
	}

	return false, nil
}

func (dao UserImpl) AddUser(user models.User) error {

	query := "INSERT INTO usuario (nombre, contraseña) VALUES $1, $2 RETURNING id_usuario"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	row := stmt.QueryRow(user.Nombre, user.Contraseña)
	err = row.Scan(&user.Id_user)
	if err != nil {
		return err
	}

	return nil
}
