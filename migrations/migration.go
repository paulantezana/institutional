package migrations

import (
	"github.com/paulantezana/institutional/config"
    "github.com/paulantezana/institutional/models"
)

// Migrate function
func Migrate() {
	db := config.GetConnection()
	defer db.Close()

	// Migrate tables
	db.AutoMigrate(
		&models.Core{},
		&models.CoreMenu{},
		&models.CoreModulo{},
		&models.CorePermit{},
		&models.CoreRol{},
		&models.CoreSubMenu{},

		&models.Instituto{},
		&models.Filial{},
		&models.Carrera{},
		&models.Semestre{},
		&models.Modulo{},
		&models.Unidad{},

		&models.Usuario{},
		&models.Alumno{},
		&models.Profesor{},

		&models.Monto{},
		&models.Pago{},
		&models.Beca{},
		&models.Matricula{},

		&models.Empresa{},
		&models.Practica{},
		&models.PracticaNota{},
		&models.Representante{},

		&models.Personal{},

		&models.Nota{},

		&models.CategoriaLibro{},
		&models.Libro{},
		&models.Prestamo{},
	)

	db.Model(&models.Filial{}).AddForeignKey("instituto_id", "institutos(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Carrera{}).AddForeignKey("filial_id", "filiales(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Semestre{}).AddForeignKey("carrera_id", "carreras(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Modulo{}).AddForeignKey("semestre_id", "semestres(id)", "RESTRICT", "RESTRICT")
	db.Model(&models.Unidad{}).AddForeignKey("modulo_id", "modulos(id)", "RESTRICT", "RESTRICT")
}
