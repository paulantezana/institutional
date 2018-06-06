package mutations

import (
    "github.com/paulantezana/institutional/models"
    "github.com/graphql-go/graphql"
    "github.com/paulantezana/institutional/config"
)

func CreateCorePermitMutation() *graphql.Field  {
    return &graphql.Field{
        Type: models.CorePermitType,
        Args: graphql.FieldConfigArgument{

        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            permit := models.CorePermit{

            }

            // get connection
            db := config.GetConnection()
            defer db.Close()

            // Execute operations
            if err := db.Create(&permit).Error; err != nil {
                return nil, err
            }

            return permit, nil
        },
    }
}

func UpdateCorePermitMutation() *graphql.Field  {
    return &graphql.Field{
        Type: models.CorePermitType,
        Args: graphql.FieldConfigArgument{

        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            permit := models.CorePermit{

            }

            // get connection
            db := config.GetConnection()
            defer db.Close()

            // Execute operations
            if err := db.Create(&permit).Error; err != nil {
                return nil, err
            }

            return permit, nil
        },
    }
}

func DeleteCorePermitMutation() *graphql.Field  {
    return &graphql.Field{
        Type: models.CorePermitType,
        Args: graphql.FieldConfigArgument{

        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            permit := models.CorePermit{

            }

            // get connection
            db := config.GetConnection()
            defer db.Close()

            // Execute operations
            if err := db.Create(&permit).Error; err != nil {
                return nil, err
            }

            return permit, nil
        },
    }
}