package page

import (
	"awesomeProject/dao/factory"
	"awesomeProject/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func AddUser (w http.ResponseWriter, req *http.Request) {

	user := models.User{}
	userDAO := factory.FactoryUser()

	_ = json.NewDecoder(req.Body).Decode(&user)

	param := mux.Vars(req)
	id, _ := strconv.Atoi(param["id"])

	err := userDAO.AddUser(user, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusOK)
		log.Print("Error: ", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func Login (w http.ResponseWriter, req *http.Request) {

	user := models.User{}
	userDAO := factory.FactoryUser()

	_ = json.NewDecoder(req.Body).Decode(&user)

	err, ok := userDAO.Login(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusOK)
		log.Print("Error: ", err)
		return
	}

	if ok {
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

func Logout (w http.ResponseWriter, req *http.Request) {

	userDAO := factory.FactoryUser()

	param := mux.Vars(req)
	id, _ := strconv.Atoi(param["id"])

	err := userDAO.Logout(id); if err != nil {
		http.Error(w, err.Error(), http.StatusOK)
		log.Print("Error: ", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}