package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models"
)

func CoreRolQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(models.CoreRolType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			db := config.GetConnection()
			defer db.Close()
			roles := make([]models.CoreRol, 0)
			if err := db.Find(&roles).Error; err != nil {
				return nil, err
			}
			return roles, nil
		},
	}
}
