package mutations

import (
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models"
)

func CreateCoreRolMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.CoreRolType,
		Args: graphql.FieldConfigArgument{
			"name":        &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"description": &graphql.ArgumentConfig{Type: graphql.String},
			"estado":      &graphql.ArgumentConfig{Type: graphql.Boolean},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			rol := models.CoreRol{
				Name: p.Args["name"].(string),
			}

			// Optional arguments
			if p.Args["description"] != nil {
				rol.Description = p.Args["description"].(string)
			}
			if p.Args["estado"] != nil {
				rol.Estado = p.Args["estado"].(bool)
			}

			// get connection
			db := config.GetConnection()
			defer db.Close()

			// Execute operations
			if err := db.Create(&rol).Error; err != nil {
				return nil, err
			}

			return rol, nil
		},
	}
}

func UpdateCoreRolMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.CoreRolType,
		Args: graphql.FieldConfigArgument{
			"id":          &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			"name":        &graphql.ArgumentConfig{Type: graphql.String},
			"description": &graphql.ArgumentConfig{Type: graphql.String},
			"estado":      &graphql.ArgumentConfig{Type: graphql.Boolean},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			rol := models.CoreRol{
				ID: uint(p.Args["id"].(int)),
			}

			// get connection
			db := config.GetConnection()
			defer db.Close()

			if db.First(&rol).RecordNotFound() {
				return nil, errors.New(fmt.Sprintf("The record with the id %d was not found", rol.ID))
			}

			// Optional arguments
			if p.Args["name"] != nil {
				rol.Name = p.Args["name"].(string)
			}
			if p.Args["description"] != nil {
				rol.Description = p.Args["description"].(string)
			}
			if p.Args["estado"] != nil {
				rol.Estado = p.Args["estado"].(bool)
			}

			// Execute operations
			if err := db.Model(&rol).Update(rol).Error; err != nil {
				return nil, err
			}

			return rol, nil
		},
	}
}

func DeleteCoreRolMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.CoreRolType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			rol := models.CoreRol{
				ID: uint(p.Args["id"].(int)),
			}

			// get connection
			db := config.GetConnection()
			defer db.Close()

			// Validations
			if db.First(&rol).RecordNotFound() {
				return nil, errors.New(fmt.Sprintf("The record with the id %d was not found", rol.ID))
			}
			if rol.Freeze {
				return nil, errors.New(fmt.Sprintf("this %s role is not allowed to eliminate", rol.Name))
			}

			// Execute operations
			if err := db.Delete(&rol).Error; err != nil {
				return nil, err
			}

			return rol, nil
		},
	}
}
