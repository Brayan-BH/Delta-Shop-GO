package routes

import (
	"delta/src/controller"
	"delta/src/library/sqlquery"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func RutasProductos(r *mux.Router) {

	s := r.PathPrefix("/store").Subrouter()
	s.HandleFunc("/productos/list-store/{id}", allProductos).Methods("GET")
	s.HandleFunc("/productos/list-store-categoria", filterCategoria).Methods("GET")
}

func allProductos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()
	params := mux.Vars(r)
	id := params["id"]
	//get allData from database
	var dataProductos []map[string]interface{}
	if id == "all" {
		dataProductos = sqlquery.NewQuerys("Productos").Select("nombre,descripcion,modelo,color,preciolista,idCategoria,nombreImagen,codstripe").Where("nivel", "=", 3).Exec().All()
	} else {
		dataProductos = sqlquery.NewQuerys("Productos").Select("nombre,descripcion,modelo,color,preciolista,idCategoria,nombreImagen,codstripe").Where("nivel", "=", 3).AndLike("idProducto", id+"%").Exec().All()
	}

	response.Data["productosList"] = dataProductos
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func insertProductos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	params := mux.Vars(r)
	_ = params["idProducto"]
	file, _, err := r.FormFile("file")
	if err != nil {
		response.Msg = err.Error()
		response.StatusCode = 300
		json.NewEncoder(w).Encode(response)
		return
	}
	defer file.Close()

	fmt.Println(file)
}

func filterCategoria(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()
	// params := mux.Vars(r)
	// id := params["id"]
	//get allData from database
	dataProductos := sqlquery.NewQuerys("Productos").Select("nombre,nivel,idCategoria,idProducto").Where("nivel", "<>", 3).Exec().OrderBy("idProducto").All()
	response.Data["filterCategoria"] = dataProductos
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

