package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models/institucional"
)

func GetUsuarioQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(institucional.UsuarioType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			db := config.GetConnection()
			defer db.Close()
			usuarios := make([]institucional.Usuario, 0)
			if err := db.Find(&usuarios).Error; err != nil {
				return nil, err
			}
			return usuarios, nil
		},
	}
}
