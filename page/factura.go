package page

import (
	"net/http"
	"awesomeProject/dao/factory"
	"log"
	"encoding/json"
	"awesomeProject/models"
	"time"
)

//SELECT
func GetFacturas (w http.ResponseWriter, req *http.Request) {

	facturaDAO := factory.FactoryFactura()

	facturas, err := facturaDAO.GetAll()
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(&facturas)
}

//INSERT
func InsertFacturas(w http.ResponseWriter, req *http.Request)  {

	facturaDAO := factory.FactoryFactura()
	factura := models.Factura{}

	_ = json.NewDecoder(req.Body).Decode(&factura)
	factura.Fecha = time.Now()
	err := facturaDAO.Create(&factura)

	if err != nil {
		log.Fatal(err)
	}

	facturas, err := facturaDAO.GetAll()
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(&facturas)
}
