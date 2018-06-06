package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models"
)

func CoreSubMenuQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(models.CoreSubMenuType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			db := config.GetConnection()
			defer db.Close()
			submenus := make([]models.CoreSubMenu, 0)
			if err := db.Find(&submenus).Error; err != nil {
				return nil, err
			}
			return submenus, nil
		},
	}
}
