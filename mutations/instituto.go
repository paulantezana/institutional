package mutations

import (
    "github.com/graphql-go/graphql"
    "github.com/paulantezana/institutional/models"
    "github.com/paulantezana/institutional/config"
)

func CreateInstitutoMutation() *graphql.Field  {
    return &graphql.Field{
        Type: models.InstitutoType,
        Args: graphql.FieldConfigArgument{

        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            instituto := models.Instituto{

            }

            // get connection
            db := config.GetConnection()
            defer db.Close()

            // Execute operations
            if err := db.Create(&instituto).Error; err != nil {
                return nil, err
            }

            return instituto, nil
        },
    }
}

func UpdateInstitutoMutation() *graphql.Field  {
    return &graphql.Field{
        Type: models.InstitutoType,
        Args: graphql.FieldConfigArgument{

        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            instituto := models.Instituto{

            }

            // get connection
            db := config.GetConnection()
            defer db.Close()

            // Execute operations
            if err := db.Create(&instituto).Error; err != nil {
                return nil, err
            }

            return instituto, nil
        },
    }
}

func DeleteInstitutoMutation() *graphql.Field  {
    return &graphql.Field{
        Type: models.InstitutoType,
        Args: graphql.FieldConfigArgument{

        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            instituto := models.Instituto{

            }

            // get connection
            db := config.GetConnection()
            defer db.Close()

            // Execute operations
            if err := db.Create(&instituto).Error; err != nil {
                return nil, err
            }

            return instituto, nil
        },
    }
}