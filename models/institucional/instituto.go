package institucional

import (
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

// Instituto type definition
type Instituto struct {
	gorm.Model
	Estado bool `json:"estado" gorm:"default:'true'"`

	Filiales []Filial `json:"filiales,omitempty"`
}

// Model GraphQL
var InstitutoType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:   "Instituto",
		Fields: graphql.Fields{},
	},
)
