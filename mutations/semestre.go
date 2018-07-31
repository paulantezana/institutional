package mutations

import (
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models"
)

func CreateSemestreMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.SemestreType,
		Args: graphql.FieldConfigArgument{
			"nombre":     &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"year":       &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			"estado":     &graphql.ArgumentConfig{Type: graphql.Boolean},
			"carrera_id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			semestre := models.Semestre{
				Nombre:    p.Args["nombre"].(string),
				Year:      uint16(p.Args["year"].(int)),
				CarreraID: uint(p.Args["carrera_id"].(int)),
			}

			// Arguments optionals
			if p.Args["estado"] != nil {
				semestre.Estado = p.Args["estado"].(bool)
			}

			// get connection
			db := config.GetConnection()
			defer db.Close()

			// Execute operations
			if err := db.Create(&semestre).Error; err != nil {
				return nil, err
			}

			return semestre, nil
		},
	}
}

func UpdateSemestreMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.SemestreType,
		Args: graphql.FieldConfigArgument{
			"id":         &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			"nombre":     &graphql.ArgumentConfig{Type: graphql.String},
			"year":       &graphql.ArgumentConfig{Type: graphql.Int},
			"estado":     &graphql.ArgumentConfig{Type: graphql.Boolean},
			"carrera_id": &graphql.ArgumentConfig{Type: graphql.Int},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			semestre := models.Semestre{
				ID: uint(p.Args["id"].(int)),
			}

			// get connection
			db := config.GetConnection()
			defer db.Close()

			if db.First(&semestre).RecordNotFound() {
				return nil, errors.New(fmt.Sprintf("The record with the id %d was not found", semestre.ID))
			}

			// Arguments optionals
			if p.Args["nombre"] != nil {
				semestre.Nombre = p.Args["nombre"].(string)
			}
			if p.Args["year"] != nil {
				semestre.Year = uint16(p.Args["year"].(int))
			}
			if p.Args["carrera_id"] != nil {
				semestre.CarreraID = uint(p.Args["carrera_id"].(int))
			}
			if p.Args["estado"] != nil {
				semestre.Estado = p.Args["estado"].(bool)
			}

			// Execute operations
			if err := db.Model(&semestre).Update(semestre).Error; err != nil {
				return nil, err
			}

			return semestre, nil
		},
	}
}

func DeleteSemestreMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.SemestreType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			semestre := models.Semestre{
				ID: uint(p.Args["id"].(int)),
			}

			// get connection
			db := config.GetConnection()
			defer db.Close()

			if db.First(&semestre).RecordNotFound() {
				return nil, errors.New(fmt.Sprintf("The record with the id %d was not found", semestre.ID))
			}

			// Execute operations
			if err := db.Delete(&semestre).Error; err != nil {
				return nil, err
			}

			return semestre, nil
		},
	}
}
