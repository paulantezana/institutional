package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models"
)

func UnidadQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(models.UnidadType),
        Args: graphql.FieldConfigArgument{
            "search": &graphql.ArgumentConfig{Type: graphql.String},
        },
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            // Declare Variables
            search := ""
            unidades := make([]models.Unidad, 0)

            // Query Params
            if p.Args["search"] != nil {
                search = p.Args["search"].(string)
            }

            // Get connection
            db := config.GetConnection()
            defer db.Close()

			if err := db.Find(&unidades).Error; err != nil {
				return nil, err
			}

            // Execute instructions
            if err := db.Where("nombre LIKE ?", "%"+search+"%").Find(&unidades).Error; err != nil {
                return nil, err
            }

			return unidades, nil
		},
	}
}

func UnidadIDQuery() *graphql.Field {
	return &graphql.Field{
		Type: models.UnidadType,
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
