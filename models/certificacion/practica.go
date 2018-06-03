package certificacion

import "github.com/jinzhu/gorm"

type Practica struct {
	gorm.Model
}

// Database custom table name
func (Practica) TableName() string {
    return "practicas"
}