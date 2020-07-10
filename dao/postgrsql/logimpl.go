package postgrsql

type LogImpl struct {}

/*

type Login struct {

	Id_Login       	 int
	Hora_Conexion  	 time.Time
	Hora_Desconexion time.Time
}

//Returns the log by day
func Log () ([]models.Login, error) {

	logs := make([]models.Login, 0)
	var row models.Login

	query := "SELECT * FROM Login"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query); if err != nil {return logs, err}
	defer stmt.Close()

	rows, err := stmt.Query(); if err != nil {return logs, err}

	for rows.Next() {

		err := rows.Scan(&row.Id_Login, &row.Hora_Conexion, &row.Hora_Desconexion); if err != nil {return logs, err}
		logs = append(logs, row)
	}

	return nil, logs
}*/

