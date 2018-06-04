package mutations

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models/institucional"
)

func CreateCarreraMutation() *graphql.Field {
	return &graphql.Field{
		Type: institucional.CarreraType,
		Args: graphql.FieldConfigArgument{
			"nombre":      &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"logo":        &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"descripcion": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"creacion":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			carrera := institucional.Carrera{
				Nombre:      p.Args["nombre"].(string),
				Logo:        p.Args["logo"].(string),
				Descripcion: p.Args["descripcion"].(string),
				Creacion:    uint16(p.Args["creacion"].(int)),
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
		Type: institucional.CarreraType,
		Args: graphql.FieldConfigArgument{
			"id":          &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			"nombre":      &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"logo":        &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"descripcion": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"creacion":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			carrera := institucional.Carrera{
				ID:          p.Args["nombre"].(uint),
				Nombre:      p.Args["nombre"].(string),
				Logo:        p.Args["logo"].(string),
				Descripcion: p.Args["descripcion"].(string),
				Creacion:    uint16(p.Args["creacion"].(int)),
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
