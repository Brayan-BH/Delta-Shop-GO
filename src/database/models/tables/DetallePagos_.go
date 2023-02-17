package tables

import "delta/src/database/models"

func DetallePagos_GetSchema() ([]models.Base, string) {
	var DetallePagos []models.Base
	tableName := "DetallePagos"
	DetallePagos = append(DetallePagos, models.Base{
		Name:        "idDetallePago",
		Description: "idDetallePago",
		Important:   true,
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
		},
	})
	DetallePagos = append(DetallePagos, models.Base{
		Name:        "metodoPago",
		Description: "metodoPago",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       6,
			Max:       50,
			LowerCase: true,
		},
	})
	DetallePagos = append(DetallePagos, models.Base{
		Name:        "detallePago",
		Description: "detallePago",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       10,
			Max:       100,
			LowerCase: true,
		},
	})
	DetallePagos = append(DetallePagos, models.Base{
		Name:        "oservaciones",
		Description: "oservaciones",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       7,
			Max:       50,
			LowerCase: true,
		},
	})
	DetallePagos = append(DetallePagos, models.Base{
		Name:        "idPedido",
		Description: "idPedido",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
		},
	})
	return DetallePagos, tableName
}
