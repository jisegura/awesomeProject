package main

import (
	"awesomeProject/page"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	router := mux.NewRouter()

	f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Print("Error")
	}
	defer f.Close()

	log.SetOutput(io.MultiWriter(os.Stderr, f))
	log.Println("This is a test log entry")

	//EMPLEADO/////
	router.HandleFunc("/empleado/", page.InsertEmpleado).Methods("POST")
	router.HandleFunc("/empleado/login/{id}", page.AddLogin).Methods("POST")
	router.HandleFunc("/empleado/", page.UpdateEmpleado).Methods("PUT")
	router.HandleFunc("/empleado/baja/", page.UpdateBaja).Methods("PUT")
	router.HandleFunc("/empleado/{id}", page.DeleteEmpleado).Methods("DELETE")
	router.HandleFunc("/empleado/", page.GetEmpleados).Methods("GET")
	router.HandleFunc("/empleado/login/", page.Login).Methods("POST")

	//CATEGORIA/////
	router.HandleFunc("/categoria/", page.GetCategorias).Methods("GET")
	router.HandleFunc("/categoria/{id}", page.GetCategoriaById).Methods("GET")
	router.HandleFunc("/categoria/", page.InsertCategoria).Methods("POST")
	router.HandleFunc("/categoria/{id}", page.DeleteCategoria).Methods("DELETE")
	router.HandleFunc("/categoria/", page.UpdateCategoria).Methods("PUT")

	//PRODUCTO/////
	router.HandleFunc("/producto/", page.GetProductos).Methods("GET")
	router.HandleFunc("/producto/{id}", page.GetProductoById).Methods("GET")
	router.HandleFunc("/producto/", page.InsertProducto).Methods("POST")
	router.HandleFunc("/producto/{id}", page.DeleteProducto).Methods("DELETE")
	router.HandleFunc("/producto/", page.UpdateProducto).Methods("PUT")

	//CAJA/////
	router.HandleFunc("/caja/", page.InsertCaja).Methods("POST")
	router.HandleFunc("/caja/", page.CerrarCaja).Methods("PUT")
	router.HandleFunc("/caja/", page.GetCajas).Methods("GET")
	router.HandleFunc("/caja/open/", page.GetCaja).Methods("GET")
	router.HandleFunc("/caja/export/{fechaInicio}/{fechaFin}/", page.GetCajasByFechas).Methods("GET")
	router.HandleFunc("/caja/efectivo/{id}", page.GetIngresosEfectivo).Methods("GET")
	router.HandleFunc("/caja/debito/{id}", page.GetIngresosDebito).Methods("GET")
	router.HandleFunc("/caja/credito/{id}", page.GetIngresosCredito).Methods("GET")

	//FACTURA/////
	router.HandleFunc("/factura/retiros/", page.InsertFactura).Methods("POST")
	router.HandleFunc("/factura/clientes/", page.InsertCliente).Methods("POST")
	router.HandleFunc("/factura/otros/", page.InsertOtro).Methods("POST")

	router.HandleFunc("/factura/retiros/{id}", page.GetFacturasRetiro).Methods("GET")
	router.HandleFunc("/factura/clientes/{id}", page.GetFacturasCliente).Methods("GET")
	router.HandleFunc("/factura/otros/{id}", page.GetFacturasOtro).Methods("GET")
	router.HandleFunc("/factura/{id}", page.GetAllIdFactura).Methods("GET")
	router.HandleFunc("/factura/all/{id}", page.GetAllFacturas).Methods("GET")
	router.HandleFunc("/factura/", page.GetFacturasEliminadas).Methods("GET")
	router.HandleFunc("/factura/ultimas/{id}", page.GetLastFacturas).Methods("GET")

	router.HandleFunc("/factura/", page.UpdateFactura).Methods("PUT")

	router.HandleFunc("/upload/", page.Upload).Methods("POST")
	router.HandleFunc("/upload/", page.Download).Methods("GET")

	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	log.Print("Escuchando en localhost:3000")
	log.Fatal(http.ListenAndServe("localhost:3000", handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(router)))
}
