package core

import "github.com/jinzhu/gorm"

type CoreModulo struct {
	gorm.Model
}

// Database custom table name
func (CoreModulo) TableName() string {
    return "core_modulos"
}