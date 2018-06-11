package models

import (
	"github.com/graphql-go/graphql"
	"time"
)

// Semestre type definition
type Semestre struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Nombre    string     `json:"nombre" gorm:"type:varchar(128); not null"`
	Year      uint16     `json:"year" gorm:"not null"`
	Estado    bool       `json:"estado" gorm:"default:'true'"`
	CarreraID uint       `json:"carrera_id"`
	Modulos   []Modulo   `json:"modulos,omitempty"`
}

// Database custom table name
func (Semestre) TableName() string {
	return "semestres"
}

// Model GraphQL
var SemestreType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Semestre",
		Fields: graphql.Fields{
			"id":         &graphql.Field{Type: graphql.Int},
			"created_at": &graphql.Field{Type: graphql.DateTime},
			"updated_at": &graphql.Field{Type: graphql.DateTime},
			"deleted_at": &graphql.Field{Type: graphql.DateTime},
			"nombre":     &graphql.Field{Type: graphql.String},
			"year":       &graphql.Field{Type: graphql.Int},
			"carrera_id": &graphql.Field{Type: graphql.Int},
			"estado":     &graphql.Field{Type: graphql.Boolean},
            "modulos":     &graphql.Field{Type: graphql.NewList(ModuloType)},
		},
	},
)
