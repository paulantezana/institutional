package institucional

import (
	"time"

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
