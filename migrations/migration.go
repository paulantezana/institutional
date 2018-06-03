package migrations

import (
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models/administracion"
	"github.com/paulantezana/institutional/models/biblioteca"
	"github.com/paulantezana/institutional/models/certificacion"
	"github.com/paulantezana/institutional/models/core"
	"github.com/paulantezana/institutional/models/institucional"
	"github.com/paulantezana/institutional/models/matricula"
	"github.com/paulantezana/institutional/models/nota"
)

// Migrate function
func Migrate() {
	db := config.GetConnection()
	defer db.Close()

	// Migrate tables
	db.AutoMigrate(
		&core.Core{},
		&core.CoreMenu{},
		&core.CoreModulo{},
		&core.CorePermit{},
		&core.CoreRol{},
		&core.CoreSubMenu{},

		&institucional.Instituto{},
		&institucional.Filial{},
		&institucional.Carrera{},
		&institucional.Semestre{},
		&institucional.Modulo{},
		&institucional.Unidad{},

		&institucional.Usuario{},
		&institucional.Alumno{},
		&institucional.Profesor{},

		&matricula.Monto{},
		&matricula.Pago{},
		&matricula.Beca{},
		&matricula.Matricula{},

		&certificacion.Empresa{},
		&certificacion.Practica{},
		&certificacion.PracticaNota{},
		&certificacion.Representante{},

		&administracion.Personal{},

		&nota.Nota{},

		&biblioteca.CategoriaLibro{},
		&biblioteca.Libro{},
		&biblioteca.Prestamo{},
	)

	db.Model(&institucional.Filial{}).AddForeignKey("instituto_id", "institutos(id)", "RESTRICT", "RESTRICT")
    db.Model(&institucional.Carrera{}).AddForeignKey("filial_id", "filiales(id)", "RESTRICT", "RESTRICT")
    db.Model(&institucional.Semestre{}).AddForeignKey("carrera_id", "carreras(id)", "RESTRICT", "RESTRICT")
    db.Model(&institucional.Modulo{}).AddForeignKey("semestre_id", "semestres(id)", "RESTRICT", "RESTRICT")
    db.Model(&institucional.Unidad{}).AddForeignKey("modulo_id", "modulos(id)", "RESTRICT", "RESTRICT")
}
