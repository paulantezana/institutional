package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models"
)

func AlumnoQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(models.AlumnoType),
        Args: graphql.FieldConfigArgument{
            "search":          &graphql.ArgumentConfig{Type: graphql.String},
        },
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		    // Declare Variables
            search := ""
			alumnos := make([]models.Alumno, 0)

		    // Query Params
            if p.Args["search"] != nil {
                search = p.Args["search"].(string)
            }

            // Get connection
			db := config.GetConnection()
			defer db.Close()

            // Execute instructions
            if err := db.Where("dni LIKE ?","%"+search+"%").Or("nombres LIKE ?","%"+search+"%").Or("apellido_paterno LIKE ?","%"+search+"%").Or("apellido_materno LIKE ?","%"+search+"%") .Find(&alumnos).Error; err != nil {
                return nil, err
            }

			return alumnos, nil
		},
	}
}

func AlumnoIDQuery() *graphql.Field {
    return &graphql.Field{
        Type: models.AlumnoType,
        Args: graphql.FieldConfigArgument{
            "id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
        },
        Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            alumno := models.Alumno{}
            id := p.Args["id"].(int)

            // Get connection
            db := config.GetConnection()
            defer db.Close()

            // Execute instructions
            if err := db.First(&alumno, id).Error; err != nil {
                return nil, err
            }
            return alumno, nil
        },
    }
}
