package certificacion

import "github.com/jinzhu/gorm"

type PracticaNota struct {
	gorm.Model
}

func (PracticaNota) TableName() string {
    return "practica_notas"
}
