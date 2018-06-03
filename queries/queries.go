package queries

import "github.com/graphql-go/graphql"

func GetRootFields() graphql.Fields {
	return graphql.Fields{
		"Instituto": GetInstitutoQuery(),
		"Filial":    GetFilialQuery(),
		"Carrera":   GetCarreraQuery(),
		"Semestre":  GetSemestreQuery(),
		"Modulo":    GetModuloQuery(),
		"Unidad":    GetUnidadQuery(),
		"Usuario":   GetUsuarioQuery(),
	}
}
