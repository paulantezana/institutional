package certificacion

import "github.com/jinzhu/gorm"

type Empresa struct {
	gorm.Model
}

// Database custom table name
func (Empresa) TableName() string {
    return "empresas"
}