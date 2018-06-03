package core

import "github.com/jinzhu/gorm"

type Core struct {
	gorm.Model
}

func (Core) TableName() string {
    return "core"
}