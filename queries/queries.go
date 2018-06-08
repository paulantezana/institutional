package queries

import "github.com/graphql-go/graphql"

func GetRootFields() graphql.Fields {
	return graphql.Fields{
		"Institutos": InstitutoQuery(),
		"Filiales":   FilialQuery(),
		"Carreras":   CarreraQuery(),
		"Semestres":  SemestreQuery(),
		"Modulos":    ModuloQuery(),
		"Unidades":   UnidadQuery(),

		"Core":         CoreQuery(),
		"CoreRol":      CoreRolQuery(),
		"CoreMoudulos": CoreModuloQuery(),
		"CoreMenus":    CoreMenuQuery(),
		"CoreSubMenus": CoreSubMenuQuery(),
		"CorePermits":  CorePermitQuery(),

		"Login": LoginQuery(),
		"Usuarios":   UsuarioQuery(),
		"Alumnos":    AlumnoQuery(),
		"Profesores": ProfesorQuery(),
		"Personales": PersonalQuery(),
	}
}
