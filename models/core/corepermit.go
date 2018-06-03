package core

import "github.com/jinzhu/gorm"

type CorePermit struct {
	gorm.Model
}

// Database custom table name
func (CorePermit) TableName() string {
    return "core_permits"
}