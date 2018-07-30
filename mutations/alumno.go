package mutations

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models"
    "errors"
    "fmt"
    "time"
)

func CreateAlumnoMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.AlumnoType,
		Args: graphql.FieldConfigArgument{
            "dni":              &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
            "nombres":          &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
            "apellido_paterno": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
            "apellido_materno": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
            "fecha_nacimiento": &graphql.ArgumentConfig{Type: graphql.DateTime},
            "sexo":             &graphql.ArgumentConfig{Type: graphql.String},
            "pais":             &graphql.ArgumentConfig{Type: graphql.String},
            "departamento":     &graphql.ArgumentConfig{Type: graphql.String},
            "direccion":        &graphql.ArgumentConfig{Type: graphql.String},
            "celular":          &graphql.ArgumentConfig{Type: graphql.String},
            "telefono":         &graphql.ArgumentConfig{Type: graphql.String},
            "estado_civil":     &graphql.ArgumentConfig{Type: graphql.String},
            "fecha_admicion":   &graphql.ArgumentConfig{Type: graphql.DateTime},
            "fecha_promocion":  &graphql.ArgumentConfig{Type: graphql.DateTime},
            "estado":           &graphql.ArgumentConfig{Type: graphql.Boolean},
        },
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			alumno := models.Alumno{
                Dni: p.Args["dni"].(string),
                Nombres: p.Args["nombres"].(string),
                ApellidoPaterno: p.Args["apellido_paterno"].(string),
                ApellidoMaterno: p.Args["apellido_materno"].(string),
            }

            // Optional arguments
            if p.Args["fecha_nacimiento"] != nil {
                alumno.FechaNacimiento = p.Args["fecha_nacimiento"].(time.Time)
            }

            if p.Args["sexo"] != nil {
                alumno.Sexo = p.Args["sexo"].(string)
            }

            if p.Args["pais"] != nil {
                alumno.Pais = p.Args["pais"].(string)
            }

            if p.Args["departamento"] != nil {
                alumno.Departamento = p.Args["departamento"].(string)
            }

            if p.Args["direccion"] != nil {
                alumno.Direccion = p.Args["direccion"].(string)
            }

            if p.Args["celular"] != nil {
                alumno.Celular = p.Args["celular"].(string)
            }

            if p.Args["telefono"] != nil {
                alumno.Telefono = p.Args["telefono"].(string)
            }

            if p.Args["estado_civil"] != nil {
                alumno.EstadoCivil = p.Args["estado_civil"].(string)
            }

            if p.Args["fecha_admicion"] != nil {
                alumno.FechaAdmicion = p.Args["fecha_admicion"].(time.Time)
            }

            if p.Args["fecha_promocion"] != nil {
                alumno.FechaPromocion = p.Args["fecha_promocion"].(time.Time)
            }

            if p.Args["estado"] != nil {
                alumno.Estado = p.Args["estado"].(bool)
            }

			// get connection
			db := config.GetConnection()
			defer db.Close()

			// Execute operations
			if err := db.Create(&alumno).Error; err != nil {
				return nil, err
			}

			return alumno, nil
		},
	}
}

func UpdateAlumnoMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.AlumnoType,
		Args: graphql.FieldConfigArgument{
            "id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
            "dni":              &graphql.ArgumentConfig{Type: graphql.String},
            "nombres":          &graphql.ArgumentConfig{Type: graphql.String},
            "apellido_paterno": &graphql.ArgumentConfig{Type: graphql.String},
            "apellido_materno": &graphql.ArgumentConfig{Type: graphql.String},
            "fecha_nacimiento": &graphql.ArgumentConfig{Type: graphql.DateTime},
            "sexo":             &graphql.ArgumentConfig{Type: graphql.String},
            "pais":             &graphql.ArgumentConfig{Type: graphql.String},
            "departamento":     &graphql.ArgumentConfig{Type: graphql.String},
            "direccion":        &graphql.ArgumentConfig{Type: graphql.String},
            "celular":          &graphql.ArgumentConfig{Type: graphql.String},
            "telefono":         &graphql.ArgumentConfig{Type: graphql.String},
            "estado_civil":     &graphql.ArgumentConfig{Type: graphql.String},
            "fecha_admicion":   &graphql.ArgumentConfig{Type: graphql.DateTime},
            "fecha_promocion":  &graphql.ArgumentConfig{Type: graphql.DateTime},
            "estado":           &graphql.ArgumentConfig{Type: graphql.Boolean},
        },
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			alumno := models.Alumno{
                ID: uint(p.Args["id"].(int)),
            }

			// get connection
			db := config.GetConnection()
			defer db.Close()

            if db.First(&alumno).RecordNotFound() {
                return nil, errors.New(fmt.Sprintf("The record with the id %d was not found",alumno.ID))
            }

            // Optional arguments
            if p.Args["dni"] != nil {
                alumno.Dni = p.Args["dni"].(string)
            }

            if p.Args["nombres"] != nil {
                alumno.Nombres = p.Args["nombres"].(string)
            }

            if p.Args["apellido_paterno"] != nil {
                alumno.ApellidoPaterno = p.Args["apellido_paterno"].(string)
            }

            if p.Args["apellido_materno"] != nil {
                alumno.ApellidoMaterno = p.Args["apellido_materno"].(string)
            }

            if p.Args["fecha_nacimiento"] != nil {
                alumno.FechaNacimiento = p.Args["fecha_nacimiento"].(time.Time)
            }

            if p.Args["sexo"] != nil {
                alumno.Sexo = p.Args["sexo"].(string)
            }

            if p.Args["pais"] != nil {
                alumno.Pais = p.Args["pais"].(string)
            }

            if p.Args["departamento"] != nil {
                alumno.Departamento = p.Args["departamento"].(string)
            }

            if p.Args["direccion"] != nil {
                alumno.Direccion = p.Args["direccion"].(string)
            }

            if p.Args["celular"] != nil {
                alumno.Celular = p.Args["celular"].(string)
            }

            if p.Args["telefono"] != nil {
                alumno.Telefono = p.Args["telefono"].(string)
            }

            if p.Args["estado_civil"] != nil {
                alumno.EstadoCivil = p.Args["estado_civil"].(string)
            }

            if p.Args["fecha_admicion"] != nil {
                alumno.FechaAdmicion = p.Args["fecha_admicion"].(time.Time)
            }

            if p.Args["fecha_promocion"] != nil {
                alumno.FechaPromocion = p.Args["fecha_promocion"].(time.Time)
            }

            if p.Args["estado"] != nil {
                alumno.Estado = p.Args["estado"].(bool)
            }

			// Execute operations
			if err := db.Model(&alumno).Update(alumno).Error; err != nil {
				return nil, err
			}

			return alumno, nil
		},
	}
}

func DeleteAlumnoMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.AlumnoType,
		Args: graphql.FieldConfigArgument{
            "id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
        },
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			alumno := models.Alumno{
                ID: uint(p.Args["id"].(int)),
            }

			// get connection
			db := config.GetConnection()
			defer db.Close()

            if db.First(&alumno).RecordNotFound() {
                return nil, errors.New(fmt.Sprintf("The record with the id %d was not found",alumno.ID))
            }

			// Execute operations
			if err := db.Delete(&alumno).Error; err != nil {
				return nil, err
			}

			return alumno, nil
		},
	}
}
