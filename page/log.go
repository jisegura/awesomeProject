package page

/*
func Create_Log (w http.ResponseWriter, req *http.Request) {

	logDAO := factory.FactoryLog()
	param := mux.Vars(req)
	pass := param["pass"]

	err := logDAO.Create_Log(pass)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Error: ", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

*/