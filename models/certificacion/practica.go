package certificacion

import "github.com/jinzhu/gorm"

type Practica struct {
	gorm.Model
}

func (Practica) TableName() string {
    return "practicas"
}