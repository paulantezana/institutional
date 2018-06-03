package institucional

import (
	"github.com/jinzhu/gorm"
)

// Instituto type definition
type Instituto struct {
	gorm.Model
	Estado bool `json:"estado" gorm:"default:'true'"`

	Filiales []Filial `json:"filiales,omitempty"`
}

// Database custom table name
func (Instituto) TableName() string {
    return "institutos"
}
