package models

import (
	"github.com/graphql-go/graphql"
	"time"
)

type CoreSubMenu struct {
	ID         uint       `json:"id" gorm:"primary_key"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
	Name       string     `json:"name" gorm:"type:varchar(64); not null"`
	Icon       string     `json:"icon" gorm:"type:varchar(64)"`
	Url        string     `json:"url" gorm:"type:varchar(128); not null"`
	Target     string     `json:"target" gorm:"type:varchar(64); not null"`
	CoreMenuID uint       `json:"core_menu_id"`
	CarreraID  uint       `json:"carrera_id"`
	SemestreID uint       `json:"semestre_id"`
	Permission uint16     `json:"permission"`
	Estado     bool       `json:"estado" gorm:"default:'true'"`
}

// Database custom table name
func (CoreSubMenu) TableName() string {
	return "core_sub_menus"
}

var CoreSubMenuType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "CoreSubMenu",
		Fields: graphql.Fields{
			"id":           &graphql.Field{Type: graphql.Int},
			"created_at":   &graphql.Field{Type: graphql.DateTime},
			"updated_at":   &graphql.Field{Type: graphql.DateTime},
			"deleted_at":   &graphql.Field{Type: graphql.DateTime},
			"name":         &graphql.Field{Type: graphql.String},
			"icon":         &graphql.Field{Type: graphql.String},
			"url":          &graphql.Field{Type: graphql.String},
			"target":       &graphql.Field{Type: graphql.String},
			"core_menu_id": &graphql.Field{Type: graphql.Int},
			"carrera_id":   &graphql.Field{Type: graphql.Int},
			"semestre_id":  &graphql.Field{Type: graphql.Int},
			"permission":   &graphql.Field{Type: graphql.Int},
			"estado":       &graphql.Field{Type: graphql.Boolean},
		},
	},
)
