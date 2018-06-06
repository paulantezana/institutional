package mutations

import (
    "github.com/graphql-go/graphql"
    "github.com/paulantezana/institutional/models"
    "github.com/paulantezana/institutional/config"
)

func CreateFilialMutation() *graphql.Field  {
    return &graphql.Field{
        Type: models.FilialType,
        Args: graphql.FieldConfigArgument{

        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            filial := models.Instituto{

            }

            // get connection
            db := config.GetConnection()
            defer db.Close()

            // Execute operations
            if err := db.Create(&filial).Error; err != nil {
                return nil, err
            }

            return filial, nil
        },
    }
}

func UpdateFilialMutation() *graphql.Field  {
    return &graphql.Field{
        Type: models.FilialType,
        Args: graphql.FieldConfigArgument{

        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            filial := models.Instituto{

            }

            // get connection
            db := config.GetConnection()
            defer db.Close()

            // Execute operations
            if err := db.Create(&filial).Error; err != nil {
                return nil, err
            }

            return filial, nil
        },
    }
}

func DeleteFilialMutation() *graphql.Field  {
    return &graphql.Field{
        Type: models.FilialType,
        Args: graphql.FieldConfigArgument{

        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            filial := models.Instituto{

            }

            // get connection
            db := config.GetConnection()
            defer db.Close()

            // Execute operations
            if err := db.Create(&filial).Error; err != nil {
                return nil, err
            }

            return filial, nil
        },
    }
}