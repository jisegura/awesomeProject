package page

import (
	"awesomeProject/dao/factory"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//Allows login a user
func Login (w http.ResponseWriter, req *http.Request) {

	personDAO := factory.FactoryPersona()
	param := mux.Vars(req)
	user, _ := param["user"]
	pass, _ := param["pass"]

	err, firstLogin := factory.FactoryPersona().Is_First_Login(user); if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Error: ", err)
	}

	if firstLogin {
		w.WriteHeader(http.StatusSeeOther)
		return
	}

	err, ok := personDAO.Login(user, pass)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}
	if ok {
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusUnauthorized)
}

func Logout (w http.ResponseWriter, req *http.Request) {

	personDAO := factory.FactoryPersona()

	err := personDAO.Logout(); if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func User_Is_Logged_In (w http.ResponseWriter, req *http.Request) {

	personDAO := factory.FactoryPersona()

	user, err := personDAO.User_Is_Logged_In()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Error: ", err)

	}

	json.NewEncoder(w).Encode(&user)
}
