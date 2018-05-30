package migrations

import (
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models/institucional"
)

// Migrate function
func Migrate() {
	db := config.GetConnection()
	defer db.Close()

	// Create tables in database INSTITUCIONAL
	db.CreateTable(&institucional.Instituto{})
	db.CreateTable(&institucional.Filial{})
	db.CreateTable(&institucional.Carrera{})
	db.CreateTable(&institucional.Semestre{})
	db.CreateTable(&institucional.Modulo{})
	db.CreateTable(&institucional.Unidad{})

	// ForeignKey Relations table
	db.Model(&institucional.Instituto{}).Related(&institucional.Filial{})

	db.CreateTable(&institucional.Usuario{})
}
