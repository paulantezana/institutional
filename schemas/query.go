package schemas

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
        "Alumnos":    queries.AlumnoQuery(),
        "AlumnoID": queries.AlumnoIDQuery(),
        "Profesores": queries.ProfesorQuery(),
        "Personales": queries.PersonalQuery(),
    }
}
