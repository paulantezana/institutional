package models

import (
	"github.com/graphql-go/graphql"
	"time"
)

type CorePermit struct {
	ID            uint       `json:"id" gorm:"primary_key"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
	Name          string     `json:"name" gorm:"type:varchar(128), not null"`
	CoreSubMenuID uint       `json:"core_sub_menu_id"`
	CarreraID     uint       `json:"carrera_id"`
	SemestreID    uint       `json:"semestre_id"`
	Permission    uint16     `json:"permission"`
	Estado        bool       `json:"estado" gorm:"default:'true'"`
}

// Database custom table name
func (CorePermit) TableName() string {
	return "core_permits"
}

var CorePermitType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "CorePermit",
		Fields: graphql.Fields{
			"id":               &graphql.Field{Type: graphql.Int},
			"created_at":       &graphql.Field{Type: graphql.DateTime},
			"updated_at":       &graphql.Field{Type: graphql.DateTime},
			"deleted_at":       &graphql.Field{Type: graphql.DateTime},
			"name":             &graphql.Field{Type: graphql.String},
			"core_sub_menu_id": &graphql.Field{Type: graphql.Int},
			"carrera_id":       &graphql.Field{Type: graphql.Int},
			"semestre_id":      &graphql.Field{Type: graphql.Int},
			"permission":       &graphql.Field{Type: graphql.Int},
			"estado":           &graphql.Field{Type: graphql.Boolean},
		},
	},
)
