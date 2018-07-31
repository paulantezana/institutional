package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models"
)

func UnidadQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(models.UnidadType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			db := config.GetConnection()
			defer db.Close()
			unidades := make([]models.Unidad, 0)
			if err := db.Find(&unidades).Error; err != nil {
				return nil, err
			}
			return unidades, nil
		},
	}
}

func UnidadIDQuery() *graphql.Field {
	return &graphql.Field{
		Type: models.ModuloType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			unidad := models.Unidad{}
			id := p.Args["id"].(int)

			// Get connection
			db := config.GetConnection()
			defer db.Close()

			// Execute instructions
			if err := db.First(&unidad, id).Error; err != nil {
				return nil, err
			}
			return unidad, nil
		},
	}
}
