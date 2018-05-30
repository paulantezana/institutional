package biblioteca

import "github.com/jinzhu/gorm"

type Libro struct {
	gorm.Model
	Codigo            string `json:"codigo"`
	Tipo              string `json:"tipo"`
	Titulo            string `json:"titulo"`
	Autor             string `json:"autor"`
	Editorial         string `json:"editorial"`
	NCopias           int32  `json:"n_copias"`
	HavilitarPrestamo bool   `json:"havilitar_prestamo"`
	Volumen           int32  `json:"volumen"`
	Paginas           int32  `json:"paginas"`
	NumeroEdicion     int32  `json:"numero_edicion"`
	YearEdicion       int32  `json:"year_edicion"`
	LugarEdicion      string `json:"lugar_edicion"`
	Estado            bool   `json:"estado" gorm:"default:'true'"`
}
