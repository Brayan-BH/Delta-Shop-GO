package routes_admin

import (
	"delta/src/controller"
	"delta/src/database/models/tables"
	"delta/src/library/lib"
	"delta/src/library/sqlquery"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func RutasProductos(r *mux.Router) {

	s := r.PathPrefix("/admin").Subrouter()
	s.HandleFunc("/productos/create-img/{id}", insertProductosImagen).Methods("POST")
	s.HandleFunc("/productos/create/", insertProductos).Methods("POST")
	s.HandleFunc("/productos/list/{id}", allProductos).Methods("GET")
	s.HandleFunc("/productos/list-one-producto/{idProducto}", oneProducto).Methods("GET")
	s.HandleFunc("/update/productos/{idProducto}", updateProduct).Methods("PUT")

}

func insertProductosImagen(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := controller.NewResponseManager()

	params := mux.Vars(r)
	_ = params["idProducto"]
	file, handle, err := r.FormFile("foto")
	if err != nil {
		response.Msg = err.Error()
		response.StatusCode = 300
		json.NewEncoder(w).Encode(response)
		return
	}
	defer file.Close()

	mimeType := handle.Header.Get("Content-Type")
	fmt.Println(mimeType)
	path := ""
	fmt.Println(mimeType)
	switch mimeType {
	case "image/jpeg":
		path_file, err_saveFile := lib.SaveFile(file, handle)
		if err_saveFile != nil {
			response.Msg = err_saveFile.Error()
			response.StatusCode = 300
			json.NewEncoder(w).Encode(response)
			return
		}
		path = path_file
	case "image/png":
		path_file, err_saveFile := lib.SaveFile(file, handle)
		if err_saveFile != nil {
			response.Msg = err_saveFile.Error()
			response.StatusCode = 300
			json.NewEncoder(w).Encode(response)
			return
		}
		path = path_file
	default:
		response.Msg = "Tipo de Archivo no soportado"
		response.StatusCode = 300
		json.NewEncoder(w).Encode(response)
		return
	}

	fmt.Println(path)
}

func allProductos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()
	params := mux.Vars(r)
	id := params["id"]
	//get allData from database
	var dataProductos []map[string]interface{}
	if id == "all" {
		dataProductos = sqlquery.NewQuerys("Productos").Select("nombre,descripcion,modelo,color,preciolista,idCategoria,nombreImagen,nivel,codStripe").Where("nivel", "=", 3).Exec().All()
	} else {
		dataProductos = sqlquery.NewQuerys("Productos").Select("nombre,descripcion,modelo,color,preciolista,idCategoria,nombreImagen,,nivel,codStripe").Where("nivel", "=", 3).AndLike("idProducto", id+"%").Exec().All()
	}

	response.Data["productosList"] = dataProductos
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func insertProductos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()
	request_body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Msg = err.Error()
		response.StatusCode = 400
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	data_body := make(map[string]interface{})
	json.Unmarshal(request_body, &data_body)
	var data_insert []map[string]interface{}
	data_insert = append(data_insert, data_body)

	schema, table := tables.Productos_GetSchema()
	productos := sqlquery.SqlLibExec{}
	err = productos.New(data_insert, table).Insert(schema)
	if err != nil {
		response.Msg = err.Error()
		response.StatusCode = 300
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = productos.Exec()
	if err != nil {
		response.Msg = err.Error()
		response.StatusCode = 300
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()
	params := mux.Vars(r)
	idProducto := params["idProducto"]
	if idProducto == "" {
		response.Msg = "Error to write Product"
		response.StatusCode = 400
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	request_body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Msg = err.Error()
		response.StatusCode = 400
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	data_body := make(map[string]interface{})
	json.Unmarshal(request_body, &data_body)
	if len(data_body) <= 0 {
		response.Msg = "No se encontraron datos para actualizar"
		response.StatusCode = 400
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	data_body["where"] = map[string]interface{}{"idProducto": idProducto}
	var data_update []map[string]interface{}
	data_update = append(data_update, data_body)

	schema, table := tables.Productos_GetSchema()
	productos := sqlquery.SqlLibExec{}
	err = productos.New(data_update, table).Update(schema)
	if err != nil {
		response.Msg = err.Error()
		response.StatusCode = 300
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = productos.Exec()
	if err != nil {
		response.Msg = err.Error()
		response.StatusCode = 300
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func oneProducto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()
	// params := mux.Vars(r)
	// id := params["id"]
	//get allData from database
	dataProductos := sqlquery.NewQuerys("Productos").Select("idProducto,nombre,nivel,descripcion,modelo,color,preciolista,codStripe,idCategoria").Where("nivel", "=", 3).Exec().OrderBy("idProducto").One()
	response.Data["filterProducto"] = dataProductos
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
