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

//Allows to add a new rol
func Create_Role (w http.ResponseWriter, req *http.Request) {

	rolDAO := factory.FactoryRol()
	role := models.Rol{}

	_ = json.NewDecoder(req.Body).Decode(&role)

	role , err := rolDAO.Create(role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	json.NewEncoder(w).Encode(&role)
}
//Returns a role given an id
func Get_Role_By_Id (w http.ResponseWriter, req *http.Request) {

	roleDAO := factory.FactoryRol()

	param := mux.Vars(req)
	id, _ := strconv.Atoi(param["id"])

	role, err := roleDAO.Get_By_Id(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	json.NewEncoder(w).Encode(&role)
}

//Returns all roles
func Get_All_Roles (w http.ResponseWriter, req *http.Request) {

	roleDAO := factory.FactoryRol()

	roles, err := roleDAO.Get_All()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	json.NewEncoder(w).Encode(&roles)
}

//Returns an attribute of the role
func Get_Role_Attr (w http.ResponseWriter, req *http.Request) {

	roleDAO := factory.FactoryRol()

	param := mux.Vars(req)
	id, _ := strconv.Atoi(param["id"])
	attrName, _ := param["attrName"]

	role, err := roleDAO.Get_Attr(id, attrName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	json.NewEncoder(w).Encode(role)

}

//Allows updating an attribute
func Update_Rol_Attr (w http.ResponseWriter, req *http.Request) {

	roleDAO := factory.FactoryRol()

	param := mux.Vars(req)
	id, _ := strconv.Atoi(param["id"])
	attrName, _ := param["attrName"]
	attr, _ := param["attr"]

	role, err := roleDAO.Update(id, attrName, attr)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	json.NewEncoder(w).Encode(&role)
}