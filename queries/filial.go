package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models"
)

func FilialQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(models.FilialType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			db := config.GetConnection()
			defer db.Close()
			filiales := make([]models.Filial, 0)
			if err := db.Find(&filiales).Error; err != nil {
				return nil, err
			}
			return filiales, nil
		},
	}
}

func FilialIDQuery() *graphql.Field {
	return &graphql.Field{
		Type: models.FilialType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			filial := models.Filial{}
			id := p.Args["id"].(int)

			// Get connection
			db := config.GetConnection()
			defer db.Close()

			// Execute instructions
			if err := db.First(&filial, id).Error; err != nil {
				return nil, err
			}
			return filial, nil
		},
	}
}
