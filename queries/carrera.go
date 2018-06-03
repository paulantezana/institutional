package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models/institucional"
)

func GetCarreraQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(institucional.CarreraType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			db := config.GetConnection()
			defer db.Close()
			carreras := make([]institucional.Carrera, 0)
			if err := db.Find(&carreras).Error; err != nil {
				return nil, err
			}
			return carreras, nil
		},
	}
}
