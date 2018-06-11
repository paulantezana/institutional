package queries

import (
	"crypto/sha256"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models"
	"github.com/paulantezana/institutional/security"
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

func LoginQuery() *graphql.Field {
	return &graphql.Field{
		Type: security.TokenType,
		Args: graphql.FieldConfigArgument{
			"usuario": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"clave":   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			user := models.Usuario{
				Usuario: p.Args["nombre"].(string),
				Clave:   p.Args["clave"].(string),
			}

			// Get connection database
			db := config.GetConnection()
			defer db.Close()

			// Execute operations
			cc := sha256.Sum256([]byte(user.Clave))
			pwd := fmt.Sprintf("%x", cc)

			db.Where("usuario=? and clave =?", user.Usuario, pwd).First(&user)
			if user.ID > 0 {
				user.Clave = ""
				token := security.GenerateJWT(user)
				return token, nil
			} else {
				return nil, nil
			}
		},
	}
}
