package core

import "github.com/jinzhu/gorm"

type CoreRol struct {
	gorm.Model
}

// Database custom table name
func (CoreRol) TableName() string {
    return "core_roles"
}