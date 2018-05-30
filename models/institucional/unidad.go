package institucional

import "github.com/jinzhu/gorm"

// Unidad type definition
type Unidad struct {
	gorm.Model
	Nombre  string `json:"nombre" gorm:"not null"`
	Credito int32  `json:"credito" gorm:"not null"`
	Horas   int32  `json:"horas" gorm:"not null"`
	Estado  bool   `json:"estado" gorm:"default:'true'"`
}
