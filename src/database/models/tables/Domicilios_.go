package tables

import "delta/src/database/models"

func Domicilios_GetSchema() ([]models.Base, string) {
	var Domicilios []models.Base
	tableName := "Domicilios"
	Domicilios = append(Domicilios, models.Base{
		Name:        "idDomicilio",
		Description: "idDomicilio",
		Important:   true,
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
		},
	})
	Domicilios = append(Domicilios, models.Base{
		Name:        "idPedido",
		Description: "idPedido",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       5,
			Max:       8,
			LowerCase: true,
		},
	})
	Domicilios = append(Domicilios, models.Base{
		Name:        "direccion",
		Description: "direccion",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       10,
			Max:       50,
			LowerCase: true,
		},
	})
	Domicilios = append(Domicilios, models.Base{
		Name:        "distrito",
		Description: "distrito",
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
	Domicilios = append(Domicilios, models.Base{
		Name:        "provincia",
		Description: "provincia",
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
	Domicilios = append(Domicilios, models.Base{
		Name:        "departamento",
		Description: "departamento",
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
	Domicilios = append(Domicilios, models.Base{
		Name:        "ubigeo",
		Description: "ubigeo",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
			Min:  6,
			Max:  6,
		},
	})
	Domicilios = append(Domicilios, models.Base{
		Name:        "referencia",
		Description: "referencia",
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
	return Domicilios, tableName
}
