package mutations

import (
    "github.com/graphql-go/graphql"
    "github.com/paulantezana/institutional/models"
    "github.com/paulantezana/institutional/config"
)

func CreateCoreSubMenuMutation() *graphql.Field  {
    return &graphql.Field{
        Type: models.CoreSubMenuType,
        Args: graphql.FieldConfigArgument{

        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            submenu := models.CoreSubMenu{

            }

            // get connection
            db := config.GetConnection()
            defer db.Close()

            // Execute operations
            if err := db.Create(&submenu).Error; err != nil {
                return nil, err
            }

            return submenu, nil
        },
    }
}

func UpdateCoreSubMenuMutation() *graphql.Field  {
    return &graphql.Field{
        Type: models.CoreSubMenuType,
        Args: graphql.FieldConfigArgument{

        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            submenu := models.CoreSubMenu{

            }

            // get connection
            db := config.GetConnection()
            defer db.Close()

            // Execute operations
            if err := db.Create(&submenu).Error; err != nil {
                return nil, err
            }

            return submenu, nil
        },
    }
}

func DeleteCoreSubMenuMutation() *graphql.Field  {
    return &graphql.Field{
        Type: models.CoreSubMenuType,
        Args: graphql.FieldConfigArgument{

        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            submenu := models.CoreSubMenu{

            }

            // get connection
            db := config.GetConnection()
            defer db.Close()

            // Execute operations
            if err := db.Create(&submenu).Error; err != nil {
                return nil, err
            }

            return submenu, nil
        },
    }
}