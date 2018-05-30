package models

import "github.com/jinzhu/gorm"

// Semestre type definition
type Semestre struct {
	gorm.Model
	Nombre string `json:"nombre" gorm:"not null"`
	Year   int32  `json:"year"`
	Estado bool   `json:"estado" gorm:"default:'true'"`
}
