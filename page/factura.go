package page

import (
	"awesomeProject/dao/factory"
	"awesomeProject/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

//GET ID_FACTURA POR CAJA
func GetAllIdFactura(w http.ResponseWriter, req *http.Request) {

	facturaDao := factory.FactoryFactura()
	param := mux.Vars(req)
	id, _ := strconv.Atoi(param["id"])

	facturas, err := facturaDao.GetFacturasById(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Error", err)
		return
	}

	json.NewEncoder(w).Encode(&facturas)
}

//GET ALL FACTURAS POR CAJA
func GetAllFacturas(w http.ResponseWriter, req *http.Request) {

	facturaDAO := factory.FactoryFactura()
	param := mux.Vars(req)
	id, _ := strconv.Atoi(param["id"])

	facturas, err := facturaDAO.GetAllFacturasById(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Error", err)
		return
	}

	json.NewEncoder(w).Encode(facturas)
}

//SELECT RETIRO
func GetFacturasRetiro(w http.ResponseWriter, req *http.Request) {

	facturaDAO := factory.FactoryFactura()
	param := mux.Vars(req)
	id, _ := strconv.Atoi(param["id"])

	facturas, err := facturaDAO.GetAll(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	json.NewEncoder(w).Encode(&facturas)
}

//SELECT CLIENTE
func GetFacturasCliente(w http.ResponseWriter, req *http.Request) {

	facturaDAO := factory.FactoryFactura()
	param := mux.Vars(req)
	id, _ := strconv.Atoi(param["id"])

	facturas, err := facturaDAO.GetFacturasCliente(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	json.NewEncoder(w).Encode(&facturas)
}

//SELECT OTRO
func GetFacturasOtro(w http.ResponseWriter, req *http.Request) {

	facturaDAO := factory.FactoryFactura()
	param := mux.Vars(req)
	id, _ := strconv.Atoi(param["id"])

	facturas, err := facturaDAO.GetFacturasOtro(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	json.NewEncoder(w).Encode(&facturas)
}

//INSERT RETIRO
func InsertFactura(w http.ResponseWriter, req *http.Request) {

	facturaDAO := factory.FactoryFactura()
	factura := models.Factura{}

	_ = json.NewDecoder(req.Body).Decode(&factura)
	err := facturaDAO.Create(&factura)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	f, err := facturaDAO.GetById(factura.Id_factura)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	json.NewEncoder(w).Encode(&f)
}

//INSERT CLIENTE
func InsertCliente(w http.ResponseWriter, req *http.Request) {

	facturaDAO := factory.FactoryFactura()
	factura := models.Factura{}

	_ = json.NewDecoder(req.Body).Decode(&factura)

	newFactura, err := facturaDAO.CreateCliente(&factura)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	json.NewEncoder(w).Encode(&newFactura)
}

//INSERT OTRO
func InsertOtro(w http.ResponseWriter, req *http.Request) {

	facturaDAO := factory.FactoryFactura()
	factura := models.Factura{}

	_ = json.NewDecoder(req.Body).Decode(&factura)
	newFactura, err := facturaDAO.CreateOtro(&factura)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	json.NewEncoder(w).Encode(&newFactura)
}

func UpdateFactura(w http.ResponseWriter, req *http.Request) {

	facturaDAO := factory.FactoryFactura()
	factura := models.Factura{}

	_ = json.NewDecoder(req.Body).Decode(&factura)
	fmt.Println(factura)

	err := facturaDAO.Update(&factura)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	f, err := facturaDAO.GetById(factura.Id_factura)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	json.NewEncoder(w).Encode(f)
}

func GetFacturasEliminadas(w http.ResponseWriter, req *http.Request) {

	facturaDAO := factory.FactoryFactura()

	facturas, err := facturaDAO.GetFacturasEliminadas()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	json.NewEncoder(w).Encode(facturas)
}

func GetLastFacturas(w http.ResponseWriter, req *http.Request) {

	facturaDAO := factory.FactoryFactura()

	param := mux.Vars(req)
	id, _ := strconv.Atoi(param["id"])

	facturas, err := facturaDAO.GetLastFacturas(id)
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		log.Print("Error", err)
		return
	}

	json.NewEncoder(w).Encode(&facturas)
}
