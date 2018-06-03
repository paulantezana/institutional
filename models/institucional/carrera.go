package institucional

import (
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

// Carrera type definition
type Carrera struct {
	gorm.Model
	Nombre      string `json:"nombre" gorm:"not null; unique"`
	Logo        string `json:"logo"`
	Descripcion string `json:"descripcion"`
	Creacion    int32  `json:"creacion"`
	Estado      bool   `json:"estado" gorm:"default:'true'"`

	FilialID uint `json:"filial_id"`
	Semestres []Semestre `json:"semestres,omitempty"`
}

func (Carrera) TableName() string {
    return "carreras"
}

// Model GraphQL
var CarreraType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Carrera",
		Fields: graphql.Fields{
			"nombre":      &graphql.Field{Type: graphql.String},
			"logo":        &graphql.Field{Type: graphql.String},
			"descripcion": &graphql.Field{Type: graphql.String},
			"creacion":    &graphql.Field{Type: graphql.Int},
			"estado":      &graphql.Field{Type: graphql.Boolean},
		},
	},
)
