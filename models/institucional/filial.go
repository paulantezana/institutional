package institucional

import "github.com/jinzhu/gorm"

// Filial type definition
type Filial struct {
	gorm.Model
	Nombre string `json:"nombre" gorm:"not null"`
	Estado bool   `json:"estado" gorm:"default:'true'"`
}
