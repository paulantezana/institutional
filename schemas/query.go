package schemas

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/queries"
)

func RootQuery() graphql.Fields {
	return graphql.Fields{
		"Institutos":  queries.InstitutoQuery(),
		"InstitutoID": queries.InstitutoIDQuery(),
		"Filiales":    queries.FilialQuery(),
		"FilialID":    queries.FilialIDQuery(),
		"Carreras":    queries.CarreraQuery(),
		"CarreraID":   queries.CarreraIDQuery(),
		"Semestres":   queries.SemestreQuery(),
		"SemestreID":  queries.SemestreIDQuery(),
		"Modulos":     queries.ModuloQuery(),
		"ModuloID":    queries.ModuloIDQuery(),
		"Unidades":    queries.UnidadQuery(),
		"UnidadID":    queries.UnidadIDQuery(),

		//"Core":         CoreQuery(),
		"CoreRol":      queries.CoreRolQuery(),
		"CoreMoudulos": queries.CoreModuloQuery(),
		"CoreMenus":    queries.CoreMenuQuery(),
		"CoreSubMenus": queries.CoreSubMenuQuery(),
		"CorePermits":  queries.CorePermitQuery(),

		"Usuarios":   queries.UsuarioQuery(),
		"Alumnos":    queries.AlumnoQuery(),
		"AlumnoID":   queries.AlumnoIDQuery(),
		"Profesores": queries.ProfesorQuery(),
		"Personales": queries.PersonalQuery(),
	}
}
