package mutations

import (
    "github.com/graphql-go/graphql"
    "github.com/paulantezana/institutional/models"
    "github.com/paulantezana/institutional/config"
)

func CreateAlumnoMutation() *graphql.Field  {
    return &graphql.Field{
        Type: models.AlumnoType,
        Args: graphql.FieldConfigArgument{

        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            alumno := models.Alumno{

            }

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

func UpdateAlumnoMutation() *graphql.Field  {
    return &graphql.Field{
        Type: models.AlumnoType,
        Args: graphql.FieldConfigArgument{

        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            alumno := models.Alumno{

            }

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

func DeleteAlumnoMutation() *graphql.Field  {
    return &graphql.Field{
        Type: models.AlumnoType,
        Args: graphql.FieldConfigArgument{

        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            alumno := models.Alumno{

            }

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