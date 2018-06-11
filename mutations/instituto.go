package mutations

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models"
    "errors"
    "fmt"
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
				Nombre:   p.Args["nombre"].(string),
				Director: p.Args["director"].(string),
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
		Args: graphql.FieldConfigArgument{
			"id":               &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			"nombre":           &graphql.ArgumentConfig{Type: graphql.String},
			"logo":             &graphql.ArgumentConfig{Type: graphql.String},
			"director":         &graphql.ArgumentConfig{Type: graphql.String},
			"dregre":           &graphql.ArgumentConfig{Type: graphql.String},
			"semestre_prefijo": &graphql.ArgumentConfig{Type: graphql.String},
			"semestres":        &graphql.ArgumentConfig{Type: graphql.Int},
			"semestre_sufijo":  &graphql.ArgumentConfig{Type: graphql.String},
			"estado":           &graphql.ArgumentConfig{Type: graphql.Boolean},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			instituto := models.Instituto{
				ID:       uint(p.Args["id"].(int)),
			}

            // get connection
            db := config.GetConnection()
            defer db.Close()

            if db.First(&instituto).RecordNotFound() {
                return nil, errors.New(fmt.Sprintf("The record with the id %d was not found",instituto.ID))
            }

            if p.Args["nombre"] != nil {
                instituto.Nombre = p.Args["nombre"].(string)
            }

			if p.Args["logo"] != nil {
				instituto.Logo = p.Args["logo"].(string)
			}

            if p.Args["director"] != nil {
                instituto.Director = p.Args["director"].(string)
            }

			if p.Args["dregre"] != nil {
				instituto.DREGRE = p.Args["dregre"].(string)
			}

			if p.Args["semestre_prefijo"] != nil {
				instituto.SemestrePrefijo = p.Args["semestre_prefijo"].(string)
			}

			if p.Args["semestres"] != nil {
				instituto.Semestres = uint16(p.Args["semestres"].(int))
			}

			if p.Args["semestre_sufijo"] != nil {
				instituto.SemestreSufijo = p.Args["semestre_sufijo"].(string)
			}

			if p.Args["estado"] != nil {
				instituto.Estado = p.Args["estado"].(bool)
			}

			// Execute operations
			if err := db.Model(&instituto).Update(instituto).Error; err != nil {
				return nil, err
			}

			return instituto, nil
		},
	}
}

func DeleteInstitutoMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.InstitutoType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			instituto := models.Instituto{
				ID: uint(p.Args["id"].(int)),
			}

			// get connection
			db := config.GetConnection()
			defer db.Close()

            if db.First(&instituto).RecordNotFound() {
                return nil, errors.New(fmt.Sprintf("The record with the id %d was not found",instituto.ID))
            }

			// Execute operations
			if err := db.Delete(&instituto).Error; err != nil {
				return nil, err
			}

			return instituto, nil
		},
	}
}
