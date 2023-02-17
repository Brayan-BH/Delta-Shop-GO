package tables

import (
	"delta/src/database/models"
)

func Productos_GetSchema() ([]models.Base, string) {
	var Productos []models.Base
	tableName := "Productos"
	Productos = append(Productos, models.Base{
		Name:        "idProducto",
		Description: "idProducto",
		Important:   true,
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Number(),
		},
	})
	Productos = append(Productos, models.Base{
		Name:        "nombre",
		Description: "nombre",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       5,
			Max:       100,
			LowerCase: true,
		},
	})
	Productos = append(Productos, models.Base{
		Name:        "nivel",
		Description: "nivel",
		Required:    true,
		Type:        "uint64",
		Uint: models.Uints{
			Max: 9,
		},
	})
	Productos = append(Productos, models.Base{
		Name:        "descripcion",
		Description: "descripcion",
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       10,
			Max:       200,
			LowerCase: true,
		},
	})
	Productos = append(Productos, models.Base{
		Name:        "modelo",
		Description: "modelo",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       5,
			Max:       50,
			LowerCase: true,
		},
	})
	Productos = append(Productos, models.Base{
		Name:        "color",
		Description: "color",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       5,
			Max:       15,
			LowerCase: true,
		},
	})
	Productos = append(Productos, models.Base{
		Name:        "nombreImagen",
		Description: "nombreImagen",
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       4,
			Max:       50,
			LowerCase: true,
		},
	})
	Productos = append(Productos, models.Base{
		Name:        "idCategoria",
		Description: "idCategoria",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       1,
			Max:       10,
			LowerCase: true,
		},
	})
	Productos = append(Productos, models.Base{
		Name:        "preciolista",
		Description: "preciolista",
		Required:    true,
		Update:      true,
		Type:        "float64",
	})
	Productos = append(Productos, models.Base{
		Name:        "codStripe",
		Description: "codStripe",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
		},
	})
	return Productos, tableName
}
