package page

import (
	"net/http"
	"awesomeProject/dao/factory"
	"encoding/json"
	"log"
	"awesomeProject/models"
	"github.com/gorilla/mux"
	"strconv"
	"fmt"
)

//SELECT
func GetProductos (w http.ResponseWriter, req *http.Request) {

	productoDAO := factory.FactoryProducto()

	productos, err := productoDAO.GetAll()
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(&productos)
}

//INSERT
func InsertProducto(w http.ResponseWriter, req *http.Request)  {

	productoDAO := factory.FactoryProducto()
	producto := models.Producto{}

	_ = json.NewDecoder(req.Body).Decode(&producto)
	err := productoDAO.Create(&producto)

	if err != nil {
		log.Fatal(err)
	}

	 productos, err := productoDAO.GetAll()
	 if err != nil {
	 	log.Fatal(err)
	 }

	 json.NewEncoder(w).Encode(&productos)
}

//SELECT BY ID
func GetProductoById (w http.ResponseWriter, req *http.Request) {

	productoDao := factory.FactoryProducto()
	param := mux.Vars(req)

	i, _ := strconv.Atoi(param["id"])

	producto, err := productoDao.GetById(i)
	if err != nil {
		fmt.Fprint(w, "No existe el producto")
		return
	}

	json.NewEncoder(w).Encode(producto)
}

//DELETE
func DeleteProducto (w http.ResponseWriter, req *http.Request) {

	productoDao := factory.FactoryProducto()
	param := mux.Vars(req)

	i, _ := strconv.Atoi(param["id"])

	err := productoDao.Delete(i)

	productos, err := productoDao.GetAll()
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(&productos)
}

//UPDATE
func UpdateProducto(w http.ResponseWriter, req *http.Request) {

	productoDao := factory.FactoryProducto()
	producto := models.Producto{}

	_ = json.NewDecoder(req.Body).Decode(&producto)

	err := productoDao.Update(&producto)
	if err != nil {
		return
	}

	productos, err := productoDao.GetAll()
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(&productos)
}