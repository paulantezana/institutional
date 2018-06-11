package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models"
)

func PersonalQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(models.PersonalType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			db := config.GetConnection()
			defer db.Close()
			personales := make([]models.Personal, 0)
			if err := db.Find(&personales).Error; err != nil {
				return nil, err
			}
			return personales, nil
		},
	}
}
