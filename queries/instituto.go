package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models/institucional"
)

func GetInstitutoQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(institucional.InstitutoType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			db := config.GetConnection()
			defer db.Close()
			institutos := make([]institucional.Instituto, 0)
			if err := db.Find(&institutos).Error; err != nil {
				return nil, err
			}
			return institutos, nil
		},
	}
}
