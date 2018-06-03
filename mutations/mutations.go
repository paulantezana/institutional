package mutations

import "github.com/graphql-go/graphql"

func GetRootFields() graphql.Fields {
	return graphql.Fields{
		//"CreateUser": CreateUser(),
		//"UpdateUser": UpdateUser(),
		"CreateCarrera": CreateCarreraMutation(),
	}
}
