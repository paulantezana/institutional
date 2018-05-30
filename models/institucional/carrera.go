package institucional

import "github.com/jinzhu/gorm"

// Carrera type definition
type Carrera struct {
	gorm.Model
	Nombre      string `json:"nombre" gorm:"not null; unique"`
	Logo        string `json:"logo"`
	Descripcion string `json:"descripcion"`
	Creacion    int32  `json:"creacion"`
	Estado      bool   `json:"estado" gorm:"default:'true'"`
}
