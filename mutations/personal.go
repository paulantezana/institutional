package mutations

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models"
    "errors"
    "fmt"
)

func CreatePersonalMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.PersonalType,
		Args: graphql.FieldConfigArgument{},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			personal := models.Personal{}

            // Optional arguments

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
		Args: graphql.FieldConfigArgument{
            "id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
        },
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			personal := models.Personal{}

			// get connection
			db := config.GetConnection()
			defer db.Close()

            if db.First(&personal).RecordNotFound() {
                return nil, errors.New(fmt.Sprintf("The record with the id %d was not found",personal.ID))
            }

            // Optional arguments

			// Execute operations
			if err := db.Model(&personal).Update(personal).Error; err != nil {
				return nil, err
			}

			return personal, nil
		},
	}
}

func DeletePersonalMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.PersonalType,
		Args: graphql.FieldConfigArgument{
            "id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
        },
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			personal := models.Personal{}

			// get connection
			db := config.GetConnection()
			defer db.Close()

            if db.First(&personal).RecordNotFound() {
                return nil, errors.New(fmt.Sprintf("The record with the id %d was not found",personal.ID))
            }

			// Execute operations
			if err := db.Delete(&personal).Error; err != nil {
				return nil, err
			}

			return personal, nil
		},
	}
}
