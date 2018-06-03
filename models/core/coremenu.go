package core

import "github.com/jinzhu/gorm"

type CoreMenu struct {
	gorm.Model
}

// Database custom table name
func (CoreMenu) TableName() string {
    return "core_menus"
}