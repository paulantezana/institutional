package institucional

import (
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

// Semestre type definition
type Semestre struct {
	gorm.Model
	Nombre string `json:"nombre" gorm:"not null"`
	Year   int32  `json:"year"`
	Estado bool   `json:"estado" gorm:"default:'true'"`

	CarreraID uint `json:"carrera_id"`
	Modulos []Modulo `json:"modulos,omitempty"`
}

func (Semestre) TableName() string {
    return "semestres"
}

// Model GraphQL
var SemestreType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Semestre",
		Fields: graphql.Fields{
			"nombre": &graphql.Field{Type: graphql.String},
			"year":   &graphql.Field{Type: graphql.Int},
			"estado": &graphql.Field{Type: graphql.Boolean},
		},
	},
)
