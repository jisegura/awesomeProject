package main

import (
	"awesomeProject/page"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()

	//EMPLEADO/////
	router.HandleFunc("/empleado/", page.InsertEmpleado).Methods("POST")
	router.HandleFunc("/empleado/", page.UpdateEmpleado).Methods("PUT")
	router.HandleFunc("/empleado/{id}", page.DeleteEmpleado).Methods("DELETE")
	router.HandleFunc("/empleado/", page.GetEmpleados).Methods("GET")

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

	//FACTURA/////
	router.HandleFunc("/factura/retiros/", page.InsertFactura).Methods("POST")
	router.HandleFunc("/factura/clientes/", page.InsertCliente).Methods("POST")
	router.HandleFunc("/factura/otros/", page.InsertOtro).Methods("POST")

	router.HandleFunc("/factura/retiros/{id}", page.GetFacturasRetiro).Methods("GET")
	router.HandleFunc("/factura/clientes/{id}", page.GetFacturasCliente).Methods("GET")
	router.HandleFunc("/factura/otros/{id}", page.GetFacturasOtro).Methods("GET")
	router.HandleFunc("/factura/", page.GetFacturasEliminadas).Methods("GET")

	router.HandleFunc("/factura/", page.UpdateFactura).Methods("PUT")

	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	log.Fatal(http.ListenAndServe("localhost:3000", handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(router)))
}
