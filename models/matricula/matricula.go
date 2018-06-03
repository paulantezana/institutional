package matricula

import (
	"github.com/jinzhu/gorm"
	"time"
)

// Matricula type definition
type Matricula struct {
	gorm.Model
	FechaMatricula time.Time `json:"fecha_matricula"`
	Observacion    string    `json:"observacion"`
	Estado         bool      `json:"estado" gorm:"default:'true'"`
}

// Database custom table name
func (Matricula) TableName() string {
    return "matriculas"
}