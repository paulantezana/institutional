package mutations

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models"
    "errors"
    "fmt"
)

func CreateAlumnoMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.AlumnoType,
		Args: graphql.FieldConfigArgument{},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			alumno := models.Alumno{}

            // Optional arguments

			// get connection
			db := config.GetConnection()
			defer db.Close()

			// Execute operations
			if err := db.Create(&alumno).Error; err != nil {
				return nil, err
			}

			return alumno, nil
		},
	}
}

func UpdateAlumnoMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.AlumnoType,
		Args: graphql.FieldConfigArgument{
            "id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
        },
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			alumno := models.Alumno{}

			// get connection
			db := config.GetConnection()
			defer db.Close()

            if db.First(&alumno).RecordNotFound() {
                return nil, errors.New(fmt.Sprintf("The record with the id %d was not found",alumno.ID))
            }

            // Optional arguments

			// Execute operations
			if err := db.Model(&alumno).Update(alumno).Error; err != nil {
				return nil, err
			}

			return alumno, nil
		},
	}
}

func DeleteAlumnoMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.AlumnoType,
		Args: graphql.FieldConfigArgument{
            "id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
        },
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			alumno := models.Alumno{}

			// get connection
			db := config.GetConnection()
			defer db.Close()

            if db.First(&alumno).RecordNotFound() {
                return nil, errors.New(fmt.Sprintf("The record with the id %d was not found",alumno.ID))
            }

			// Execute operations
			if err := db.Delete(&alumno).Error; err != nil {
				return nil, err
			}

			return alumno, nil
		},
	}
}
