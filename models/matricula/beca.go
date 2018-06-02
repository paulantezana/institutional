package matricula

import "github.com/jinzhu/gorm"

type Beca struct {
	gorm.Model
	Beca   string `json:"beca"`
	Estado bool   `json:"estado" gorm:"default:'true'"`
}
