package page

import (
	"awesomeProject/dao/factory"
	"awesomeProject/models"
	"encoding/json"
	"log"
	"net/http"
)

//SELECT
func GetCajas(w http.ResponseWriter, req *http.Request) {

	cajasDAO := factory.FactoryCaja()

	cajas, err := cajasDAO.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	json.NewEncoder(w).Encode(&cajas)
}

func GetCaja(w http.ResponseWriter, req *http.Request) {

	cajaDAO := factory.FactoryCaja()

	caja, err := cajaDAO.GetCaja()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	json.NewEncoder(w).Encode(&caja)
}

//INSERT
func InsertCaja(w http.ResponseWriter, req *http.Request) {

	cajasDAO := factory.FactoryCaja()
	caja := models.Caja{}

	_ = json.NewDecoder(req.Body).Decode(&caja)
	c, err := cajasDAO.Create(&caja)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	json.NewEncoder(w).Encode(&c)
}

//UPDATE
func CerrarCaja(w http.ResponseWriter, req *http.Request) {

	cajaDAO := factory.FactoryCaja()
	caja := models.Caja{}

	_ = json.NewDecoder(req.Body).Decode(&caja)

	c, err := cajaDAO.CierreCaja(&caja)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	json.NewEncoder(w).Encode(&c)
}
