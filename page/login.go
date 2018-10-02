package page

import (
	"awesomeProject/dao/factory"
	"awesomeProject/models"
	"encoding/json"
	"log"
	"net/http"
)

func HashPassword(w http.ResponseWriter, req *http.Request) {

	loginDAO := factory.FactoryLogin()
	login := models.Login{}

	_ = json.NewDecoder(req.Body).Decode(&login)

	err := loginDAO.AddLogin(login)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func Login(w http.ResponseWriter, req *http.Request) {

	loginDAO := factory.FactoryLogin()
	login := models.Login{}

	_ = json.NewDecoder(req.Body).Decode(&login)

	correct, err := loginDAO.Login(login)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	json.NewEncoder(w).Encode(&correct)
}
