package models

import (
	"github.com/graphql-go/graphql"
	"time"
)

// Carrera type definition
type Carrera struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Nombre    string     `json:"nombre" gorm:"type:varchar(128); not null; unique"`
	Logo      string     `json:"logo" gorm:"type:varchar(128)"`
	Creacion  uint16     `json:"creacion"`
	Estado    bool       `json:"estado" gorm:"default:'true'"`
	FilialID  uint       `json:"filial_id"`
	Semestres []Semestre `json:"semestres,omitempty"`
}

// Database custom table name
func (Carrera) TableName() string {
	return "carreras"
}

// Model GraphQL
var CarreraType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Carrera",
		Fields: graphql.Fields{
			"id":         &graphql.Field{Type: graphql.Int},
			"created_at": &graphql.Field{Type: graphql.DateTime},
			"updated_at": &graphql.Field{Type: graphql.DateTime},
			"deleted_at": &graphql.Field{Type: graphql.DateTime},
			"nombre":     &graphql.Field{Type: graphql.String},
			"logo":       &graphql.Field{Type: graphql.String},
			"creacion":   &graphql.Field{Type: graphql.Int},
			"filial_id":  &graphql.Field{Type: graphql.Int},
			"estado":     &graphql.Field{Type: graphql.Boolean},
			"semestres":  &graphql.Field{Type: graphql.NewList(SemestreType)},
		},
	},
)
