package postgrsql

import (
	"os/exec"
	"strings"
)

/*
func InitializeAll() error {
	query := "SELECT table_name " +
			 "FROM information_schema.tables " +
			 "WHERE table_schema not in ('information_schema', 'pg_catalog') " +
			 "AND table_type = 'BASE TABLE'"
	db := getConnection()
	defer db.Close()
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}*/

func InitilizeTable(table string) error {

	query := "TRUNCATE TABLE Empleado"
	db := getConnection()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	return nil
}

func Backup(pass string) error {

	cmd := exec.Command("pg_dump", "-h", "localhost", "-p", "5432", "-U", "postgres", "-W", "-F", "c", "-f", "/home/michelle/Backup.sql", "Iglu")
	cmd.Stdin = strings.NewReader(pass)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func Restore(pass string) error {

	InitilizeTable("empleado")

	cmd := exec.Command("pg_restore", "-h", "localhost", "-p", "5432", "-a", "-U", "postgres", "-W", "-d", "Iglu", "-F", "c", "--no-data-for-failed-tables", "/home/michelle/Backup.sql")
	cmd.Stdin = strings.NewReader(pass)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}