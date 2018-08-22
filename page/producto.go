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
func GetProductos(w http.ResponseWriter, req *http.Request) {

	productoDAO := factory.FactoryProducto()

	productos, err := productoDAO.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	json.NewEncoder(w).Encode(&productos)
}

//INSERT
func InsertProducto(w http.ResponseWriter, req *http.Request) {

	productoDAO := factory.FactoryProducto()
	producto := models.Producto{}

	_ = json.NewDecoder(req.Body).Decode(&producto)
	err := productoDAO.Create(&producto)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	p, err := productoDAO.GetById(producto.Id_producto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	json.NewEncoder(w).Encode(&p)
}

//SELECT BY ID
func GetProductoById(w http.ResponseWriter, req *http.Request) {

	productoDao := factory.FactoryProducto()
	param := mux.Vars(req)
	id, _ := strconv.Atoi(param["id"])

	producto, err := productoDao.GetById(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	json.NewEncoder(w).Encode(producto)
}

//DELETE
func DeleteProducto(w http.ResponseWriter, req *http.Request) {

	productoDao := factory.FactoryProducto()
	param := mux.Vars(req)

	id, _ := strconv.Atoi(param["id"])

	err := productoDao.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	p := models.Producto{}

	json.NewEncoder(w).Encode(&p)
}

//UPDATE
func UpdateProducto(w http.ResponseWriter, req *http.Request) {

	productoDao := factory.FactoryProducto()
	producto := models.Producto{}

	_ = json.NewDecoder(req.Body).Decode(&producto)

	err := productoDao.Update(&producto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	p, err := productoDao.GetById(producto.Id_producto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	json.NewEncoder(w).Encode(&p)
}
