package mutations

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models"
    "errors"
    "fmt"
    "time"
    "crypto/sha256"
)

func CreateUsuarioMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.UsuarioType,
		Args: graphql.FieldConfigArgument{
            "usuario":                  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
            "correo":                  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
            "clave":                    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
            "confirmar_clave":          &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
            "clave_recuperar":          &graphql.ArgumentConfig{Type: graphql.String},
            "fecha_modificacion_clave": &graphql.ArgumentConfig{Type: graphql.DateTime},
            "estado":                   &graphql.ArgumentConfig{Type: graphql.Boolean},
        },
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			usuario := models.Usuario{
                Usuario: p.Args["usuario"].(string),
                Correo: p.Args["correo"].(string),
                Clave:   p.Args["clave"].(string),
                ConfirmarClave:   p.Args["confirmar_clave"].(string),
            }

            // Optional arguments
            if p.Args["clave_recuperar"] != nil {
                usuario.FechaModificacionClave = p.Args["clave_recuperar"].(time.Time)
            }

            if p.Args["fecha_modificacion_clave"] != nil {
                usuario.FechaModificacionClave = p.Args["fecha_modificacion_clave"].(time.Time)
            }
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
            "id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
            "usuario":                  &graphql.ArgumentConfig{Type: graphql.String},
            "correo":                  &graphql.ArgumentConfig{Type: graphql.String},
            "clave":                    &graphql.ArgumentConfig{Type: graphql.String},
            "confirmar_clave":          &graphql.ArgumentConfig{Type: graphql.String},
            "clave_antigua":          &graphql.ArgumentConfig{Type: graphql.String},
            "clave_recuperar":          &graphql.ArgumentConfig{Type: graphql.String},
            "fecha_modificacion_clave": &graphql.ArgumentConfig{Type: graphql.DateTime},
            "estado":                   &graphql.ArgumentConfig{Type: graphql.Boolean},
        },
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			usuario := models.Usuario{
                ID: uint(p.Args["id"].(int)),
            }

			// get connection
			db := config.GetConnection()
			defer db.Close()

            if db.First(&usuario).RecordNotFound() {
                return nil, errors.New(fmt.Sprintf("The record with the id %d was not found",usuario.ID))
            }

			// Execute operations
            if p.Args["usuario"] != nil {
                usuario.Usuario = p.Args["usuario"].(string)
            }
            if p.Args["correo"] != nil {
                usuario.Correo = p.Args["correo"].(string)
            }
            if p.Args["clave_recuperar"] != nil {
                usuario.ClaveRecuperar = p.Args["clave_recuperar"].(string)
            }
            if p.Args["fecha_modificacion_clave"] != nil {
                usuario.FechaModificacionClave = p.Args["fecha_modificacion_clave"].(time.Time)
            }
            if p.Args["estado"] != nil {
                usuario.FechaModificacionClave = p.Args["estado"].(time.Time)
            }

            // Change password
            if p.Args["clave"] != nil || p.Args["confirmar_clave"] != nil || p.Args["clave_antigua"] != nil {
                if p.Args["clave"] == nil || p.Args["confirmar_clave"] == nil || p.Args["clave_antigua"] == nil{
                    return nil, errors.New("Para cambiar la contraseña se requiere ingresar los tres valores \"clave\", \"confirmar_clave\", \"clave_antigua\"")
                }else{
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
                    u:= models.Usuario{}
                    db.Where("clave = ?", ttt).First(&u)
                    if u.ID < 1 {
                        return nil, errors.New("La contraseña antigua es incorecta")
                    }

                    // Encrypted primary password
                    cc := sha256.Sum256([]byte(usuario.Clave))
                    pwd := fmt.Sprintf("%x", cc)
                    usuario.Clave = pwd
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
                return nil, errors.New(fmt.Sprintf("The record with the id %d was not found",usuario.ID))
            }

			// Execute operations
			if err := db.Delete(&usuario).Error; err != nil {
				return nil, err
			}

			return usuario, nil
		},
	}
}

func RecoverPassword() *graphql.Field {
    return &graphql.Field {
        Type: models.UsuarioType,
        Args: graphql.FieldConfigArgument{
            "correo":                  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            usuario := models.Usuario{
                Correo: p.Args["correo"].(string),
            }

            // get connection
            db := config.GetConnection()
            defer db.Close()

            // Validations
            if err := db.First(&usuario).Error; err != nil {
                return nil, err
            }
            if usuario.ID < 1 {
                return nil, errors.New(fmt.Sprintf("Usuario %s no encontrado",usuario.Correo))
            }

            //

            return usuario, nil
        },
    }
}