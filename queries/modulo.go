package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
    "github.com/paulantezana/institutional/models"
)

func GetModuloQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(models.ModuloType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			db := config.GetConnection()
			defer db.Close()
			modulos := make([]models.Modulo, 0)
			if err := db.Find(&modulos).Error; err != nil {
				return nil, err
			}
			return modulos, nil
		},
	}
}
