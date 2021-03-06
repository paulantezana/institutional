package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models"
)

func CoreMenuQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(models.CoreMenuType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			db := config.GetConnection()
			defer db.Close()
			menus := make([]models.CoreMenu, 0)
			if err := db.Find(&menus).Error; err != nil {
				return nil, err
			}
			return menus, nil
		},
	}
}
