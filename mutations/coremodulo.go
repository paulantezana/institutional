package mutations

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models"
)

func CreateCoreModuloMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.CoreModuloType,
		Args: graphql.FieldConfigArgument{},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			modulo := models.CoreModulo{}

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

func UpdateCoreModuloMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.CoreModuloType,
		Args: graphql.FieldConfigArgument{},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			modulo := models.CoreModulo{}

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

func DeleteCoreModuloMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.CoreModuloType,
		Args: graphql.FieldConfigArgument{},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			modulo := models.CoreModulo{}

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
