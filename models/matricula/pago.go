package matricula

import "github.com/jinzhu/gorm"

type Pago struct {
	gorm.Model
	Concepto string  `json:"concepto"`
	Cantidad float32 `json:"cantidad"`
	Estado   bool    `json:"estado" gorm:"default:'true'"`
}
