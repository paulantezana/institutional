package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models"
)

func InstitutoQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(models.InstitutoType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			db := config.GetConnection()
			defer db.Close()
			institutos := make([]models.Instituto, 0)
			if err := db.Find(&institutos).Error; err != nil {
				return nil, err
			}
			return institutos, nil
		},
	}
}

func InstitutoIDQuery() *graphql.Field {
	return &graphql.Field{
		Type: models.InstitutoType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			instituto := models.Instituto{}
			id := p.Args["id"].(int)

			// Get connection
			db := config.GetConnection()
			defer db.Close()

			// Execute instructions
			if err := db.First(&instituto, id).Error; err != nil {
				return nil, err
			}
			return instituto, nil
		},
	}
}
