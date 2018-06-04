package models

import (
	"github.com/graphql-go/graphql"
	"time"
)

type Core struct {
	ID          uint       `json:"id" gorm:"primary_key"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	Name        string     `json:"name" gorm:"type:varchar(128); not null"`
	Description string     `json:"description" `
	Version     string     `json:"version" gorm:"type:varchar(128)"`
	Author      string     `json:"author" gorm:"type:varchar(128)"`
	AuthorWeb   string     `json:"author_web" gorm:"type:varchar(128)"`
}

// Database custom table name
func (Core) TableName() string {
	return "core"
}

var CoreType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Core",
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
		},
	},
)
