package queries

import (
	"crypto/sha256"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models"
	"github.com/paulantezana/institutional/security"
    "time"
    "errors"
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

func ConfirmRecoverPasswordQuery() *graphql.Field {
    return &graphql.Field {
        Type: models.UsuarioType,
        Args: graphql.FieldConfigArgument{
            "clave_recuperar":                  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            usuario := models.Usuario{
                ClaveRecuperar: p.Args["clave_recuperar"].(string),
            }

            // get connection
            db := config.GetConnection()
            defer db.Close()

            // Validations
            if err := db.Where("clave_recuperar = ?", usuario.ClaveRecuperar).First(&usuario).Error; err != nil {
                return nil, err
            }

            // Expiration validation
            if time.Now().Year() != usuario.FechaRecuperacionClave.Year() || time.Now().Month() != usuario.FechaRecuperacionClave.Month() {
                return nil, errors.New("La clave ya expiró")
            }

            if time.Now().Day() - usuario.FechaRecuperacionClave.Day() > 7 {
                return nil, errors.New("La clave ya expiró")
            }

            return usuario, nil
        },
    }
}
