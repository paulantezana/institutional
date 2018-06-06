package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models"
)

func CorePermitQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(models.CorePermitType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			db := config.GetConnection()
			defer db.Close()
			permits := make([]models.CorePermit, 0)
			if err := db.Find(&permits).Error; err != nil {
				return nil, err
			}
			return permits, nil
		},
	}
}
