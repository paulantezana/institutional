package institucional

import (
	"github.com/graphql-go/graphql"
	"time"
)

// Instituto type definition
type Instituto struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Estado    bool       `json:"estado" gorm:"default:'true'"`

	Filiales []Filial `json:"filiales,omitempty"`
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
			"id":         &graphql.Field{Type: graphql.Int},
			"created_at": &graphql.Field{Type: graphql.DateTime},
			"updated_at": &graphql.Field{Type: graphql.DateTime},
			"deleted_at": &graphql.Field{Type: graphql.DateTime},
		},
	},
)
