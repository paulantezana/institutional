package mutations

import "github.com/graphql-go/graphql"

// GetRootFields functions get root mutations graphql
func GetRootFields() graphql.Fields {
	return graphql.Fields{
		// "CreateUser": CreateUser(),
		// "UpdateUser": UpdateUser(),
	}
}
