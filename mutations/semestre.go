package mutations

import (
    "github.com/graphql-go/graphql"
    "github.com/paulantezana/institutional/models"
    "github.com/paulantezana/institutional/config"
)

func CreateSemestreMutation() *graphql.Field  {
    return &graphql.Field{
        Type: models.SemestreType,
        Args: graphql.FieldConfigArgument{

        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            semestre := models.Semestre{

            }

            // get connection
            db := config.GetConnection()
            defer db.Close()

            // Execute operations
            if err := db.Create(&semestre).Error; err != nil {
                return nil, err
            }

            return semestre, nil
        },
    }
}

func UpdateSemestreMutation() *graphql.Field  {
    return &graphql.Field{
        Type: models.SemestreType,
        Args: graphql.FieldConfigArgument{

        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            semestre := models.Semestre{

            }

            // get connection
            db := config.GetConnection()
            defer db.Close()

            // Execute operations
            if err := db.Create(&semestre).Error; err != nil {
                return nil, err
            }

            return semestre, nil
        },
    }
}

func DeleteSemestreMutation() *graphql.Field  {
    return &graphql.Field{
        Type: models.SemestreType,
        Args: graphql.FieldConfigArgument{

        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            semestre := models.Semestre{

            }

            // get connection
            db := config.GetConnection()
            defer db.Close()

            // Execute operations
            if err := db.Create(&semestre).Error; err != nil {
                return nil, err
            }

            return semestre, nil
        },
    }
}