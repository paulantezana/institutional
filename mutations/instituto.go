package mutations

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models"
)

func CreateInstitutoMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.InstitutoType,
		Args: graphql.FieldConfigArgument{
			"nombre":           &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"logo":             &graphql.ArgumentConfig{Type: graphql.String},
			"director":         &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"dregre":           &graphql.ArgumentConfig{Type: graphql.String},
			"semestre_prefijo": &graphql.ArgumentConfig{Type: graphql.String},
			"semestres":        &graphql.ArgumentConfig{Type: graphql.Int},
			"semestre_sufijo":  &graphql.ArgumentConfig{Type: graphql.String},
			"estado":           &graphql.ArgumentConfig{Type: graphql.Boolean},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			instituto := models.Instituto{
				Nombre:          p.Args["nombre"].(string),
				Director:        p.Args["director"].(string),
			}
            if p.Args["logo"] != nil {
			    instituto.Logo = p.Args["logo"].(string)
            }

            if p.Args["logo"] != nil {
                instituto.DREGRE = p.Args["dregre"].(string)
            }

            if p.Args["logo"] != nil {
                instituto.SemestrePrefijo = p.Args["semestre_prefijo"].(string)
            }

            if p.Args["logo"] != nil {
                instituto.Semestres = uint16(p.Args["semestres"].(int))
            }

            if p.Args["logo"] != nil {
                instituto.SemestrePrefijo = p.Args["semestre_sufijo"].(string)
            }

            if p.Args["logo"] != nil {
                instituto.Estado = p.Args["estado"].(bool)
            }

			// get connection
			db := config.GetConnection()
			defer db.Close()

			// Execute operations
			if err := db.Create(&instituto).Error; err != nil {
				return nil, err
			}

			return instituto, nil
		},
	}
}

func UpdateInstitutoMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.InstitutoType,
		Args: graphql.FieldConfigArgument{},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			instituto := models.Instituto{}

			// get connection
			db := config.GetConnection()
			defer db.Close()

			// Execute operations
			if err := db.Create(&instituto).Error; err != nil {
				return nil, err
			}

			return instituto, nil
		},
	}
}

func DeleteInstitutoMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.InstitutoType,
		Args: graphql.FieldConfigArgument{},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			instituto := models.Instituto{}

			// get connection
			db := config.GetConnection()
			defer db.Close()

			// Execute operations
			if err := db.Create(&instituto).Error; err != nil {
				return nil, err
			}

			return instituto, nil
		},
	}
}
