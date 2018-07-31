package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models"
)

func SemestreQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(models.SemestreType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			db := config.GetConnection()
			defer db.Close()
			semestres := make([]models.Semestre, 0)
			if err := db.Find(&semestres).Error; err != nil {
				return nil, err
			}
			return semestres, nil
		},
	}
}

func SemestreIDQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(models.SemestreType),
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			semestre := models.Semestre{}
			id := p.Args["id"].(int)

			// Get connection
			db := config.GetConnection()
			defer db.Close()

			// Execute instructions
			if err := db.First(&semestre, id).Error; err != nil {
				return nil, err
			}
			return semestre, nil
		},
	}
}
