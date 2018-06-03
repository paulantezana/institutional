package matricula

import "github.com/jinzhu/gorm"

// Monto type definition
type Monto struct {
	gorm.Model
	Concepto string  `json:"concepto"`
	Cantidad float32 `json:"cantidad"`
	Estado   bool    `json:"estado" gorm:"default:'true'"`
}

// Database custom table name
func (Monto) TableName() string {
    return "montos"
}