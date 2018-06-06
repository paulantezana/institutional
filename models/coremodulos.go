package models

import (
	"github.com/graphql-go/graphql"
	"time"
)

type CoreModulo struct {
	ID          uint       `json:"id" gorm:"primary_key"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	Name        string     `json:"name" gorm:"type:varchar(128); not null"`
	Description string     `json:"description" `
	Version     string     `json:"version" gorm:"type:varchar(128)"`
	Author      string     `json:"author" gorm:"type:varchar(128)"`
	AuthorWeb   string     `json:"author_web" gorm:"type:varchar(128)"`
	Style       string     `json:"style"`
	Script      string     `json:"script"`
	Config      string     `json:"config"`
    CoreRolID  uint       `json:"core_rol_id"`
	Estado      bool       `json:"estado" gorm:"default:'true'"`
}

// Database custom table name
func (CoreModulo) TableName() string {
	return "core_modulos"
}

var CoreModuloType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "CoreModulo",
		Fields: graphql.Fields{
			"id":          &graphql.Field{Type: graphql.Int},
			"created_at":  &graphql.Field{Type: graphql.DateTime},
			"updated_at":  &graphql.Field{Type: graphql.DateTime},
			"deleted_at":  &graphql.Field{Type: graphql.DateTime},
			"name":        &graphql.Field{Type: graphql.String},
			"description": &graphql.Field{Type: graphql.String},
			"version":     &graphql.Field{Type: graphql.String},
			"author":      &graphql.Field{Type: graphql.String},
			"author_web":  &graphql.Field{Type: graphql.String},
			"style":       &graphql.Field{Type: graphql.String},
			"script":      &graphql.Field{Type: graphql.String},
			"config":      &graphql.Field{Type: graphql.String},
            "core_rol_id":      &graphql.Field{Type: graphql.Int},
			"estado":      &graphql.Field{Type: graphql.Boolean},
		},
	},
)
