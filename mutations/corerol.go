package mutations

import (
    "github.com/paulantezana/institutional/models"
    "github.com/graphql-go/graphql"
    "github.com/paulantezana/institutional/config"
)

func CreateCoreRolMutation() *graphql.Field  {
    return &graphql.Field{
        Type: models.CoreRolType,
        Args: graphql.FieldConfigArgument{

        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            rol := models.CoreRol{

            }

            // get connection
            db := config.GetConnection()
            defer db.Close()

            // Execute operations
            if err := db.Create(&rol).Error; err != nil {
                return nil, err
            }

            return rol, nil
        },
    }
}

func UpdateCoreRolMutation() *graphql.Field  {
    return &graphql.Field{
        Type: models.CoreRolType,
        Args: graphql.FieldConfigArgument{

        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            rol := models.CoreRol{

            }

            // get connection
            db := config.GetConnection()
            defer db.Close()

            // Execute operations
            if err := db.Create(&rol).Error; err != nil {
                return nil, err
            }

            return rol, nil
        },
    }
}

func DeleteCoreRolMutation() *graphql.Field  {
    return &graphql.Field{
        Type: models.CoreRolType,
        Args: graphql.FieldConfigArgument{

        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            rol := models.CoreRol{

            }

            // get connection
            db := config.GetConnection()
            defer db.Close()

            // Execute operations
            if err := db.Create(&rol).Error; err != nil {
                return nil, err
            }

            return rol, nil
        },
    }
}