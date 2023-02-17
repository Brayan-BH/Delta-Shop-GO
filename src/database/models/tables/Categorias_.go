package tables

import "delta/src/database/models"

func Categorias_GetSchema() ([]models.Base, string) {
	var Categorias []models.Base
	tableName := "Categorias"
	Categorias = append(Categorias, models.Base{
		Name:        "idCategoria",
		Description: "idCategoria",
		Required:    true,
		Important:   true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
		},
	})
	Categorias = append(Categorias, models.Base{
		Name:        "nombreCategoria",
		Description: "nombreCategoria",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       5,
			Max:       20,
			UpperCase: true,
		},
	})
	return Categorias, tableName
}
