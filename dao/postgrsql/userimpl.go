package postgrsql

type LoginReg struct{}
/*
//Allows login a user
func (dao LoginReg) Login (user string, pass string) (error, bool) {

	var id int

	ok, err := dao.verify_Data(user, pass); if err != nil {return err, false}

	if ok {

		query := "INSERT INTO Registro_Login (id_persona, hora_login, hora_logout) " +
				 "VALUES ($1, $2, $3) " +
			     "RETURNING Id_login"

		db := getConnection()
		defer db.Close()

		stmt, err := db.Prepare(query); if err != nil {return err, false}
		defer stmt.Close()

		id, err = dao.Get_User_Id(user); if err != nil {return err, false}

		row := stmt.QueryRow(id, time.Now(), pq.NullTime{})
		err = row.Scan(&id); if err != nil {return err, false}
	}

	return nil, true
}

// Allows deloge a user
func (dao LoginReg) Logout () error {

	query := "UPDATE Registro_Login " +
			 "SET hora_logout = $1 "

	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return err}
	defer stmt.Close()

	row, err := stmt.Exec(time.Now()); if err != nil {return err}

	i, _ := row.RowsAffected()
	if i != 1 {
		return errors.New("Error, se esperaba una fila afectada")
	}

	err = Delete_Reg(); if err != nil {return err}

	return nil
}

//Returns if there is a logged in user
func (dao LoginReg) User_Is_Logged_In () (int, error) {

	var id_persona int

	query := "SELECT Id_persona FROM Registro_Login"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return id_persona, err}
	defer stmt.Close()

	row := stmt.QueryRow()
	err = row.Scan(&id_persona); if err != nil {
		if err == sql.ErrNoRows {
			return  id_persona, nil
		}
	}

	return id_persona, err
}

//Allows removing the user from the login table
func Delete_Reg () error {

	query := "DELETE FROM Registro_Login"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return err}
	defer stmt.Close()

	row, err := stmt.Exec(); if err != nil {return err}

	i, _ := row.RowsAffected(); if i != 1 {return errors.New("Error, se esperaba una fila afectada")}

	return nil
}

func HashPassword(password string) (string, error) {

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), 10); if err != nil {return "", err}

	// Encode the entire thing as base64 and return
	hashBase64 := base64.StdEncoding.EncodeToString(hashedBytes)

	return hashBase64, nil
}

func ComparePassword(hashBase64, testPassword string) bool {

	// Decode the real hashed and salted password so we can
	//	// split out the salt
	hashBytes, err := base64.StdEncoding.DecodeString(hashBase64)
	if err != nil {
		fmt.Println("Error, we were given invalid base64 string", err)
		return false
	}

	err = bcrypt.CompareHashAndPassword(hashBytes, []byte(testPassword))
	return err == nil
}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {

	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomString(n int) (string, error) {

	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	bytes, err := GenerateRandomBytes(n)
	if err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}

	hashBase64, err := HashPassword(string(bytes))
	log.Print("pass ", string(bytes), "  hash: ", hashBase64)
	otro, err := HashPassword("cocacola")
	log.Print("otro: ", otro)

	return hashBase64, nil
}*/