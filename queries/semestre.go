package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models/institucional"
)

func GetSemestreQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(institucional.SemestreType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			db := config.GetConnection()
			defer db.Close()
			semestres := make([]institucional.Semestre, 0)
			if err := db.Find(&semestres).Error; err != nil {
				return nil, err
			}
			return semestres, nil
		},
	}
}
