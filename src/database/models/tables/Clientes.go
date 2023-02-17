package tables

import (
	"delta/src/database/models"

	"github.com/google/uuid"
)

func Clientes_GetSchema() ([]models.Base, string) {
	var Clientes []models.Base
	tableName := "Clientes"
	Clientes = append(Clientes, models.Base{
		Name:        "idCliente",
		Description: "idCliente",
		Important:   true,
		Required:    true,
		Default:     uuid.New().String(),
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
		},
	})
	Clientes = append(Clientes, models.Base{
		Name:        "nombres",
		Description: "nombres",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       3,
			Max:       100,
			LowerCase: true,
		},
	})
	Clientes = append(Clientes, models.Base{
		Name:        "apellidos",
		Description: "apellidos",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       5,
			Max:       100,
			LowerCase: true,
		},
	})
	Clientes = append(Clientes, models.Base{
		Name:        "email",
		Description: "email",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       14,
			Max:       80,
			LowerCase: true,
		},
	})
	Clientes = append(Clientes, models.Base{
		Name:        "genero",
		Description: "genero",
		Type:        "uint64",
		Uint: models.Uints{
			Max: 10,
		},
	})
	Clientes = append(Clientes, models.Base{
		Name:        "passwords",
		Description: "passwords",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       8,
			Max:       100,
			LowerCase: true,
		},
	})
	Clientes = append(Clientes, models.Base{
		Name:        "telefono",
		Description: "telefono",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       6,
			Max:       9,
			LowerCase: true,
		},
	})
	return Clientes, tableName
}
