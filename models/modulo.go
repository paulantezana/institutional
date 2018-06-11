package models

import (
	"github.com/graphql-go/graphql"
	"time"
)

// Modulo type definition
type Modulo struct {
	ID          uint       `json:"id" gorm:"primary_key"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	Nombre      string     `json:"nombre" gorm:"type:varchar(128); not null"`
	Tipo        string     `json:"tipo"`
	Descripcion string     `json:"descripcion"`
	Estado      bool       `json:"estado" gorm:"default:'true'"`
	SemestreID  uint       `json:"semestre_id"`
	Unidades    []Unidad   `json:"unidades,omitempty"`
}

// Database custom table name
func (Modulo) TableName() string {
	return "modulos"
}

// Model GraphQL
var ModuloType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Modulo",
		Fields: graphql.Fields{
			"id":          &graphql.Field{Type: graphql.Int},
			"created_at":  &graphql.Field{Type: graphql.DateTime},
			"updated_at":  &graphql.Field{Type: graphql.DateTime},
			"deleted_at":  &graphql.Field{Type: graphql.DateTime},
			"nombre":      &graphql.Field{Type: graphql.String},
			"tipo":        &graphql.Field{Type: graphql.String},
			"descripcion": &graphql.Field{Type: graphql.String},
			"semestre_id": &graphql.Field{Type: graphql.Int},
			"estado":      &graphql.Field{Type: graphql.Boolean},
            "unidades":      &graphql.Field{Type: graphql.NewList(UnidadType)},
		},
	},
)
