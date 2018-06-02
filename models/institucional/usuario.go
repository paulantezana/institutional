package institucional

import (
	"time"

	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

// Usuario type definition
type Usuario struct {
	gorm.Model
	Usuario                string    `json:"usuario"`
	Clave                  string    `json:"clave"`
	ClaveRecuperar         string    `json:"clave_recuperar"`
	FechaModificacionClave time.Time `json:"fecha_modificacion_clave"`
	Estado                 bool      `json:"estado" gorm:"default:'true'"`
}

// Model GraphQL
var UsuarioType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Usuario",
		Fields: graphql.Fields{
			"usuario":                  &graphql.Field{Type: graphql.String},
			"clave":                    &graphql.Field{Type: graphql.String},
			"clave_recuperar":          &graphql.Field{Type: graphql.String},
			"fecha_modificacion_clave": &graphql.Field{Type: graphql.DateTime},
			"estado":                   &graphql.Field{Type: graphql.Boolean},
		},
	},
)