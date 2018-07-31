package schemas

import "github.com/graphql-go/graphql"

var Schema graphql.Schema

func init() {
	// Root mutation
	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name:   "RootMutation",
		Fields: RootMutation(),
	})

	// Root query
	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name:   "RootQuery",
		Fields: RootQuery(),
	})

	// Define schema
	var err error
	Schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})

	if err != nil {
		panic(err)
	}
}
