package routes_clientes

import (
	"delta/src/controller"
	"delta/src/database/models/tables"
	"delta/src/library/lib"
	"delta/src/library/sqlquery"
	"delta/src/middleware"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func RutasCliente(r *mux.Router) {

	s := r.PathPrefix("/cliente").Subrouter()
	s.Handle("/get/info-cls-a/data/", middleware.Autentication(http.HandlerFunc(allCliente))).Methods("GET")
	s.Handle("/get/info-cla-o/data/{idCliente}", middleware.Autentication(http.HandlerFunc(oneCLiente))).Methods("GET")
	s.Handle("/update/info-reg-o/data/{idCliente}", middleware.Autentication(http.HandlerFunc(updateCliente))).Methods("PUT")
	s.Handle("/get/list-domicilio/", middleware.Autentication(http.HandlerFunc(listDomicilio))).Methods("GET")
	s.Handle("/create/cli-domicilio/", middleware.Autentication(http.HandlerFunc(insertDomicilio))).Methods("POST")
	s.Handle("/create/info-reg-o/data/", middleware.Autentication(http.HandlerFunc(insertCliente))).Methods("POST")
	s.Handle("/update/reg-cli-domicilio/", middleware.Autentication(http.HandlerFunc(updateDomicilio))).Methods("PUT")

}

func allCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()

	//get allData from database
	dataCliente := sqlquery.NewQuerys("Clientes").Select("idCliente,nombres,apellidos,email,genero,passwords,telefono").Exec().All()
	response.Data["clientes"] = dataCliente
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func updateCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()
	idCliente := lib.GetId(r)
	if idCliente == "" {
		response.Msg = "Error al actualiar el cliente"
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

	data_body["where"] = map[string]interface{}{"idCliente": idCliente}
	var data_update []map[string]interface{}
	data_update = append(data_update, data_body)

	schema, table := tables.Clientes_GetSchema()
	cliente := sqlquery.SqlLibExec{}
	err = cliente.New(data_update, table).Update(schema)
	if err != nil {
		response.Msg = err.Error()
		response.StatusCode = 300
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = cliente.Exec()
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

func oneCLiente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()
	idCliente := lib.GetId(r)
	if idCliente == "" {
		response.Msg = "Error to write Cliente"
		response.StatusCode = 400
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}
	//get allData from database
	dataCliente := sqlquery.NewQuerys("Clientes").Select("idCLiente,nombres,apellidos,email,genero,passwords,telefono").Where("idCliente", "=", idCliente).Exec().One()
	response.Data = dataCliente
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func insertDomicilio(w http.ResponseWriter, r *http.Request) {
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

	schema, table := tables.Domicilios_GetSchema()
	domicilio := sqlquery.SqlLibExec{}
	err = domicilio.New(data_insert, table).Insert(schema)
	if err != nil {
		response.Msg = err.Error()
		response.StatusCode = 300
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = domicilio.Exec()
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

func updateDomicilio(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()
	idDomicilio := lib.GetId(r)

	if idDomicilio == "" {
		response.Msg = "Error al actualiar el domicilio"
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
		response.Msg = "No se encontraron datos para actualizar el domicilio"
		response.StatusCode = 400
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	data_body["where"] = map[string]interface{}{"idDomicilio": idDomicilio}
	var data_update []map[string]interface{}
	data_update = append(data_update, data_body)

	schema, table := tables.Domicilios_GetSchema()
	domicilio := sqlquery.SqlLibExec{}
	err = domicilio.New(data_update, table).Update(schema)
	if err != nil {
		response.Msg = err.Error()
		response.StatusCode = 300
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = domicilio.Exec()
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

func listDomicilio(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()
	idDomicilio := lib.GetId(r)
	if idDomicilio == "" {
		response.Msg = "Error to write Domicilios"
		response.StatusCode = 400
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}
	//get allData from database
	dataDomicilio := sqlquery.NewQuerys("Domicilios").Select("idCLiente,idPedido,direccion,distrito,provincia,departamento,ubigeo,referencia").Where("idDomicilio", "=", idDomicilio).Exec().One()
	response.Data = dataDomicilio
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func insertCliente(w http.ResponseWriter, r *http.Request) {
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

	schema, table := tables.Clientes_GetSchema()
	cliente := sqlquery.SqlLibExec{}
	err = cliente.New(data_insert, table).Insert(schema)
	if err != nil {
		response.Msg = err.Error()
		response.StatusCode = 300
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = cliente.Exec()
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
