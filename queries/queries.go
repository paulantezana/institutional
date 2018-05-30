package queries

import "github.com/graphql-go/graphql"

// GetRootFields function get root queries graphql
func GetRootFields() graphql.Fields {
	return graphql.Fields{
		// "User": GetUsers(),
	}
}
