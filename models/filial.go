package models

import (
	"github.com/graphql-go/graphql"
	"time"
)

// Filial type definition
type Filial struct {
	ID                           uint       `json:"id" gorm:"primary_key"`
	CreatedAt                    time.Time  `json:"created_at"`
	UpdatedAt                    time.Time  `json:"updated_at"`
	DeletedAt                    *time.Time `json:"deleted_at"`
	Nombre                       string     `json:"nombre" gorm:"type:varchar(128); not null"`
	Pais                         string     `json:"pais" gorm:"type:varchar(64)"`
	Departamento                 string     `json:"departamento" gorm:"type:varchar(64)"`
	Provincia                    string     `json:"provincia" gorm:"type:varchar(64)"`
	Distrito                     string     `json:"distrito" gorm:"type:varchar(128)"`
	CentroPoblado                string     `json:"centro_poblado" gorm:"type:varchar(128)"`
	DireccionSede                string     `json:"direccion_sede" gorm:"type:varchar(128)"`
	DireccionMatricula           string     `json:"direccion_matricula" gorm:"type:varchar(128)"`
	CodigoModular                string     `json:"codigo_modular" gorm:"type:varchar(64)"`
	TipoGestion                  string     `json:"tipo_gestion" gorm:"type:varchar(128)"`
	Mension                      string     `json:"mension" gorm:"type:varchar(128)"`
	ResolucionAutorizacionTipo   string     `json:"resolucion_autorizacion_tipo" gorm:"type:varchar(64)"`
	ResolucionAutorizacionNumero string     `json:"resolucion_autorizacion_numero" gorm:"type:varchar(64)"`
	ResolucionAutorizacionFecha  time.Time  `json:"resolucion_autorizacion_fecha"`
	ResolucionRevalidacionTipo   string     `json:"resolucion_revalidacion_tipo" gorm:"type:varchar(64)"`
	ResolucionRevalidacionNumero string     `json:"resolucion_revalidacion_numero" gorm:"type:varchar(64)"`
	ResolucionRevalidacionFecha  time.Time  `json:"resolucion_revalidacion_fecha"`
	Estado                       bool       `json:"estado" gorm:"default:'true'"`
	InstitutoID                  uint       `json:"instituto_id"`
	Carreras                     []Carrera  `json:"carreras,omitempty" `
}

// Database custom table name
func (Filial) TableName() string {
	return "filiales"
}

// Model GraphQL
var FilialType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Filial",
		Fields: graphql.Fields{
			"id":                             &graphql.Field{Type: graphql.Int},
			"created_at":                     &graphql.Field{Type: graphql.DateTime},
			"updated_at":                     &graphql.Field{Type: graphql.DateTime},
			"deleted_at":                     &graphql.Field{Type: graphql.DateTime},
			"nombre":                         &graphql.Field{Type: graphql.String},
			"pais":                           &graphql.Field{Type: graphql.String},
			"departamento":                   &graphql.Field{Type: graphql.String},
			"provincia":                      &graphql.Field{Type: graphql.String},
			"distrito":                       &graphql.Field{Type: graphql.String},
			"centro_poblado":                 &graphql.Field{Type: graphql.String},
			"direccion_sede":                 &graphql.Field{Type: graphql.String},
			"direccion_matricula":            &graphql.Field{Type: graphql.String},
			"codigo_modular":                 &graphql.Field{Type: graphql.String},
			"tipo_gestion":                   &graphql.Field{Type: graphql.String},
			"mension":                        &graphql.Field{Type: graphql.String},
			"resolucion_autorizacion_tipo":   &graphql.Field{Type: graphql.String},
			"resolucion_autorizacion_numero": &graphql.Field{Type: graphql.String},
			"resolucion_autorizacion_fecha":  &graphql.Field{Type: graphql.DateTime},
			"resolucion_revalidacion_tipo":   &graphql.Field{Type: graphql.String},
			"resolucion_revalidacion_numero": &graphql.Field{Type: graphql.String},
			"resolucion_revalidacion_fecha":  &graphql.Field{Type: graphql.DateTime},
			"instituto_id":                   &graphql.Field{Type: graphql.Int},
			"estado":                         &graphql.Field{Type: graphql.Boolean},
			"carreras":                       &graphql.Field{Type: graphql.NewList(CarreraType)},
		},
	},
)
