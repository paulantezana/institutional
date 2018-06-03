package institucional

import (
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

// Unidad type definition
type Unidad struct {
	gorm.Model
	Nombre  string `json:"nombre" gorm:"not null"`
	Credito int32  `json:"credito" gorm:"not null"`
	Horas   int32  `json:"horas" gorm:"not null"`
	Estado  bool   `json:"estado" gorm:"default:'true'"`
	ModuloID uint `json:"modulo_id"`
}

func (Unidad) TableName() string {
	return "unidades"
}

// Model GraphQL
var UnidadType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Unidad",
		Fields: graphql.Fields{
			"nombre":  &graphql.Field{Type: graphql.String},
			"credito": &graphql.Field{Type: graphql.Int},
			"horas":   &graphql.Field{Type: graphql.Int},
			"estado":  &graphql.Field{Type: graphql.Boolean},
		},
	},
)
