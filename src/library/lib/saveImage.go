package lib

import (
	"io/ioutil"
	"mime/multipart"
	"strings"

	"github.com/google/uuid"
)

func SaveFile(file multipart.File, handle *multipart.FileHeader) (string, error) {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	uuid := uuid.New().String()
	temp := strings.Split(handle.Filename, ".")
	tipo := temp[len(temp)-1]
	//separar filename por si envian con puntos 
	filename := "public/" + uuid + "." + tipo

	err = ioutil.WriteFile(filename, data, 0666)
	if err != nil {
		return "", err
	}
	return filename, nil
}
