package models

import "time"

// Pago type definition
type Pago struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Concepto  string     `json:"concepto"`
	Cantidad  float32    `json:"cantidad"`
	Estado    bool       `json:"estado" gorm:"default:'true'"`
}

// Database custom table name
func (Pago) TableName() string {
	return "pagos"
}
