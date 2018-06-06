package mutations

import (
    "github.com/graphql-go/graphql"
    "github.com/paulantezana/institutional/models"
    "github.com/paulantezana/institutional/config"
)

func CreateUsuarioMutation() *graphql.Field  {
    return &graphql.Field{
        Type: models.UsuarioType,
        Args: graphql.FieldConfigArgument{

        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            usuario := models.Usuario{

            }

            // get connection
            db := config.GetConnection()
            defer db.Close()

            // Execute operations
            if err := db.Create(&usuario).Error; err != nil {
                return nil, err
            }

            return usuario, nil
        },
    }
}

func UpdateUsuarioMutation() *graphql.Field  {
    return &graphql.Field{
        Type: models.UsuarioType,
        Args: graphql.FieldConfigArgument{

        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            usuario := models.Usuario{

            }

            // get connection
            db := config.GetConnection()
            defer db.Close()

            // Execute operations
            if err := db.Create(&usuario).Error; err != nil {
                return nil, err
            }

            return usuario, nil
        },
    }
}

func DeleteUsuarioMutation() *graphql.Field  {
    return &graphql.Field{
        Type: models.UsuarioType,
        Args: graphql.FieldConfigArgument{

        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            usuario := models.Usuario{

            }

            // get connection
            db := config.GetConnection()
            defer db.Close()

            // Execute operations
            if err := db.Create(&usuario).Error; err != nil {
                return nil, err
            }

            return usuario, nil
        },
    }
}