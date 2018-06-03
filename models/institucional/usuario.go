package institucional

import (
	"time"

	"github.com/graphql-go/graphql"
)

// Usuario type definition
type Usuario struct {
	ID                     uint       `json:"id" gorm:"primary_key"`
	CreatedAt              time.Time  `json:"created_at"`
	UpdatedAt              time.Time  `json:"updated_at"`
	DeletedAt              *time.Time `json:"deleted_at"`
	Usuario                string     `json:"usuario"`
	Clave                  string     `json:"clave"`
	ConfirmarClave         string     `json:"confirmar_clave,omitempty" gorm:"-"`
	ClaveRecuperar         string     `json:"clave_recuperar"`
	FechaModificacionClave time.Time  `json:"fecha_modificacion_clave"`
	Estado                 bool       `json:"estado" gorm:"default:'true'"`
}

// Database custom table name
func (Usuario) TableName() string {
	return "usuarios"
}

// Model GraphQL
var UsuarioType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Usuario",
		Fields: graphql.Fields{
			"id":                       &graphql.Field{Type: graphql.Int},
			"created_at":               &graphql.Field{Type: graphql.DateTime},
			"updated_at":               &graphql.Field{Type: graphql.DateTime},
			"deleted_at":               &graphql.Field{Type: graphql.DateTime},
			"usuario":                  &graphql.Field{Type: graphql.String},
			"clave":                    &graphql.Field{Type: graphql.String},
			"clave_recuperar":          &graphql.Field{Type: graphql.String},
			"fecha_modificacion_clave": &graphql.Field{Type: graphql.DateTime},
			"estado":                   &graphql.Field{Type: graphql.Boolean},
		},
	},
)
