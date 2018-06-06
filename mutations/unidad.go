package mutations

import (
    "github.com/paulantezana/institutional/models"
    "github.com/graphql-go/graphql"
    "github.com/paulantezana/institutional/config"
)

func CreateUnidadMutation() *graphql.Field  {
    return &graphql.Field{
        Type: models.UnidadType,
        Args: graphql.FieldConfigArgument{

        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            unidad := models.Unidad{

            }

            // get connection
            db := config.GetConnection()
            defer db.Close()

            // Execute operations
            if err := db.Create(&unidad).Error; err != nil {
                return nil, err
            }

            return unidad, nil
        },
    }
}

func UpdateUnidadMutation() *graphql.Field  {
    return &graphql.Field{
        Type: models.UnidadType,
        Args: graphql.FieldConfigArgument{

        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            unidad := models.Unidad{

            }

            // get connection
            db := config.GetConnection()
            defer db.Close()

            // Execute operations
            if err := db.Create(&unidad).Error; err != nil {
                return nil, err
            }

            return unidad, nil
        },
    }
}

func DeleteUnidadMutation() *graphql.Field  {
    return &graphql.Field{
        Type: models.UnidadType,
        Args: graphql.FieldConfigArgument{

        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            unidad := models.Unidad{

            }

            // get connection
            db := config.GetConnection()
            defer db.Close()

            // Execute operations
            if err := db.Create(&unidad).Error; err != nil {
                return nil, err
            }

            return unidad, nil
        },
    }
}