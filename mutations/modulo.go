package mutations

import (
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models"
)

func CreateModuloMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.ModuloType,
		Args: graphql.FieldConfigArgument{
			"nombre":      &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"tipo":        &graphql.ArgumentConfig{Type: graphql.String},
			"descripcion": &graphql.ArgumentConfig{Type: graphql.String},
			"estado":      &graphql.ArgumentConfig{Type: graphql.Boolean},
			"semestre_id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			modulo := models.Modulo{
				Nombre:     p.Args["nombre"].(string),
				SemestreID: uint(p.Args["nombre"].(int)),
			}

			// Arguments optionals
			if p.Args["tipo"] != nil {
				modulo.Tipo = p.Args["tipo"].(string)
			}
			if p.Args["descripcion"] != nil {
				modulo.Descripcion = p.Args["descripcion"].(string)
			}
			if p.Args["estado"] != nil {
				modulo.Estado = p.Args["estado"].(bool)
			}

			// get connection
			db := config.GetConnection()
			defer db.Close()

			// Execute operations
			if err := db.Create(&modulo).Error; err != nil {
				return nil, err
			}

			return modulo, nil
		},
	}
}

func UpdateModuloMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.ModuloType,
		Args: graphql.FieldConfigArgument{
			"id":          &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			"nombre":      &graphql.ArgumentConfig{Type: graphql.String},
			"tipo":        &graphql.ArgumentConfig{Type: graphql.String},
			"descripcion": &graphql.ArgumentConfig{Type: graphql.String},
			"estado":      &graphql.ArgumentConfig{Type: graphql.Boolean},
			"semestre_id": &graphql.ArgumentConfig{Type: graphql.Int},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			modulo := models.Modulo{
				ID: uint(p.Args["id"].(int)),
			}

			// get connection
			db := config.GetConnection()
			defer db.Close()

			if db.First(&modulo).RecordNotFound() {
				return nil, errors.New(fmt.Sprintf("The record with the id %d was not found", modulo.ID))
			}

			// Arguments optionals
			if p.Args["nombre"] != nil {
				modulo.Nombre = p.Args["nombre"].(string)
			}
			if p.Args["tipo"] != nil {
				modulo.Tipo = p.Args["tipo"].(string)
			}
			if p.Args["descripcion"] != nil {
				modulo.Descripcion = p.Args["descripcion"].(string)
			}
			if p.Args["estado"] != nil {
				modulo.Estado = p.Args["estado"].(bool)
			}
			if p.Args["semestre_id"] != nil {
				modulo.SemestreID = uint(p.Args["semestre_id"].(int))
			}

			// Execute operations
			if err := db.Model(&modulo).Update(modulo).Error; err != nil {
				return nil, err
			}

			return modulo, nil
		},
	}
}

func DeleteModuloMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.ModuloType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			modulo := models.Modulo{
				ID: uint(p.Args["id"].(int)),
			}

			// get connection
			db := config.GetConnection()
			defer db.Close()

			if db.First(&modulo).RecordNotFound() {
				return nil, errors.New(fmt.Sprintf("The record with the id %d was not found", modulo.ID))
			}

			// Execute operations
			if err := db.Delete(&modulo).Error; err != nil {
				return nil, err
			}

			return modulo, nil
		},
	}
}
