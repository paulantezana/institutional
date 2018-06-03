package certificacion

import "github.com/jinzhu/gorm"

type Empresa struct {
	gorm.Model
}

func (Empresa) TableName() string {
    return "empresas"
}