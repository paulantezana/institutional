package core

import "github.com/jinzhu/gorm"

type CorePermit struct {
	gorm.Model
}

func (CorePermit) TableName() string {
    return "core_permits"
}