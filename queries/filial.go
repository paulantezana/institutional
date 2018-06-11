package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models"
)

func FilialQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(models.FilialType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			db := config.GetConnection()
			defer db.Close()
			filiales := make([]models.Filial, 0)
			if err := db.Find(&filiales).Error; err != nil {
				return nil, err
			}
			return filiales, nil
		},
	}
}
