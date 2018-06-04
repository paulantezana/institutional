package models

import (
	"github.com/graphql-go/graphql"
	"time"
)

type AlumnoEstado struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Nombre    string     `json:"nombre" gorm:"type:varchar(32); not null; unique"`
	Estado    bool       `json:"estado" gorm:"default:'true'"`
}

func (AlumnoEstado) TableName() string {
	return "alumno_estados"
}

var AlumnoEstadoType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "",
		Fields: graphql.Fields{
			"id":         &graphql.Field{Type: graphql.Int},
			"created_at": &graphql.Field{Type: graphql.DateTime},
			"updated_at": &graphql.Field{Type: graphql.DateTime},
			"deleted_at": &graphql.Field{Type: graphql.DateTime},
			"nombre":     &graphql.Field{Type: graphql.String},
			"estado":     &graphql.Field{Type: graphql.Boolean},
		},
	},
)
