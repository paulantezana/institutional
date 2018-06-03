package biblioteca

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Prestamo struct {
	gorm.Model
	LecturaEnSala   bool      `json:"lectura_en_sala"`
	FechaSalida     time.Time `json:"fecha_salida" gorm:"not null"`
	FechaEntrega    time.Time `json:"fecha_entrega"`
	FechaDevolucion time.Time `json:"fecha_devolucion"`
	Observacion     string    `json:"observacion"`
	Estado          bool      `json:"estado" gorm:"default:'true'"`
}

// Database custom table name
func (Prestamo) TableName() string {
    return "prestamos"
}