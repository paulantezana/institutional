package mutations

import "github.com/graphql-go/graphql"

func GetRootFields() graphql.Fields {
	return graphql.Fields{
		"CreateInstituto": CreateInstitutoMutation(),
		"UpdateInstituto": UpdateInstitutoMutation(),
		"DeleteInstituto": DeleteInstitutoMutation(),

		"CreateFilial": CreateFilialMutation(),
		"UpdateFilial": UpdateFilialMutation(),
		"DeleteFilial": DeleteFilialMutation(),

		"CreateCarrera": CreateCarreraMutation(),
		"UpdateCarrera": UpdateCarreraMutation(),
		"DeleteCarrera": DeleteCarreraMutation(),

		"CreateSemestre": CreateSemestreMutation(),
		"UpdateSemestre": UpdateSemestreMutation(),
		"DeleteSemestre": DeleteSemestreMutation(),

		"CreateModulo": CreateModuloMutation(),
		"UpdateModulo": UpdateModuloMutation(),
		"DeleteModulo": DeleteModuloMutation(),

		"CreateUnidad": CreateUnidadMutation(),
		"UpdateUnidad": UpdateUnidadMutation(),
		"DeleteUnidad": DeleteUnidadMutation(),
	}
}
