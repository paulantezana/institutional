package certificacion

import "time"

type PracticaNota struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// Database custom table name
func (PracticaNota) TableName() string {
	return "practica_notas"
}
