package matricula

import (
	"time"
)

// Matricula type definition
type Matricula struct {
	ID             uint       `json:"id" gorm:"primary_key"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at"`
	FechaMatricula time.Time  `json:"fecha_matricula"`
	Observacion    string     `json:"observacion"`
	Estado         bool       `json:"estado" gorm:"default:'true'"`
}

// Database custom table name
func (Matricula) TableName() string {
	return "matriculas"
}
