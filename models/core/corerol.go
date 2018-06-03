package core

import "github.com/jinzhu/gorm"

type CoreRol struct {
	gorm.Model
}

func (CoreRol) TableName() string {
    return "core_roles"
}