package mutations

import (
    "github.com/paulantezana/institutional/models"
    "github.com/graphql-go/graphql"
    "github.com/paulantezana/institutional/config"
)

func CreateCoreMutation() *graphql.Field  {
    return &graphql.Field{
        Type: models.CoreMenuType,
        Args: graphql.FieldConfigArgument{

        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            core := models.CoreMenu{

            }

            // get connection
            db := config.GetConnection()
            defer db.Close()

            // Execute operations
            if err := db.Create(&core).Error; err != nil {
                return nil, err
            }

            return core, nil
        },
    }
}

func UpdateCoreMutation() *graphql.Field  {
    return &graphql.Field{
        Type: models.CoreMenuType,
        Args: graphql.FieldConfigArgument{

        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            core := models.CoreMenu{

            }

            // get connection
            db := config.GetConnection()
            defer db.Close()

            // Execute operations
            if err := db.Create(&core).Error; err != nil {
                return nil, err
            }

            return core, nil
        },
    }
}

func DeleteCoreMutation() *graphql.Field  {
    return &graphql.Field{
        Type: models.CoreMenuType,
        Args: graphql.FieldConfigArgument{

        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            core := models.CoreMenu{

            }

            // get connection
            db := config.GetConnection()
            defer db.Close()

            // Execute operations
            if err := db.Create(&core).Error; err != nil {
                return nil, err
            }

            return core, nil
        },
    }
}