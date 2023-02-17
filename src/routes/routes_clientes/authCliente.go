package routes_clientes

import (
	"delta/src/controller"
	"delta/src/library/sqlquery"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type jwtclaim struct {
	Email     string `json:"email"`
	Nombres   string `json:"nombres"`
	IdCliente string `json:"id"`
	jwt.StandardClaims
}

func RutasAuth(r *mux.Router) {

	s := r.PathPrefix("/auth-cliente").Subrouter()
	s.HandleFunc("/", auth).Methods("GET")
	s.HandleFunc("/login", login).Methods("PUT")
	s.HandleFunc("/create/info-reg-o/data/", insertCliente).Methods("POST")

}

func auth(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "Aplication-Json")
	response := controller.NewResponseManager()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func login(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "Aplication-Json")
	response := controller.NewResponseManager()

	// Get the request body
	req_body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Msg = err.Error()
		response.StatusCode = 300
		response.Status = "Error"
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}
	//objeto map
	body := make(map[string]interface{})
	//leer the body y convertir en objeto
	json.Unmarshal(req_body, &body)

	dataUser := sqlquery.NewQuerys("Clientes").Select().Where("email", "=", body["email"]).Exec().One()
	if len(dataUser) <= 0 {
		response.Msg = "Email y contrasenia Incorrecto"
		response.StatusCode = 300
		response.Status = "Email y contrasenia Incorrecto"
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}
	//int8, int64, int32

	err = bcrypt.CompareHashAndPassword([]byte(dataUser["passwords"].(string)), []byte(body["password"].(string)))
	if err != nil {
		response.Msg = "Email Incorrecto o Contrasenia Incorrecto"
		response.StatusCode = 300
		response.Status = "Error"
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	//Tokens
	//tiempo de expiracion del token
	var key_token interface{}
	key_token = []byte("deltacommerce")
	claims := jwtclaim{
		dataUser["email"].(string),
		dataUser["nombres"].(string),
		dataUser["idCliente"].(string),
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * (60 * 24)).Unix(),
			Issuer:    "pdt",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token_string, err_token := token.SignedString(key_token)
	if err_token != nil {
		response.Msg = "Error signing" + err_token.Error()
		response.StatusCode = 300
		response.Status = "Error"
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Data["token"] = token_string
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
