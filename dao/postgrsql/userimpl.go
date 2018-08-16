package postgrsql

import "awesomeProject/models"

type UserImpl struct {}

func (dao UserImpl) Create(u *models.User) error {
	query := "INSERT INTO users (firstname, lastname, email) VALUES ($1, $2, $3) RETURNING id"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	row := stmt.QueryRow(u.FirstName, u.LastName, u.Email)
	row.Scan(&u.Id)
	return nil
}

func (dao UserImpl) GetAll()([]models.User, error) {
	users := make([]models.User, 0)
	query := "SELECT * FROM users"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return users, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return users, err
	}

	for rows.Next() {
		var row models.User
		err := rows.Scan(&row.Id, &row.FirstName, &row.LastName)
		if err != nil {
			return users, err
		}
		users = append(users, row)
	}
	return users, nil
}