package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models"
)

func AlumnoQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(models.AlumnoType),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			db := config.GetConnection()
			defer db.Close()
			alumnos := make([]models.Alumno, 0)
			if err := db.Find(&alumnos).Error; err != nil {
				return nil, err
			}
			return alumnos, nil
		},
	}
}