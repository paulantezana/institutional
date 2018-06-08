package security

import (
	"github.com/graphql-go/graphql"
)

// Token send json
type Token struct {
	Token string `json:"token"`
}

var TokenType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Token",
		Fields: graphql.Fields{
			"token": &graphql.Field{Type: graphql.String},
		},
	},
)
