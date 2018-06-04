package models

import (
	"github.com/graphql-go/graphql"
	"time"
)

type CoreRol struct {
	ID          uint       `json:"id" gorm:"primary_key"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	Name        string     `json:"name" gorm:"type:varchar(128), not null"`
	Description string     `json:"description" `
	Freeze      bool       `json:"freeze"`
	Estado      bool       `json:"estado" gorm:"default:'true'"`
}

// Database custom table name
func (CoreRol) TableName() string {
	return "core_roles"
}

var CoreRolType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "CoreRol",
		Fields: graphql.Fields{
			"id":          &graphql.Field{Type: graphql.Int},
			"created_at":  &graphql.Field{Type: graphql.DateTime},
			"updated_at":  &graphql.Field{Type: graphql.DateTime},
			"deleted_at":  &graphql.Field{Type: graphql.DateTime},
			"name":        &graphql.Field{Type: graphql.String},
			"description": &graphql.Field{Type: graphql.String},
			"freeze":      &graphql.Field{Type: graphql.Boolean},
			"estado":      &graphql.Field{Type: graphql.Boolean},
		},
	},
)
