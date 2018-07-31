package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models"
)

func CarreraQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(models.CarreraType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			db := config.GetConnection()
			defer db.Close()
			carreras := make([]models.Carrera, 0)
			if err := db.Find(&carreras).Error; err != nil {
				return nil, err
			}
			return carreras, nil
		},
	}
}

func CarreraIDQuery() *graphql.Field {
	return &graphql.Field{
		Type: models.CarreraType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			carrera := models.Carrera{}
			id := p.Args["id"].(int)

			// Get connection
			db := config.GetConnection()
			defer db.Close()

			// Execute instructions
			if err := db.First(&carrera, id).Error; err != nil {
				return nil, err
			}
			return carrera, nil
		},
	}
}
