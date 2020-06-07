package postgrsql

import (
	"awesomeProject/utilities"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os/exec"
	"strings"
	"time"
)

func getConnection() *sql.DB {

	config, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatalln(err)
	}

	//postgres ://user:password@server:port/database?sslmode=false
	dns := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.User, config.Password, config.Host, config.Port, config.Database)
	db, err := sql.Open("postgres", dns)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Backup() error {

	cmd := exec.Command("pg_dump", "-h", "localhost", "-p", "5432", "-U", "postgres", "-W", "-F", "c", "-f", "/home/michelle/Backup.sql", "Iglu")
	cmd.Stdin = strings.NewReader("cocacola")
	err := cmd.Run()
	if err == nil {
		fmt.Println(cmd.CombinedOutput())
	}
	return err
}

func Restore() error {
	today := time.Now()
	fmt.Println(today)
	cmd := exec.Command("pg_restore", "-h", "localhost", "-p", "5432", "-U", "postgres", "-W", "-d", "Iglu", "/home/michelle/Backup.sql")
	cmd.Stdin = strings.NewReader("cocacola")
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
