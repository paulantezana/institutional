package administracion

import (
	"github.com/graphql-go/graphql"
	"time"
)

type Personal struct {
	ID              uint       `json:"id" gorm:"primary_key"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at"`
	Dni             string     `json:"dni" gorm:"not null; unique"`
	Nombres         string     `json:"nombres" gorm:"not null"`
	ApellidoPaterno string     `json:"apellido_paterno" gorm:"not null"`
	ApellidoMaterno string     `json:"apellido_materno" gorm:"not null"`
	FechaNacimiento time.Time  `json:"fecha_nacimiento"`
	Sexo            string     `json:"sexo"`
	Pais            string     `json:"pais" gorm:"not null"`
	Departamento    string     `json:"departamento"`
	Direccion       string     `json:"direccion"`
	Celular         string     `json:"celular" gorm:"unique"`
	Telefono        string     `json:"telefono" gorm:"unique"`
	Correo          string     `json:"correo" gorm:"unique"`
	EstadoCivil     string     `json:"estado_civil"`
	Foto            string     `json:"foto"`
	FechaIngreso    time.Time  `json:"fecha_ingreso"`
	FechaRetiro     time.Time  `json:"fecha_retiro"`
	Estado          bool       `json:"estado" gorm:"default:'true'"`
}

// Database custom table name
func (Personal) TableName() string {
	return "personales"
}

// Model GraphQL
var PersonalType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Personal",
		Fields: graphql.Fields{
			"id":               &graphql.Field{Type: graphql.Int},
			"created_at":       &graphql.Field{Type: graphql.DateTime},
			"updated_at":       &graphql.Field{Type: graphql.DateTime},
			"deleted_at":       &graphql.Field{Type: graphql.DateTime},
			"dni":              &graphql.Field{Type: graphql.String},
			"nombres":          &graphql.Field{Type: graphql.String},
			"apellido_paterno": &graphql.Field{Type: graphql.String},
			"apellido_materno": &graphql.Field{Type: graphql.String},
			"fecha_nacimiento": &graphql.Field{Type: graphql.DateTime},
			"sexo":             &graphql.Field{Type: graphql.String},
			"pais":             &graphql.Field{Type: graphql.String},
			"departamento":     &graphql.Field{Type: graphql.String},
			"direccion":        &graphql.Field{Type: graphql.String},
			"celular":          &graphql.Field{Type: graphql.String},
			"telefono":         &graphql.Field{Type: graphql.String},
			"correo":           &graphql.Field{Type: graphql.String},
			"estado_civil":     &graphql.Field{Type: graphql.String},
			"foto":             &graphql.Field{Type: graphql.String},
			"fecha_ingreso":    &graphql.Field{Type: graphql.DateTime},
			"fecha_retiro":     &graphql.Field{Type: graphql.DateTime},
		},
	},
)
