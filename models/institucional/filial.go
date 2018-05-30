package institucional

import (
	"github.com/jinzhu/gorm"
	"time"
)

// Filial type definition
type Filial struct {
	gorm.Model
	Nombre                       string    `json:"nombre" gorm:"not null"`
	Pais                         string    `json:"pais"`
	Departamento                 string    `json:"departamento"`
	Provincia                    string    `json:"provincia"`
	Distrito                     string    `json:"distrito"`
	CentroPoblado                string    `json:"centro_poblado"`
	DireccionSede                string    `json:"direccion_sede"`
	DireccionMatricula           string    `json:"direccion_matricula"`
	CodigoModular                string    `json:"codigo_modular"`
	TipoGestion                  string    `json:"tipo_gestion"`
	Mension                      string    `json:"mension"`
	ResolucionAutorizacionTipo   string    `json:"resolucion_autorizacion_tipo"`
	ResolucionAutorizacionNumero string    `json:"resolucion_autorizacion_numero"`
	ResolucionAutorizacionFecha  time.Time `json:"resolucion_autorizacion_fecha"`
	ResolucionRevalidacionTipo   string    `json:"resolucion_revalidacion_tipo"`
	ResolucionRevalidacionNumero string    `json:"resolucion_revalidacion_numero"`
	ResolucionRevalidacionFecha  time.Time `json:"resolucion_revalidacion_fecha"`
	Estado                       bool      `json:"estado" gorm:"default:'true'"`
}
