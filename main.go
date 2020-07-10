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
	if err != nil {log.Print("Error")}
	defer f.Close()

	log.SetOutput(io.MultiWriter(os.Stderr, f))
	log.Println("This is a test log entry")

	//------------------------------------------------------------------------------------//
	//------------------------DATA BASE INIT ---------------------------------------------//
	//------------------------------------------------------------------------------------//
	router.HandleFunc("/init/", page.Initilize).Methods("POST")
	//------------------------------------------------------------------------------------//
	//------------------------BACKUP------------------------------------------------------//
	//------------------------------------------------------------------------------------//
	router.HandleFunc("/backup/{pass}", page.Backup).Methods("GET")

	//------------------------------------------------------------------------------------//
	//------------------------RESTORE-----------------------------------------------------//
	//------------------------------------------------------------------------------------//
	router.HandleFunc("/restore/{pass}", page.Restore).Methods("GET")

	//------------------------------------------------------------------------------------//
	//------------------------LOG---------------------------------------------------------//
	//------------------------------------------------------------------------------------//

	//------------------------------------------------------------------------------------//
	//------------------------ROLE---------------------------------------------------------//
	//------------------------------------------------------------------------------------//
	router.HandleFunc("/role/", page.Create_Role).Methods("POST")

	router.HandleFunc("/role/{id}", page.Get_Role_By_Id).Methods("GET")
	router.HandleFunc("/role/", page.Get_All_Roles).Methods("GET")
	router.HandleFunc("/role/{id},{attrName}", page.Get_Role_Attr).Methods("GET")

	router.HandleFunc("/role/{id},{attrName},{attr}", page.Update_Rol_Attr).Methods("PUT")

	//------------------------------------------------------------------------------------//
	//------------------------PERSONA-----------------------------------------------------//
	//------------------------------------------------------------------------------------//
	router.HandleFunc("/person/", page.Create_Person).Methods("POST")

	router.HandleFunc("/person/", page.Get_People).Methods("GET")
	router.HandleFunc("/person/{id}", page.Get_Person_By_Id).Methods("GET")
	router.HandleFunc("/person/{attr}/{id}", page.Get_Attr).Methods("GET")
	router.HandleFunc("/person/role/{id}", page.Get_Role).Methods("GET")

	router.HandleFunc("/person/{id}", page.Delete_Person).Methods("DELETE")

	router.HandleFunc("/person/update/{id}/{attrName}/{attr}", page.Update_Attr).Methods("PUT")
	router.HandleFunc("/person/pass/change/{user}/{pass}/{newPass}", page.Change_Password).Methods("PUT")
	router.HandleFunc("/person/pass/set/{user}/{pass}", page.Set_Password).Methods("PUT")
	router.HandleFunc("/person/pass/reset/{user}", page.Reset_Password).Methods("PUT")
	router.HandleFunc("/person/{id}", page.Unsubscribe).Methods("PUT")
	//------------------------------------------------------------------------------------//
	//------------------------USER--------------------------------------------------------//
	//------------------------------------------------------------------------------------//
	router.HandleFunc("/user/login/{user}/{pass}", page.Login).Methods("POST")

	router.HandleFunc("/user/", page.User_Is_Logged_In).Methods("GET")
	router.HandleFunc("/user/logout/", page.Logout).Methods("PUT")

	//------------------------------------------------------------------------------------//
	//------------------------CATEGORIA---------------------------------------------------//
	//------------------------------------------------------------------------------------//
/*	router.HandleFunc("/categoria/", page.GetCategorias).Methods("GET")
	router.HandleFunc("/categoria/{id}", page.GetCategoriaById).Methods("GET")
	router.HandleFunc("/categoria/", page.InsertCategoria).Methods("POST")
	router.HandleFunc("/categoria/{id}", page.DeleteCategoria).Methods("DELETE")
	router.HandleFunc("/categoria/", page.UpdateCategoria).Methods("PUT")

	//------------------------------------------------------------------------------------//
	//------------------------PRODUCTO----------------------------------------------------//
	//------------------------------------------------------------------------------------//
	router.HandleFunc("/producto/", page.GetProductos).Methods("GET")
	router.HandleFunc("/producto/{id}", page.GetProductoById).Methods("GET")
	router.HandleFunc("/producto/", page.InsertProducto).Methods("POST")
	router.HandleFunc("/producto/{id}", page.DeleteProducto).Methods("DELETE")
	router.HandleFunc("/producto/", page.UpdateProducto).Methods("PUT")

	//------------------------------------------------------------------------------------//
	//------------------------CAJA--------------------------------------------------------//
	//------------------------------------------------------------------------------------//
	router.HandleFunc("/caja/", page.InsertCaja).Methods("POST")
	router.HandleFunc("/caja/", page.CerrarCaja).Methods("PUT")
	router.HandleFunc("/caja/", page.GetCajas).Methods("GET")
	router.HandleFunc("/caja/open/", page.GetCaja).Methods("GET")
	router.HandleFunc("/caja/export/{fechaInicio}/{fechaFin}/", page.ExportByFecha).Methods("GET")
	router.HandleFunc("/caja/efectivo/{id}", page.GetIngresosEfectivo).Methods("GET")
	router.HandleFunc("/caja/debito/{id}", page.GetIngresosDebito).Methods("GET")
	router.HandleFunc("/caja/credito/{id}", page.GetIngresosCredito).Methods("GET")

	//------------------------------------------------------------------------------------//
	//------------------------FACTURA-----------------------------------------------------//
	//------------------------------------------------------------------------------------//
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

	//------------------------------------------------------------------------------------//
	//------------------------IMAGE-------------------------------------------------------//
	//------------------------------------------------------------------------------------//
	router.HandleFunc("/upload/", page.Upload).Methods("POST")
	router.HandleFunc("/upload/", page.Download).Methods("GET")
*/
	//------------------------------------------------------------------------------------//
	//------------------------REQUEST-----------------------------------------------------//
	//------------------------------------------------------------------------------------//
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})

	//------------------------------------------------------------------------------------//
	//------------------------SERVER------------------------------------------------------//
	//------------------------------------------------------------------------------------//
	log.Print("Version 1.2")
	log.Print("Escuchando en localhost:3000")
	log.Fatal(http.ListenAndServe("localhost:3000", handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(router)))
}