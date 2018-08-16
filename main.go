package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"awesomeProject/page"
)

func main(){
/*
	userDAO := factory.FactoryDAO()
	user := models.User{}

	fmt.Print("Nombre: ")
	fmt.Scan(&user.FirstName)
	fmt.Print("Apellido: ")
	fmt.Scan(&user.LastName)
	fmt.Print("Mail: ")
	fmt.Scan(&user.Email)

	err := userDAO.Create(&user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(user)*/

	router := mux.NewRouter()

	//router.HandleFunc("/producto/", page.GetProductos).Methods("GET")
	//router.HandleFunc("/producto/{id}", page.GetProductoById).Methods("GET")
	//router.HandleFunc("/producto/", page.InsertProducto).Methods("POST")
	//router.HandleFunc("/producto/{id}", page.DeleteProducto).Methods("DELETE")
	//router.HandleFunc("/producto/", page.UpdateProducto).Methods("PUT")
	//router.HandleFunc("/caja/", page.InsertCaja).Methods("POST")
	router.HandleFunc("/factura/", page.InsertFacturas).Methods("POST")
	log.Fatal(http.ListenAndServe("192.168.1.35:3000", router))
}
