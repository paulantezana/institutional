package mutations

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models"
)

func CreateModuloMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.ModuloType,
		Args: graphql.FieldConfigArgument{},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			modulo := models.Modulo{}

			// get connection
			db := config.GetConnection()
			defer db.Close()

			// Execute operations
			if err := db.Create(&modulo).Error; err != nil {
				return nil, err
			}

			return modulo, nil
		},
	}
}

func UpdateModuloMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.ModuloType,
		Args: graphql.FieldConfigArgument{},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			modulo := models.Modulo{}

			// get connection
			db := config.GetConnection()
			defer db.Close()

			// Execute operations
			if err := db.Create(&modulo).Error; err != nil {
				return nil, err
			}

			return modulo, nil
		},
	}
}

func DeleteModuloMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.ModuloType,
		Args: graphql.FieldConfigArgument{},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			modulo := models.Modulo{}

			// get connection
			db := config.GetConnection()
			defer db.Close()

			// Execute operations
			if err := db.Create(&modulo).Error; err != nil {
				return nil, err
			}

			return modulo, nil
		},
	}
}
