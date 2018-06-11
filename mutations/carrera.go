package mutations

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models"
    "errors"
    "fmt"
)

func CreateCarreraMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.CarreraType,
		Args: graphql.FieldConfigArgument{
			"nombre":      &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"logo":        &graphql.ArgumentConfig{Type: graphql.String},
			"descripcion": &graphql.ArgumentConfig{Type: graphql.String},
			"creacion":    &graphql.ArgumentConfig{Type: graphql.Int},
			"estado":      &graphql.ArgumentConfig{Type: graphql.Boolean},
			"filial_id":   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			carrera := models.Carrera{
				Nombre:      p.Args["nombre"].(string),
				FilialID:    uint(p.Args["filial_id"].(int)),
			}

			// Arguments optionals
            if p.Args["logo"] != nil {
                carrera.Logo = p.Args["logo"].(string)
            }
            if p.Args["descripcion"] != nil {
                carrera.Descripcion = p.Args["descripcion"].(string)
            }
            if p.Args["creacion"] != nil {
                carrera.Creacion = uint16(p.Args["creacion"].(int))
            }
            if p.Args["estado"] != nil {
                carrera.Estado = p.Args["estado"].(bool)
            }

			db := config.GetConnection()
			defer db.Close()

			// Execute operations
			if err := db.Create(&carrera).Error; err != nil {
				return nil, err
			}

			return carrera, nil
		},
	}
}

func UpdateCarreraMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.CarreraType,
		Args: graphql.FieldConfigArgument{
			"id":          &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
            "nombre":      &graphql.ArgumentConfig{Type: graphql.String},
            "logo":        &graphql.ArgumentConfig{Type: graphql.String},
            "descripcion": &graphql.ArgumentConfig{Type: graphql.String},
            "creacion":    &graphql.ArgumentConfig{Type: graphql.Int},
            "estado":      &graphql.ArgumentConfig{Type: graphql.Boolean},
            "filial_id":   &graphql.ArgumentConfig{Type: graphql.Int},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			carrera := models.Carrera{
				ID:          uint(p.Args["id"].(int)),
			}

			db := config.GetConnection()
			defer db.Close()

            if db.First(&carrera).RecordNotFound() {
                return nil, errors.New(fmt.Sprintf("The record with the id %d was not found",carrera.ID))
            }

            // Arguments optionals
            if p.Args["nombre"] != nil {
                carrera.Nombre = p.Args["nombre"].(string)
            }
            if p.Args["logo"] != nil {
                carrera.Logo = p.Args["logo"].(string)
            }
            if p.Args["descripcion"] != nil {
                carrera.Descripcion = p.Args["descripcion"].(string)
            }
            if p.Args["creacion"] != nil {
                carrera.Creacion = uint16(p.Args["creacion"].(int))
            }
            if p.Args["estado"] != nil {
                carrera.Estado = p.Args["estado"].(bool)
            }
            if p.Args["filial_id"] != nil {
                carrera.Logo = p.Args["filial_id"].(string)
            }

			// Execute operations
			if err := db.Model(&carrera).Update(carrera).Error; err != nil {
				return nil, err
			}

			return carrera, nil
		},
	}
}

func DeleteCarreraMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.CarreraType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			carrera := models.Carrera{
				ID: uint(p.Args["id"].(int)),
			}

			// get connection
			db := config.GetConnection()
			defer db.Close()

            if db.First(&carrera).RecordNotFound() {
                return nil, errors.New(fmt.Sprintf("The record with the id %d was not found",carrera.ID))
            }

			// Execute operations
			if err := db.Delete(&carrera).Error; err != nil {
				return nil, err
			}

			return carrera, nil
		},
	}
}
