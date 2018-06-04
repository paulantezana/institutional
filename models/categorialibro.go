package models

import (
	"time"
)

type CategoriaLibro struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Nombre    string     `json:"nombre" gorm:"not null"`
	Estado    bool       `json:"estado" gorm:"default:'true'"`

	Libros []Libro `json:"libros,omitempty"`
}

// Database custom table name
func (CategoriaLibro) TableName() string {
	return "categoria_libros"
}
