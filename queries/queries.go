package queries

import "github.com/graphql-go/graphql"

func GetRootFields() graphql.Fields {
	return graphql.Fields{
		"Instituto": InstitutoQuery(),
		"Filial":    FilialQuery(),
		"Carrera":   CarreraQuery(),
		"Semestre":  SemestreQuery(),
		"Modulo":    ModuloQuery(),
		"Unidad":    UnidadQuery(),

		"Core":        CoreQuery(),
		"CoreRol":     CoreRolQuery(),
		"CoreMoudulo": CoreModuloQuery(),
		"CoreMenu":    CoreMenuQuery(),
		"CoreSubMenu": CoreSubMenuQuery(),
		"CorePermit":  CorePermitQuery(),

		"Usuario":  UsuarioQuery(),
		"Alumno":   AlumnoQuery(),
		"Profesor": ProfesorQuery(),
		"Personal": PersonalQuery(),
	}
}
