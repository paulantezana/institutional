package matricula

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Matricula struct {
	gorm.Model
	FechaMatricula time.Time `json:"fecha_matricula"`
	Observacion    string    `json:"observacion"`
	Estado         bool      `json:"estado" gorm:"default:'true'"`
}
