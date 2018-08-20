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
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(&cajas)
}

//INSERT
func InsertCaja(w http.ResponseWriter, req *http.Request) {

	cajasDAO := factory.FactoryCaja()
	caja := models.Caja{}

	_ = json.NewDecoder(req.Body).Decode(&caja)
	err := cajasDAO.Create(&caja)

	if err != nil {
		log.Fatal(err)
	}

	cajas, err := cajasDAO.GetAll()
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(&cajas)
}

/*
//UPDATE
func cerrarCaja (w http.ResponseWriter, req *http.Request) {

}*/
