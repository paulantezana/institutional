package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models/institucional"
)

func GetProfesorQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(institucional.ProfesorType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			db := config.GetConnection()
			defer db.Close()
			profesores := make([]institucional.Semestre, 0)
			if err := db.Find(&profesores).Error; err != nil {
				return nil, err
			}
			return profesores, nil
		},
	}
}
