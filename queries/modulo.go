package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models"
)

func ModuloQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(models.ModuloType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			db := config.GetConnection()
			defer db.Close()
			modulos := make([]models.Modulo, 0)
			if err := db.Find(&modulos).Error; err != nil {
				return nil, err
			}
			return modulos, nil
		},
	}
}

func ModuloIDQuery() *graphql.Field {
	return &graphql.Field{
		Type: models.ModuloType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			modulo := models.Modulo{}
			id := p.Args["id"].(int)

			// Get connection
			db := config.GetConnection()
			defer db.Close()

			// Execute instructions
			if err := db.First(&modulo, id).Error; err != nil {
				return nil, err
			}
			return modulo, nil
		},
	}
}
