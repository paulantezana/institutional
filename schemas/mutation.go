package schemas

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/mutations"
)

func RootMutation() graphql.Fields {
	return graphql.Fields{
		"CreateInstituto": mutations.CreateInstitutoMutation(),
		"UpdateInstituto": mutations.UpdateInstitutoMutation(),
		"DeleteInstituto": mutations.DeleteInstitutoMutation(),

		"CreateFilial": mutations.CreateFilialMutation(),
		"UpdateFilial": mutations.UpdateFilialMutation(),
		"DeleteFilial": mutations.DeleteFilialMutation(),

		"CreateCarrera": mutations.CreateCarreraMutation(),
		"UpdateCarrera": mutations.UpdateCarreraMutation(),
		"DeleteCarrera": mutations.DeleteCarreraMutation(),

		"CreateSemestre": mutations.CreateSemestreMutation(),
		"UpdateSemestre": mutations.UpdateSemestreMutation(),
		"DeleteSemestre": mutations.DeleteSemestreMutation(),

		"CreateModulo": mutations.CreateModuloMutation(),
		"UpdateModulo": mutations.UpdateModuloMutation(),
		"DeleteModulo": mutations.DeleteModuloMutation(),

		"CreateUnidad": mutations.CreateUnidadMutation(),
		"UpdateUnidad": mutations.UpdateUnidadMutation(),
		"DeleteUnidad": mutations.DeleteUnidadMutation(),

		// roles and user permissions
		"CreateCoreRol": mutations.CreateCoreRolMutation(),
		"UpdateCoreRol": mutations.UpdateCoreRolMutation(),
		"DeleteCoreRol": mutations.DeleteCoreRolMutation(),

		"CreateUsuario": mutations.CreateUsuarioMutation(),
		"UpdateUsuario": mutations.UpdateUsuarioMutation(),
		"DeleteUsuario": mutations.DeleteUsuarioMutation(),

		"CreateProfesor": mutations.CreateProfesorMutation(),
		"UpdateProfesor": mutations.UpdateProfesorMutation(),
		"DeleteProfesor": mutations.DeleteProfesorMutation(),

		"CreateAlumno": mutations.CreateAlumnoMutation(),
		"UpdateAlumno": mutations.UpdateAlumnoMutation(),
		"DeleteAlumno": mutations.DeleteAlumnoMutation(),

		"CreatePersonal": mutations.CreatePersonalMutation(),
		"UpdatePersonal": mutations.UpdatePersonalMutation(),
		"DeletePersonal": mutations.DeletePersonalMutation(),
	}
}
