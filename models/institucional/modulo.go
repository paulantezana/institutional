package institucional

import (
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

// Modulo type definition
type Modulo struct {
	gorm.Model
	Nombre      string `json:"nombre" gorm:"not null"`
	Tipo        string `json:"tipo"`
	Descripcion string `json:"descripcion"`
	Estado      bool   `json:"estado" gorm:"default:'true'"`

	SemestreID uint `json:"semestre_id"`
	Unidades []Unidad `json:"unidades,omitempty"`
}

func (Modulo) TableName() string {
	return "modulos"
}

// Model GraphQL
var ModuloType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Modulo",
		Fields: graphql.Fields{
			"nombre":      &graphql.Field{Type: graphql.String},
			"tipo":        &graphql.Field{Type: graphql.String},
			"descripcion": &graphql.Field{Type: graphql.String},
			"estado":      &graphql.Field{Type: graphql.Boolean},
		},
	},
)
