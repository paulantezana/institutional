package schema

import (
    "github.com/graphql-go/graphql"
    "github.com/paulantezana/institutional/queries"
)

func RootQuery() graphql.Fields {
    return graphql.Fields{
        "Institutos": queries.InstitutoQuery(),
        "Filiales":   queries.FilialQuery(),
        "Carreras":   queries.CarreraQuery(),
        "Semestres":  queries.SemestreQuery(),
        "Modulos":    queries.ModuloQuery(),
        "Unidades":   queries.UnidadQuery(),

        //"Core":         CoreQuery(),
        "CoreRol":      queries.CoreRolQuery(),
        "CoreMoudulos": queries.CoreModuloQuery(),
        "CoreMenus":    queries.CoreMenuQuery(),
        "CoreSubMenus": queries.CoreSubMenuQuery(),
        "CorePermits":  queries.CorePermitQuery(),

        "Usuarios":   queries.UsuarioQuery(),
        "ConfirmRecoverPassword":queries.ConfirmRecoverPasswordQuery(),
        "Alumnos":    queries.AlumnoQuery(),
        "Profesores": queries.ProfesorQuery(),
        "Personales": queries.PersonalQuery(),
    }
}
