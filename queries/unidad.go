package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models"
)

func UnidadQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(models.UnidadType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			db := config.GetConnection()
			defer db.Close()
			unidades := make([]models.Unidad, 0)
			if err := db.Find(&unidades).Error; err != nil {
				return nil, err
			}
			return unidades, nil
		},
	}
}
