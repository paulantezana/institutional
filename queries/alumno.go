package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models"
    "fmt"
)

func AlumnoQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(models.AlumnoType),
        Args: graphql.FieldConfigArgument{
            "dni":              &graphql.ArgumentConfig{Type: graphql.String},
            //"nombres":          &graphql.ArgumentConfig{Type: graphql.String},
        },
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		    a := models.Alumno{}
            if p.Args["dni"] != nil {
               a.Dni = p.Args["dni"].(string)
            }

            //if p.Args["nombres"] != nil {
            //    alumno.Nombres = p.Args["nombres"].(string)
            //}

			db := config.GetConnection()
			defer db.Close()

			alumnos := make([]models.Alumno, 0)

			if err := db.Where("name LIKE ?", fmt.Sprint("%"+a.Dni+"%")).Find(&alumnos).Error; err != nil {
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
