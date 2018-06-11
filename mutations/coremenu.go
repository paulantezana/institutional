package mutations

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models"
)

func CreateCoreMenuMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.CoreMenuType,
		Args: graphql.FieldConfigArgument{},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			menu := models.CoreMenu{}

			// get connection
			db := config.GetConnection()
			defer db.Close()

			// Execute operations
			if err := db.Create(&menu).Error; err != nil {
				return nil, err
			}

			return menu, nil
		},
	}
}

func UpdateCoreMenuMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.CoreMenuType,
		Args: graphql.FieldConfigArgument{},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			menu := models.CoreMenu{}

			// get connection
			db := config.GetConnection()
			defer db.Close()

			// Execute operations
			if err := db.Create(&menu).Error; err != nil {
				return nil, err
			}

			return menu, nil
		},
	}
}

func DeleteCoreMenuMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.CoreMenuType,
		Args: graphql.FieldConfigArgument{},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			menu := models.CoreMenu{}

			// get connection
			db := config.GetConnection()
			defer db.Close()

			// Execute operations
			if err := db.Create(&menu).Error; err != nil {
				return nil, err
			}

			return menu, nil
		},
	}
}
