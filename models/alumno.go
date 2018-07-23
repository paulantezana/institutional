package models

import (
	"github.com/graphql-go/graphql"
	"time"
)

// Alumno type definition
type Alumno struct {
	ID              uint       `json:"id" gorm:"primary_key"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at"`
	Dni             string     `json:"dni" gorm:"type:varchar(15); not null; unique"`
	Nombres         string     `json:"nombres" gorm:"type:varchar(128); not null"`
	ApellidoPaterno string     `json:"apellido_paterno" gorm:"type:varchar(128); not null"`
	ApellidoMaterno string     `json:"apellido_materno" gorm:"type:varchar(128); not null"`
	FechaNacimiento time.Time  `json:"fecha_nacimiento"`
	Sexo            string     `json:"sexo"`
	Pais            string     `json:"pais" gorm:"type:varchar(64)"`
	Departamento    string     `json:"departamento" gorm:"type:varchar(64);"`
	Direccion       string     `json:"direccion" gorm:"type:varchar(128);"`
	Celular         string     `json:"celular" gorm:"type:varchar(16); unique"`
	Telefono        string     `json:"telefono" gorm:"type:varchar(16); unique"`
	EstadoCivil     string     `json:"estado_civil" gorm:"type:varchar(64)"`
	FechaAdmicion   time.Time  `json:"fecha_admicion"`
	FechaPromocion  time.Time  `json:"fecha_promocion"`
	Estado          bool       `json:"estado" gorm:"default:'true'"`

	UsuarioID       uint       `json:"usuario_id"`
}

// Database custom table name
func (Alumno) TableName() string {
	return "alumnos"
}

// Model GraphQL
var AlumnoType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Alumno",
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
			"estado_civil":     &graphql.Field{Type: graphql.String},
			"fecha_admicion":   &graphql.Field{Type: graphql.DateTime},
			"fecha_promocion":  &graphql.Field{Type: graphql.DateTime},
			"estado":           &graphql.Field{Type: graphql.Boolean},
		},
	},
)
