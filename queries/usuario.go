package queries

import (
    "github.com/graphql-go/graphql"
    "github.com/paulantezana/institutional/models/institucional"
    "log"
    "github.com/paulantezana/institutional/config"
)

func GetUsuarioQuery() *graphql.Field  {
    return &graphql.Field{
        Type: graphql.NewList(institucional.UsuarioType),
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            log.Printf("[query] usuarios\n")
            db:= config.GetConnection()
            defer db.Close()
            usuarios := make([]institucional.Usuario, 0)
            if err := db.Find(&usuarios).Error; err != nil {
                return nil, err
            }
            return usuarios, nil
        },
    }
}