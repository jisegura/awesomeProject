package page

/*
//SELECT
func GetCategorias(w http.ResponseWriter, req *http.Request) {

	categoriaDAO := factory.FactoryCategoria()

	categorias, err := categoriaDAO.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	json.NewEncoder(w).Encode(&categorias)
}

//INSERT
func InsertCategoria(w http.ResponseWriter, req *http.Request) {

	categoriaDAO := factory.FactoryCategoria()
	categoria := models.Categoria{}

	_ = json.NewDecoder(req.Body).Decode(&categoria)
	err := categoriaDAO.Create(&categoria)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	cat, err := categoriaDAO.GetById(categoria.Id_categoria)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	json.NewEncoder(w).Encode(&cat)
}

//SELECT BY ID
func GetCategoriaById(w http.ResponseWriter, req *http.Request) {

	categoriaDAO := factory.FactoryCategoria()
	param := mux.Vars(req)

	i, _ := strconv.Atoi(param["id"])

	categoria, err := categoriaDAO.GetById(i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	json.NewEncoder(w).Encode(categoria)
}

//DELETE
func DeleteCategoria(w http.ResponseWriter, req *http.Request) {

	categoriaDAO := factory.FactoryCategoria()
	param := mux.Vars(req)

	i, _ := strconv.Atoi(param["id"])

	err := categoriaDAO.Delete(i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	cat := models.Categoria{}

	json.NewEncoder(w).Encode(&cat)
}

//UPDATE
func UpdateCategoria(w http.ResponseWriter, req *http.Request) {

	categoriaDAO := factory.FactoryCategoria()
	categoria := models.Categoria{}

	_ = json.NewDecoder(req.Body).Decode(&categoria)

	err := categoriaDAO.Update(&categoria)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	cat, err := categoriaDAO.GetById(categoria.Id_categoria)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	json.NewEncoder(w).Encode(&cat)
}
*/