package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
    "github.com/paulantezana/institutional/models"
)

func GetProfesorQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(models.ProfesorType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			db := config.GetConnection()
			defer db.Close()
			profesores := make([]models.Semestre, 0)
			if err := db.Find(&profesores).Error; err != nil {
				return nil, err
			}
			return profesores, nil
		},
	}
}
