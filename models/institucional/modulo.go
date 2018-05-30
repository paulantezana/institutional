package institucional

import "github.com/jinzhu/gorm"

// Modulo type definition
type Modulo struct {
	gorm.Model
	Nombre      string `json:"nombre" gorm:"not null"`
	Tipo        string `json:"tipo"`
	Descripcion string `json:"descripcion"`
	Estado      bool   `json:"estado" gorm:"default:'true'"`
}
