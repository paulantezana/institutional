package core

import "github.com/jinzhu/gorm"

type Core struct {
	gorm.Model
}

// Database custom table name
func (Core) TableName() string {
    return "core"
}