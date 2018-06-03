package core

import "github.com/jinzhu/gorm"

type CoreModulo struct {
	gorm.Model
}

func (CoreModulo) TableName() string {
    return "core_modulos"
}