package institucional

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
	Nombre                       string     `json:"nombre" gorm:"not null"`
	Pais                         string     `json:"pais"`
	Departamento                 string     `json:"departamento"`
	Provincia                    string     `json:"provincia"`
	Distrito                     string     `json:"distrito"`
	CentroPoblado                string     `json:"centro_poblado"`
	DireccionSede                string     `json:"direccion_sede"`
	DireccionMatricula           string     `json:"direccion_matricula"`
	CodigoModular                string     `json:"codigo_modular"`
	TipoGestion                  string     `json:"tipo_gestion"`
	Mension                      string     `json:"mension"`
	ResolucionAutorizacionTipo   string     `json:"resolucion_autorizacion_tipo"`
	ResolucionAutorizacionNumero string     `json:"resolucion_autorizacion_numero"`
	ResolucionAutorizacionFecha  time.Time  `json:"resolucion_autorizacion_fecha"`
	ResolucionRevalidacionTipo   string     `json:"resolucion_revalidacion_tipo"`
	ResolucionRevalidacionNumero string     `json:"resolucion_revalidacion_numero"`
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
			"id":         &graphql.Field{Type: graphql.Int},
			"created_at": &graphql.Field{Type: graphql.DateTime},
			"updated_at": &graphql.Field{Type: graphql.DateTime},
			"deleted_at": &graphql.Field{Type: graphql.DateTime},
		},
	},
)
