package tables
import "server-go/src/models" 
func Productos_GetSchema_DB() ([]models.Base_DB, string) {
	var Productos []models.Base_DB
	tableName := "_" + "Productos"
	Productos = append(Productos, models.Base_DB{
		Name:"idProducto",
		Description:"idProducto",
		Required: true,
		Update: true,
		Type:"string",
		Strings: models.Strings{
			Expr:      *models.Null(),
		},
	})
	Productos = append(Productos, models.Base_DB{
		Name:"nombre",
		Description:"nombre",
		Required: true,
		Update: true,
		Type:"string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:10.000000,
			Max:100,
			UpperCase:true,
		},
	})
	Productos = append(Productos, models.Base_DB{
		Name:"nivel",
		Description:"nivel",
		Required: true,
		Update: true,
		Type:"uint64",
		Uint: models.Uints{
			Max: 10,
		},
	})
	Productos = append(Productos, models.Base_DB{
		Name:"descripcion",
		Description:"descripcion",
		Required: true,
		Update: true,
		Type:"string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:20.000000,
			Max:200,
			UpperCase:true,
		},
	})
	Productos = append(Productos, models.Base_DB{
		Name:"modelo",
		Description:"modelo",
		Required: true,
		Update: true,
		Type:"string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:10.000000,
			Max:100,
			UpperCase:true,
		},
	})
	Productos = append(Productos, models.Base_DB{
		Name:"color",
		Description:"color",
		Required: true,
		Update: true,
		Type:"string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:10.000000,
			Max:100,
			UpperCase:true,
		},
	})
	Productos = append(Productos, models.Base_DB{
		Name:"nombreImagen",
		Description:"nombreImagen",
		Required: true,
		Update: true,
		Type:"string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:4.000000,
			Max:40,
			UpperCase:true,
		},
	})
	Productos = append(Productos, models.Base_DB{
		Name:"idCategoria",
		Description:"idCategoria",
		Required: true,
		Update: true,
		Type:"string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:0.300000,
			Max:3,
			UpperCase:true,
		},
	})
	Productos = append(Productos, models.Base_DB{
		Name:"preciolista",
		Description:"preciolista",
		Required: true,
		Update: true,
	})
	Productos = append(Productos, models.Base_DB{
		Name:"codStripe",
		Description:"codStripe",
		Required: true,
		Update: true,
		Type:"string",
		Strings: models.Strings{
			Expr:      *models.Null(),
		},
	})
	return Productos, tableName
}
