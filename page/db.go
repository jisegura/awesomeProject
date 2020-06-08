package page

import (
	"awesomeProject/dao/postgrsql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Backup(w http.ResponseWriter, req *http.Request) {

	param := mux.Vars(req)
	pass, _ := (param["pass"])

	err := postgrsql.Backup(pass)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

func Restore(w http.ResponseWriter, req *http.Request) {

	param := mux.Vars(req)
	pass, _ := (param["pass"])

	err := postgrsql.Restore(pass)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

func Initilize(w http.ResponseWriter, req *http.Request) {

	err := postgrsql.InitilizeTable("empleado")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
