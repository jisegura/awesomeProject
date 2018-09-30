package page

import (
	"awesomeProject/dao/factory"
	"awesomeProject/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
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

func GetCajasByFechas(w http.ResponseWriter, req *http.Request) {

	cajaDAO := factory.FactoryCaja()
	facturaDAO := factory.FactoryFactura()
	excelDAO := factory.FactoryExcel()

	param := mux.Vars(req)
	i, err := strconv.ParseInt(param["fechaInicio"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}
	j, err := strconv.ParseInt(param["fechaFin"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	fechaInicio := time.Unix(i, 0)
	fechaFin := time.Unix(j, 0)

	cajas, err := cajaDAO.GetCajasByFechas(fechaInicio, fechaFin)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}
	var movimientos []models.Movimientos

	for i := range cajas {
		var movimiento models.Movimientos
		facturas, err := facturaDAO.GetAllFacturasById(cajas[i].Id_caja)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Print("Error: ", err)
			return
		}
		movimiento.Caja = cajas[i]
		movimiento.Facturas = facturas
		movimientos = append(movimientos, movimiento)
	}

	err = excelDAO.Export(movimientos)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetIngresosEfectivo(w http.ResponseWriter, req *http.Request) {

	facturaDAO := factory.FactoryFactura()

	param := mux.Vars(req)
	id, _ := strconv.Atoi(param["id"])

	total, err := facturaDAO.GetIngresos(id, 1)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	json.NewEncoder(w).Encode(&total)
}

func GetIngresosDebito(w http.ResponseWriter, req *http.Request) {

	facturaDAO := factory.FactoryFactura()

	param := mux.Vars(req)
	id, _ := strconv.Atoi(param["id"])

	total, err := facturaDAO.GetIngresos(id, 2)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	json.NewEncoder(w).Encode(&total)
}

func GetIngresosCredito(w http.ResponseWriter, req *http.Request) {

	facturaDAO := factory.FactoryFactura()

	param := mux.Vars(req)
	id, _ := strconv.Atoi(param["id"])

	total, err := facturaDAO.GetIngresos(id, 3)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	json.NewEncoder(w).Encode(&total)
}
