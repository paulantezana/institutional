package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models"
)

func UsuarioQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(models.UsuarioType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			db := config.GetConnection()
			defer db.Close()
			usuarios := make([]models.Usuario, 0)
			if err := db.Find(&usuarios).Error; err != nil {
				return nil, err
			}
			return usuarios, nil
		},
	}
}