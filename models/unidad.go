package models

import (
	"github.com/graphql-go/graphql"
	"time"
)

// Unidad type definition
type Unidad struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Nombre    string     `json:"nombre" gorm:"type:varchar(128); not null"`
	Credito   float32    `json:"credito" gorm:"not null"`
	Horas     uint16     `json:"horas" gorm:"not null"`
	Estado    bool       `json:"estado" gorm:"default:'true'"`
	ModuloID  uint       `json:"modulo_id"`
}

// Database custom table name
func (Unidad) TableName() string {
	return "unidades"
}

// Model GraphQL
var UnidadType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Unidad",
		Fields: graphql.Fields{
			"id":         &graphql.Field{Type: graphql.Int},
			"created_at": &graphql.Field{Type: graphql.DateTime},
			"updated_at": &graphql.Field{Type: graphql.DateTime},
			"deleted_at": &graphql.Field{Type: graphql.DateTime},
			"nombre":     &graphql.Field{Type: graphql.String},
			"credito":    &graphql.Field{Type: graphql.Float},
			"horas":      &graphql.Field{Type: graphql.Int},
			"modulo_id":  &graphql.Field{Type: graphql.Int},
			"estado":     &graphql.Field{Type: graphql.Boolean},
		},
	},
)
