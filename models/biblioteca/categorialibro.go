package biblioteca

import "github.com/jinzhu/gorm"

type CategoriaLibro struct {
	gorm.Model
	Nombre string `json:"nombre" gorm:"not null"`
	Estado bool   `json:"estado" gorm:"default:'true'"`

	Libros []Libro `json:"libros,omitempty"`
}

func (CategoriaLibro) TableName() string {
    return "categoria_libros"
}
