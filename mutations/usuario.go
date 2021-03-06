package mutations

import (
	"crypto/md5"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models"
	"time"
)

func CreateUsuarioMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.UsuarioType,
		Args: graphql.FieldConfigArgument{
			"usuario":         &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"correo":          &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"clave":           &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"confirmar_clave": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"estado":          &graphql.ArgumentConfig{Type: graphql.Boolean},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			usuario := models.Usuario{
				Usuario:                p.Args["usuario"].(string),
				Correo:                 p.Args["correo"].(string),
				Clave:                  p.Args["clave"].(string),
				ConfirmarClave:         p.Args["confirmar_clave"].(string),
				FechaModificacionClave: time.Now(),
			}

			// Optional arguments
			if p.Args["estado"] != nil {
				usuario.Estado = p.Args["estado"].(bool)
			}

			// get connection
			db := config.GetConnection()
			defer db.Close()

			// Execute operations
			if usuario.Clave != usuario.ConfirmarClave {
				return nil, errors.New("Las claves no coinciden")
			}

			cc := sha256.Sum256([]byte(usuario.Clave))
			pwd := fmt.Sprintf("%x", cc)
			usuario.Clave = pwd

			picmd5 := md5.Sum([]byte(usuario.Correo))
			picstr := fmt.Sprintf("%x", picmd5)
			usuario.Avatar = "https://gravatar.com/avatar/" + picstr + "?s=100"

			if err := db.Create(&usuario).Error; err != nil {
				return nil, err
			}

			return usuario, nil
		},
	}
}

// UpdateUsuarioMutation : change username, password, etc.
func UpdateUsuarioMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.UsuarioType,
		Args: graphql.FieldConfigArgument{
			"id":              &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			"usuario":         &graphql.ArgumentConfig{Type: graphql.String},
			"correo":          &graphql.ArgumentConfig{Type: graphql.String},
			"clave":           &graphql.ArgumentConfig{Type: graphql.String},
			"confirmar_clave": &graphql.ArgumentConfig{Type: graphql.String},
			"clave_antigua":   &graphql.ArgumentConfig{Type: graphql.String},
			"estado":          &graphql.ArgumentConfig{Type: graphql.Boolean},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			usuario := models.Usuario{
				ID: uint(p.Args["id"].(int)),
			}

			// get connection
			db := config.GetConnection()
			defer db.Close()

			if db.First(&usuario).RecordNotFound() {
				return nil, errors.New(fmt.Sprintf("The record with the id %d was not found", usuario.ID))
			}

			// Execute operations
			if p.Args["usuario"] != nil {
				usuario.Usuario = p.Args["usuario"].(string)
			}
			if p.Args["correo"] != nil {
				usuario.Correo = p.Args["correo"].(string)
			}
			if p.Args["estado"] != nil {
				usuario.Estado = p.Args["estado"].(bool)
			}

			// Change password
			if p.Args["clave"] != nil || p.Args["confirmar_clave"] != nil || p.Args["clave_antigua"] != nil {
				if p.Args["clave"] == nil || p.Args["confirmar_clave"] == nil || p.Args["clave_antigua"] == nil {
					return nil, errors.New("Para cambiar la contraseña se requiere ingresar los tres valores \"clave\", \"confirmar_clave\", \"clave_antigua\"")
				} else {
					usuario.Clave = p.Args["clave"].(string)
					usuario.ConfirmarClave = p.Args["confirmar_clave"].(string)
					usuario.ClaveAntigua = p.Args["clave_antigua"].(string)

					// Validate confirm password
					if usuario.Clave != usuario.ConfirmarClave {
						return nil, errors.New("Las nuevas claves no coinciden")
					}

					// Encrypted old password
					tt := sha256.Sum256([]byte(usuario.ClaveAntigua))
					ttt := fmt.Sprintf("%x", tt)

					// Validate old password
					u := models.Usuario{}
					db.Where("clave = ?", ttt).First(&u)
					if u.ID < 1 {
						return nil, errors.New("La contraseña antigua es incorecta")
					}

					// Encrypted primary password
					cc := sha256.Sum256([]byte(usuario.Clave))
					pwd := fmt.Sprintf("%x", cc)
					usuario.Clave = pwd
					usuario.FechaModificacionClave = time.Now()
				}
			}

			if err := db.Model(&usuario).Update(usuario).Error; err != nil {
				return nil, err
			}

			return usuario, nil
		},
	}
}

func DeleteUsuarioMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.UsuarioType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			usuario := models.Usuario{}

			// get connection
			db := config.GetConnection()
			defer db.Close()

			if db.First(&usuario).RecordNotFound() {
				return nil, errors.New(fmt.Sprintf("The record with the id %d was not found", usuario.ID))
			}

			// Execute operations
			if err := db.Delete(&usuario).Error; err != nil {
				return nil, err
			}

			return usuario, nil
		},
	}
}
