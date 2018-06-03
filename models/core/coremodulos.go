package core

import "time"

type CoreModulo struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// Database custom table name
func (CoreModulo) TableName() string {
	return "core_modulos"
}
