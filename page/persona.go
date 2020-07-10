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

//Allows to add a new person
func Create_Person (w http.ResponseWriter, req *http.Request) {

	personDAO := factory.FactoryPersona()
	person := models.Persona{}

	_ = json.NewDecoder(req.Body).Decode(&person)
	err := personDAO.Create(&person)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	personAux, err := personDAO.Get_By_Id(person.Id_persona)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	json.NewEncoder(w).Encode(&personAux)
}


//Returns all the people
func Get_People(w http.ResponseWriter, req *http.Request) {

	personaDAO := factory.FactoryPersona()

	people, err := personaDAO.Get_All()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	json.NewEncoder(w).Encode(&people)
}

//Returns a person
func Get_Person_By_Id(w http.ResponseWriter, req *http.Request) {

	personDAO := factory.FactoryPersona()
	param := mux.Vars(req)

	id, _ := strconv.Atoi(param["id"])
	person, err := personDAO.Get_By_Id(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	json.NewEncoder(w).Encode(&person)
}

//Allows to delete a person
func Delete_Person(w http.ResponseWriter, req *http.Request) {

	personDAO := factory.FactoryPersona()
	param := mux.Vars(req)

	id, _ := strconv.Atoi(param["id"])

	err := personDAO.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}
	e := models.Persona{}

	json.NewEncoder(w).Encode(&e)
}

//Allows updating an attribute of a person
func Update_Attr(w http.ResponseWriter, req *http.Request) {

	personDAO := factory.FactoryPersona()
	param := mux.Vars(req)

	id, _ := strconv.Atoi(param["id"])
	attrName, _ := param["attrName"]
	attr, _ := param["attr"]

	err := personDAO.Update_Attribute(id, attrName, attr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	e, err := personDAO.Get_By_Id(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	json.NewEncoder(w).Encode(&e)
}

//Returns an attribute of a person
func Get_Attr (w http.ResponseWriter, req *http.Request) {

	personDAO := factory.FactoryPersona()
	param := mux.Vars(req)

	id, _ := strconv.Atoi(param["id"])
	attrName, _ := param["attr"]


	attribute, err := personDAO.Get_Attr(id, attrName)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	json.NewEncoder(w).Encode(&attribute)
}

//Returns the rol of a person
func Get_Role (w http.ResponseWriter, req *http.Request) {

	personDAO := factory.FactoryPersona()
	param := mux.Vars(req)

	id, _ := strconv.Atoi(param["id"])
	rol, err := personDAO.Get_Role(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	json.NewEncoder(w).Encode(&rol)
}

//Allows a person to change the password
func Change_Password (w http.ResponseWriter, req *http.Request) {

	personDAO := factory.FactoryPersona()
	param := mux.Vars(req)

	user := param["user"]
	pass := param["pass"]
	newPass := param["newPass"]

	log.Print("user:", user, " pass: ", pass, " newPass: ", newPass)
	err := personDAO.Change_Password(user, pass, newPass)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

//Allows a person to set the password
func Set_Password (w http.ResponseWriter, req *http.Request) {

	personDAO := factory.FactoryPersona()
	param := mux.Vars(req)

	user := param["user"]
	pass := param["pass"]

	err := personDAO.Set_Password(user, pass)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

//Allows a person to reset the password
func Reset_Password (w http.ResponseWriter, req *http.Request) {

	personDAO := factory.FactoryPersona()
	param := mux.Vars(req)

	user := param["user"]

	err := personDAO.Reset_Password(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

//Allows a person to unsubscribe
func Unsubscribe(w http.ResponseWriter, req *http.Request) {

	personDAO := factory.FactoryPersona()

	param := mux.Vars(req)

	id, _ := strconv.Atoi(param["id"])

	err := personDAO.Unsubscribe(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	e, err := personDAO.Get_By_Id(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	json.NewEncoder(w).Encode(&e)
}