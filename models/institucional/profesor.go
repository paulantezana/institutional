package institucional

import (
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
	"time"
)

type Profesor struct {
	gorm.Model
	Dni              string    `json:"dni" gorm:"not null; unique"`
	Nombres          string    `json:"nombres" gorm:"not null"`
	ApellidoPaterno  string    `json:"apellido_paterno" gorm:"not null"`
	ApellidoMaterno  string    `json:"apellido_materno" gorm:"not null"`
	FechaNacimiento  time.Time `json:"fecha_nacimiento"`
	Sexo             string    `json:"sexo"`
	Pais             string    `json:"pais" gorm:"not null"`
	Departamento     string    `json:"departamento"`
	Direccion        string    `json:"direccion"`
	Celular          string    `json:"celular" gorm:"unique"`
	Telefono         string    `json:"telefono" gorm:"unique"`
	Correo           string    `json:"correo" gorm:"unique"`
	EstadoCivil      string    `json:"estado_civil"`
	Foto             string    `json:"foto"`
	CondicionLaboral string    `json:"condicion_laboral"`
	NivelEducativo   string    `json:"nivel_educativo"`
	FechaIngreso     time.Time `json:"fecha_ingreso" gorm:"not null"`
	FechaRetiro      time.Time `json:"fecha_retiro"`
	AnosDocencia     int32     `json:"anos_docencia"`
	MesesDocencia    int32     `json:"meses_docencia"`
	Especialidad     string    `json:"especialidad"`
	NumeroHoras      int32     `json:"numero_horas"`
	Estado           bool      `json:"estado" gorm:"default:'true'"`
}

// Model GraphQL
var ProfesorType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:   "Profesor",
		Fields: graphql.Fields{},
	},
)
