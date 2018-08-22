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

//SELECT
func GetEmpleados(w http.ResponseWriter, req *http.Request) {

	empleadoDAO := factory.FactoryEmpleado()

	empleados, err := empleadoDAO.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	json.NewEncoder(w).Encode(&empleados)
}

//INSERT
func InsertEmpleado(w http.ResponseWriter, req *http.Request) {

	empleadoDAO := factory.FactoryEmpleado()
	empleado := models.Empleado{}

	_ = json.NewDecoder(req.Body).Decode(&empleado)
	err := empleadoDAO.Create(&empleado)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	e, err := empleadoDAO.GetById(empleado.Id_empleado)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	json.NewEncoder(w).Encode(&e)
}

/*
//SELECT BY ID
func GetEmpleadoById (w http.ResponseWriter, req *http.Request) {

	empleadoDao := factory.FactoryEmpleado()
	param := mux.Vars(req)

	i, _ := strconv.Atoi(param["id"])

	empleado, err := empleadoDao.GetById(i)
	if err != nil {
		fmt.Fprint(w, "No existe el empleado")
		return
	}

	json.NewEncoder(w).Encode(empleado)
}*/

//DELETE
func DeleteEmpleado(w http.ResponseWriter, req *http.Request) {

	empleadoDao := factory.FactoryEmpleado()
	param := mux.Vars(req)

	id, _ := strconv.Atoi(param["id"])

	err := empleadoDao.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}
	e := models.Empleado{}

	json.NewEncoder(w).Encode(&e)
}

//UPDATE
func UpdateEmpleado(w http.ResponseWriter, req *http.Request) {

	empleadoDAO := factory.FactoryEmpleado()
	empleado := models.Empleado{}

	_ = json.NewDecoder(req.Body).Decode(&empleado)

	err := empleadoDAO.Update(&empleado)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	e, err := empleadoDAO.GetById(empleado.Id_empleado)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	json.NewEncoder(w).Encode(&e)
}
