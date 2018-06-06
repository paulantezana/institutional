package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models"
)

func CoreQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(models.CoreType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			db := config.GetConnection()
			defer db.Close()
			cores := make([]models.Core, 0)
			if err := db.Find(&cores).Error; err != nil {
				return nil, err
			}
			return cores, nil
		},
	}
}
