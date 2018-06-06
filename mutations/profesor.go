package mutations

import (
    "github.com/paulantezana/institutional/models"
    "github.com/graphql-go/graphql"
    "github.com/paulantezana/institutional/config"
)

func CreateProfesorMutation() *graphql.Field  {
    return &graphql.Field{
        Type: models.ProfesorType,
        Args: graphql.FieldConfigArgument{

        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            profesor := models.Profesor{

            }

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

func UpdateProfesorMutation() *graphql.Field  {
    return &graphql.Field{
        Type: models.ProfesorType,
        Args: graphql.FieldConfigArgument{

        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            profesor := models.Profesor{

            }

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

func DeleteProfesorMutation() *graphql.Field  {
    return &graphql.Field{
        Type: models.ProfesorType,
        Args: graphql.FieldConfigArgument{

        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            profesor := models.Profesor{

            }

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