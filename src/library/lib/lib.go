package lib

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

type Jwtclaim struct {
	Email     string `json:"email"`
	Nombres   string `json:"nombres"`
	IdCliente string `json:"id"`
	jwt.StandardClaims
}

//typeof
//recibe un valor interface que no se reconoce su tipo y devuelve un string
func InterfaceToString(params ...interface{}) string {
	typeValue := reflect.TypeOf(params[0]).String()
	value := params[0]
	valueReturn := ""
	if strings.Contains(typeValue, "string") {
		toSql := false
		if len(params) == 2 && reflect.TypeOf(params[1]).Kind() == reflect.Bool {
			toSql = params[1].(bool)
		}

		if toSql {
			valueReturn = fmt.Sprintf("'%s'", value)
		} else {
			valueReturn = fmt.Sprintf("%s", value)
		}
	} else if strings.Contains(typeValue, "int") {
		valueReturn = fmt.Sprintf("%d", value)
	} else if strings.Contains(typeValue, "float") {
		valueReturn = fmt.Sprintf("%f", value)
	} else if strings.Contains(typeValue, "bool") {
		valueReturn = fmt.Sprintf("%t", value)
	}
	return valueReturn
}

func BytesToFloat64(bytes []byte) float64 {

	text := bytes // A decimal value represented as Latin-1 text

	f, err := strconv.ParseFloat(string(text), 64)
	if err != nil {
		fmt.Print("Error Conv:", err)
	}

	return f
}

func IndexOfStrings(data []string, key string) int {
	for i, v := range data {
		if v == key {
			return i
		}
	}
	return -1
}

func GetId(r *http.Request) string {
	key_token := []byte("deltacommerce")

	access_token := r.Header.Get("Access-Token")
	token, _ := jwt.ParseWithClaims(access_token, &Jwtclaim{}, func(tk *jwt.Token) (interface{}, error) {
		return key_token, nil
	})

	claims, _ := token.Claims.(*Jwtclaim)

	return claims.IdCliente
}
