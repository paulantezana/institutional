package biblioteca

import (
	"time"
)

type Prestamo struct {
	ID              uint       `json:"id" gorm:"primary_key"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at"`
	LecturaEnSala   bool       `json:"lectura_en_sala"`
	FechaSalida     time.Time  `json:"fecha_salida" gorm:"not null"`
	FechaEntrega    time.Time  `json:"fecha_entrega"`
	FechaDevolucion time.Time  `json:"fecha_devolucion"`
	Observacion     string     `json:"observacion"`
	Estado          bool       `json:"estado" gorm:"default:'true'"`
}

// Database custom table name
func (Prestamo) TableName() string {
	return "prestamos"
}
