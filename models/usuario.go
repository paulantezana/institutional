package models

import (
	"time"

	"github.com/graphql-go/graphql"
)

// Usuario type definition
type Usuario struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`

	Usuario                string     `json:"usuario" gorm:"type:varchar(64); unique"`
	Correo                 string     `json:"correo" gorm:"type:varchar(64); unique"`
	Avatar                 string     `json:"avatar,omitempty" gorm:"type:varchar(128)"`
	Clave                  string     `json:"clave,omitempty"`
	ClaveRecuperar         string     `json:"clave_recuperar,omitempty"`
	FechaRecuperacionClave time.Time  `json:"fecha_recuperacion_clave,omitempty"`
	FechaModificacionClave time.Time  `json:"fecha_modificacion_clave,omitempty"`
	Estado                 bool       `json:"estado,omitempty" gorm:"default:'true'"`
	Alumno                 []Alumno   `json:"alumno,omitempty"`
	Profesor               []Profesor `json:"profesor,omitempty"`

	ConfirmarClave string `json:"confirmar_clave,omitempty" gorm:"-"`
	ClaveAntigua   string `json:"clave_antigua,omitempty" gorm:"-"`
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
			"correo":                   &graphql.Field{Type: graphql.String},
			"clave":                    &graphql.Field{Type: graphql.String},
			"clave_recuperar":          &graphql.Field{Type: graphql.String},
			"fecha_modificacion_clave": &graphql.Field{Type: graphql.DateTime},
			"fecha_recuperacion_clave": &graphql.Field{Type: graphql.DateTime},
			"estado":                   &graphql.Field{Type: graphql.Boolean},
		},
	},
)
