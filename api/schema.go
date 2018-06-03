package api

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/mutations"
	"github.com/paulantezana/institutional/queries"
)

var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{

		//Queries
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootQuery",
			Fields: queries.GetRootFields(),
		}),

		// Mutations
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootMutation",
			Fields: mutations.GetRootFields(),
		}),
	},
)
