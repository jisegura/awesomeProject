package postgrsql

func InitializeAll() error {

	query := "SELECT table_name" +
		"FROM information_schema.tables" +
		"WHERE table_schema not in ('information_schema', 'pg_catalog')" +
		"AND table_type = 'BASE TABLE'"

	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	//row := stmt.QueryRow()

	return nil
}

func InitilizeTable(table string) error {

	query := "TRUNCATE table CASCADE"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return nil
}
