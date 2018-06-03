package matricula

import "github.com/jinzhu/gorm"

// Pago type definition
type Pago struct {
	gorm.Model
	Concepto string  `json:"concepto"`
	Cantidad float32 `json:"cantidad"`
	Estado   bool    `json:"estado" gorm:"default:'true'"`
}

// Database custom table name
func (Pago) TableName() string {
    return "pagos"
}