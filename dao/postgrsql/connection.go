package postgrsql

import (
	"awesomeProject/utilities"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
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
