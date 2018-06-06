package models

import (
	"github.com/graphql-go/graphql"
	"time"
)

// Beca type definition
type Beca struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Beca      string     `json:"beca"`
	Estado    bool       `json:"estado" gorm:"default:'true'"`
}

// Database custom table name
func (Beca) TableName() string {
	return "becas"
}

var BecaType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "",
		Fields: graphql.Fields{
			"id":         &graphql.Field{Type: graphql.Int},
			"created_at": &graphql.Field{Type: graphql.DateTime},
			"updated_at": &graphql.Field{Type: graphql.DateTime},
			"deleted_at": &graphql.Field{Type: graphql.DateTime},
		},
	},
)
