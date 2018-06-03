package core

import "github.com/jinzhu/gorm"

type CoreSubMenu struct {
	gorm.Model
}

func (CoreSubMenu) TableName() string {
    return "core_sub_menus"
}