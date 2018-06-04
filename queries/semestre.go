package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
    "github.com/paulantezana/institutional/models"
)

func GetSemestreQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(models.SemestreType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			db := config.GetConnection()
			defer db.Close()
			semestres := make([]models.Semestre, 0)
			if err := db.Find(&semestres).Error; err != nil {
				return nil, err
			}
			return semestres, nil
		},
	}
}
