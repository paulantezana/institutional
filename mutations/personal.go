package mutations

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models"
)

func CreatePersonalMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.PersonalType,
		Args: graphql.FieldConfigArgument{},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			personal := models.Personal{}

			// get connection
			db := config.GetConnection()
			defer db.Close()

			// Execute operations
			if err := db.Create(&personal).Error; err != nil {
				return nil, err
			}

			return personal, nil
		},
	}
}

func UpdatePersonalMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.PersonalType,
		Args: graphql.FieldConfigArgument{},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			personal := models.Personal{}

			// get connection
			db := config.GetConnection()
			defer db.Close()

			// Execute operations
			if err := db.Create(&personal).Error; err != nil {
				return nil, err
			}

			return personal, nil
		},
	}
}

func DeletePersonalMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.PersonalType,
		Args: graphql.FieldConfigArgument{},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			personal := models.Personal{}

			// get connection
			db := config.GetConnection()
			defer db.Close()

			// Execute operations
			if err := db.Create(&personal).Error; err != nil {
				return nil, err
			}

			return personal, nil
		},
	}
}
