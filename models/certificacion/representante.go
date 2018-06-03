package certificacion

import "github.com/jinzhu/gorm"

type Representante struct {
	gorm.Model
}

func (Representante) TableName() string {
    return "representantes"
}