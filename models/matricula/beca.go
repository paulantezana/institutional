package matricula

import "github.com/jinzhu/gorm"

// Beca type definition
type Beca struct {
	gorm.Model
	Beca   string `json:"beca"`
	Estado bool   `json:"estado" gorm:"default:'true'"`
}

// Database custom table name
func (Beca) TableName() string {
    return "becas"
}