package mutations

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models"
    "errors"
    "fmt"
)

func CreateProfesorMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.ProfesorType,
		Args: graphql.FieldConfigArgument{},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			profesor := models.Profesor{}

            // Optional arguments

			// get connection
			db := config.GetConnection()
			defer db.Close()

			// Execute operations
			if err := db.Create(&profesor).Error; err != nil {
				return nil, err
			}

			return profesor, nil
		},
	}
}

func UpdateProfesorMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.ProfesorType,
		Args: graphql.FieldConfigArgument{
            "id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
        },
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			profesor := models.Profesor{}

			// get connection
			db := config.GetConnection()
			defer db.Close()

            if db.First(&profesor).RecordNotFound() {
                return nil, errors.New(fmt.Sprintf("The record with the id %d was not found",profesor.ID))
            }

            // Optional arguments

			// Execute operations
			if err := db.Model(&profesor).Update(profesor).Error; err != nil {
				return nil, err
			}

			return profesor, nil
		},
	}
}

func DeleteProfesorMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.ProfesorType,
		Args: graphql.FieldConfigArgument{
            "id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
        },
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			profesor := models.Profesor{}

			// get connection
			db := config.GetConnection()
			defer db.Close()

            if db.First(&profesor).RecordNotFound() {
                return nil, errors.New(fmt.Sprintf("The record with the id %d was not found",profesor.ID))
            }

			// Execute operations
			if err := db.Delete(&profesor).Error; err != nil {
				return nil, err
			}

			return profesor, nil
		},
	}
}
