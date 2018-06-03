package core

import "github.com/jinzhu/gorm"

type CoreMenu struct {
	gorm.Model
}

func (CoreMenu) TableName() string {
    return "core_menus"
}