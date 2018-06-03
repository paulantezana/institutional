package certificacion

import "github.com/jinzhu/gorm"

type Representante struct {
	gorm.Model
}

// Database custom table name
func (Representante) TableName() string {
    return "representantes"
}