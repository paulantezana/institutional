package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models"
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
