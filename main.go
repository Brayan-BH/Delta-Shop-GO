package main

import (
	"delta/src/routes/routes_admin"
	"delta/src/routes/routes_clientes"
	routes "delta/src/routes/store"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	enableCORS(r)
	r.HandleFunc("/", HomeHandler)
	//servir imagenes estaticas
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	// r.HandleFunc("/static/img", func(w http.ResponseWriter, r *http.Request) {
	// 	file, err := ioutil.ReadFile("./public/b6d71ada-82cb-49f6-a30e-bdc5ba8df1b3logo.jpeg")
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	w.Write(file)
	// })

	//rutas de autentificacion\
	routes_clientes.RutasAuth(r)
	routes_clientes.RutasCliente(r)
	routes.RutasProductos(r)
	routes_admin.RutasProductos(r)
	fmt.Println("Server on port 5000")
	log.Fatal(http.ListenAndServe(":5000", r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	data := map[string]interface{}{"api": "apiDelta", "version": 1.1}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func enableCORS(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
	}).Methods(http.MethodOptions)
	router.Use(middlewareCors)
}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			// Just put some headers to allow CORS...
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Auth-Date, Auth-Periodo, Access-Token")
			// and call next handler!
			next.ServeHTTP(w, req)
		})
}
