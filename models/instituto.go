package models

import (
	"github.com/graphql-go/graphql"
	"time"
)

// Instituto type definition
type Instituto struct {
	ID              uint       `json:"id" gorm:"primary_key"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at"`
	Nombre          string     `json:"nombre" gorm:"type:varchar(128); not null"`
	Logo            string     `json:"logo" gorm:"type:varchar(128)"`
	Director        string     `json:"director" gorm:"type:varchar(128); not null"`
	DREGRE          string     `json:"dregre" gorm:"type:varchar(128)"`
	SemestrePrefijo string     `json:"semestre_prefijo" gorm:"type:varchar(64)"`
	Semestres       uint16     `json:"semestres"`
	SemestreSufijo  string     `json:"semestre_sufijo" gorm:"type:varchar(64)"`
	Estado          bool       `json:"estado" gorm:"default:'true'"`
	Filiales        []Filial   `json:"filiales,omitempty"`
}

// Database custom table name
func (Instituto) TableName() string {
	return "institutos"
}

// Model GraphQL
var InstitutoType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Instituto",
		Fields: graphql.Fields{
			"id":               &graphql.Field{Type: graphql.Int},
			"created_at":       &graphql.Field{Type: graphql.DateTime},
			"updated_at":       &graphql.Field{Type: graphql.DateTime},
			"deleted_at":       &graphql.Field{Type: graphql.DateTime},
			"nombre":           &graphql.Field{Type: graphql.String},
			"logo":             &graphql.Field{Type: graphql.String},
			"director":         &graphql.Field{Type: graphql.String},
			"dregre":           &graphql.Field{Type: graphql.String},
			"semestre_prefijo": &graphql.Field{Type: graphql.String},
			"semestres":        &graphql.Field{Type: graphql.Int},
			"semestre_sufijo":  &graphql.Field{Type: graphql.String},
		},
	},
)
