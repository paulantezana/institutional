package mutations

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models"
    "errors"
    "fmt"
)

func CreateUnidadMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.UnidadType,
		Args: graphql.FieldConfigArgument{
			"nombre":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"credito":   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Float)},
			"horas":     &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			"estado":    &graphql.ArgumentConfig{Type: graphql.Boolean},
			"modulo_id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			unidad := models.Unidad{
				Nombre:   p.Args["nombre"].(string),
				Credito:  float32(p.Args["credito"].(float64)),
				Horas:    uint16(p.Args["horas"].(int)),
				ModuloID: uint(p.Args["modulo_id"].(int)),
			}

            // Arguments optionals
            if p.Args["estado"] != nil {
                unidad.Estado = p.Args["estado"].(bool)
            }

			// get connection
			db := config.GetConnection()
			defer db.Close()

			// Execute operations
			if err := db.Create(&unidad).Error; err != nil {
				return nil, err
			}

			return unidad, nil
		},
	}
}

func UpdateUnidadMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.UnidadType,
		Args: graphql.FieldConfigArgument{
			"id":        &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			"nombre":    &graphql.ArgumentConfig{Type: graphql.String},
			"credito":   &graphql.ArgumentConfig{Type: graphql.Float},
			"horas":     &graphql.ArgumentConfig{Type: graphql.Int},
			"estado":    &graphql.ArgumentConfig{Type: graphql.Boolean},
			"modulo_id": &graphql.ArgumentConfig{Type: graphql.Int},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			unidad := models.Unidad{
				ID:       uint(p.Args["id"].(int)),
			}

			// get connection
			db := config.GetConnection()
			defer db.Close()

            if db.First(&unidad).RecordNotFound() {
                return nil, errors.New(fmt.Sprintf("The record with the id %d was not found",unidad.ID))
            }

            // Arguments optionals
            if p.Args["nombre"] != nil {
                unidad.Nombre = p.Args["nombre"].(string)
            }
            if p.Args["credito"] != nil {
                unidad.Credito = float32(p.Args["credito"].(float64))
            }
            if p.Args["horas"] != nil {
                unidad.Horas = uint16(p.Args["horas"].(int))
            }
            if p.Args["estado"] != nil {
                unidad.Estado = p.Args["estado"].(bool)
            }
            if p.Args["modulo_id"] != nil {
                unidad.ModuloID =  uint(p.Args["modulo_id"].(int))
            }

			// Execute operations
			if err := db.Model(&unidad).Update(unidad).Error; err != nil {
				return nil, err
			}

			return unidad, nil
		},
	}
}

func DeleteUnidadMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.UnidadType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			unidad := models.Unidad{
				ID: uint(p.Args["id"].(int)),
			}

			// get connection
			db := config.GetConnection()
			defer db.Close()

            if db.First(&unidad).RecordNotFound() {
                return nil, errors.New(fmt.Sprintf("The record with the id %d was not found",unidad.ID))
            }

			// Execute operations
			if err := db.Delete(&unidad).Error; err != nil {
				return nil, err
			}

			return unidad, nil
		},
	}
}
