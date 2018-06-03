package certificacion

import "github.com/jinzhu/gorm"

type PracticaNota struct {
	gorm.Model
}

// Database custom table name
func (PracticaNota) TableName() string {
    return "practica_notas"
}
