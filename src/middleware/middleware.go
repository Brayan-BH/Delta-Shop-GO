package middleware

import (
	"delta/src/controller"
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Jwtclaim struct {
	Email     string `json:"email"`
	Nombres   string `json:"nombres"`
	IdCliente string `json:"id"`
	jwt.StandardClaims
}

//Seguridad Autentication
func Autentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := controller.NewResponseManager()
		//recibir token
		access_token := r.Header.Get("Access-Token")
		if access_token == "" {
			response.Msg = "Access token is missing"
			response.StatusCode = 405
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			return
		} else {

			//ValidateToken
			key_token := []byte("deltacommerce")
			token, err := jwt.ParseWithClaims(access_token, &Jwtclaim{}, func(tk *jwt.Token) (interface{}, error) {
				return key_token, nil
			})

			if err != nil {
				response.Msg = "Error signing" + err.Error()
				response.StatusCode = 300
				response.Status = "Error"
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(response)
				return
			}
			claims, ok := token.Claims.(*Jwtclaim)
			if !ok {
				response.Msg = "Error signing"
				response.StatusCode = 300
				response.Status = "Error"
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(response)
				return
			}

			if claims.ExpiresAt < time.Now().Local().Unix() {
				response.Msg = "Session Expired"
				response.StatusCode = 300
				response.Status = "Error"
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(response)
				return
			}
			// fmt.Print(Jwtclaim{})
			//valida autenticacion json
			next.ServeHTTP(w, r)
		}
	})
}
