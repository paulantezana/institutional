package administracion

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Personal struct {
	gorm.Model
	Dni             string    `json:"dni" gorm:"not null; unique"`
	Nombres         string    `json:"nombres" gorm:"not null"`
	ApellidoPaterno string    `json:"apellido_paterno" gorm:"not null"`
	ApellidoMaterno string    `json:"apellido_materno" gorm:"not null"`
	FechaNacimiento time.Time `json:"fecha_nacimiento"`
	Sexo            string    `json:"sexo"`
	Pais            string    `json:"pais" gorm:"not null"`
	Departamento    string    `json:"departamento"`
	Direccion       string    `json:"direccion"`
	Celular         string    `json:"celular" gorm:"unique"`
	Telefono        string    `json:"telefono" gorm:"unique"`
	Correo          string    `json:"correo" gorm:"unique"`
	EstadoCivil     string    `json:"estado_civil"`
	Foto            string    `json:"foto"`
	FechaIngreso    time.Time `json:"fecha_ingreso"`
	FechaRetiro     time.Time `json:"fecha_retiro"`
	Estado          bool      `json:"estado" gorm:"default:'true'"`
}

func (Personal) TableName() string {
    return "personales"
}