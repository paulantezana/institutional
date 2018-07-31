package mutations

import (
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models"
    "errors"
    "fmt"
    "time"
)

func CreateProfesorMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.ProfesorType,
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
            "foto":              &graphql.ArgumentConfig{Type: graphql.String},
            "condicion_laboral": &graphql.ArgumentConfig{Type: graphql.String},
            "nivel_educativo":   &graphql.ArgumentConfig{Type: graphql.String},
            "fecha_ingreso":     &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.DateTime)},
            "fecha_retiro":      &graphql.ArgumentConfig{Type: graphql.DateTime},
            "anos_docencia":     &graphql.ArgumentConfig{Type: graphql.Int},
            "meses_docencia":    &graphql.ArgumentConfig{Type: graphql.Int},
            "especialidad":      &graphql.ArgumentConfig{Type: graphql.String},
            "numero_horas":      &graphql.ArgumentConfig{Type: graphql.Int},
            "estado":           &graphql.ArgumentConfig{Type: graphql.Boolean},
        },
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
            profesor := models.Profesor{
                Dni: p.Args["dni"].(string),
                Nombres: p.Args["nombres"].(string),
                ApellidoPaterno: p.Args["apellido_paterno"].(string),
                ApellidoMaterno: p.Args["apellido_materno"].(string),
            }

            // Optional arguments
            if p.Args["fecha_nacimiento"] != nil {
                profesor.FechaNacimiento = p.Args["fecha_nacimiento"].(time.Time)
            }

            if p.Args["sexo"] != nil {
                profesor.Sexo = p.Args["sexo"].(string)
            }

            if p.Args["pais"] != nil {
                profesor.Pais = p.Args["pais"].(string)
            }

            if p.Args["departamento"] != nil {
                profesor.Departamento = p.Args["departamento"].(string)
            }

            if p.Args["direccion"] != nil {
                profesor.Direccion = p.Args["direccion"].(string)
            }

            if p.Args["celular"] != nil {
                profesor.Celular = p.Args["celular"].(string)
            }

            if p.Args["telefono"] != nil {
                profesor.Telefono = p.Args["telefono"].(string)
            }

            if p.Args["estado_civil"] != nil {
                profesor.EstadoCivil = p.Args["estado_civil"].(string)
            }

            if p.Args["foto"] != nil {
                profesor.Foto = p.Args["foto"].(string)
            }

            if p.Args["condicion_laboral"] != nil {
                profesor.CondicionLaboral = p.Args["condicion_laboral"].(string)
            }

            if p.Args["nivel_educativo"] != nil {
                profesor.NivelEducativo = p.Args["nivel_educativo"].(string)
            }

            if p.Args["fecha_ingreso"] != nil {
                profesor.FechaIngreso = p.Args["fecha_ingreso"].(time.Time)
            }

            if p.Args["fecha_retiro"] != nil {
                profesor.FechaRetiro = p.Args["fecha_retiro"].(time.Time)
            }

            if p.Args["anos_docencia"] != nil {
                profesor.AnosDocencia = int32(p.Args["anos_docencia"].(int))
            }

            if p.Args["meses_docencia"] != nil {
                profesor.MesesDocencia = int32(p.Args["meses_docencia"].(int))
            }

            if p.Args["especialidad"] != nil {
                profesor.Especialidad = p.Args["especialidad"].(string)
            }

            if p.Args["numero_horas"] != nil {
                profesor.NumeroHoras = int32(p.Args["numero_horas"].(int))
            }

            if p.Args["estado"] != nil {
                profesor.Estado = p.Args["estado"].(bool)
            }

            // Optional arguments

			// get connection
			db := config.GetConnection()
			defer db.Close()

			// Execute operations
			if err := db.Create(&profesor).Error; err != nil {
				return nil, err
			}

			return profesor, nil
		},
	}
}

func UpdateProfesorMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.ProfesorType,
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
            "foto":              &graphql.ArgumentConfig{Type: graphql.String},
            "condicion_laboral": &graphql.ArgumentConfig{Type: graphql.String},
            "nivel_educativo":   &graphql.ArgumentConfig{Type: graphql.String},
            "fecha_ingreso":     &graphql.ArgumentConfig{Type: graphql.DateTime},
            "fecha_retiro":      &graphql.ArgumentConfig{Type: graphql.DateTime},
            "anos_docencia":     &graphql.ArgumentConfig{Type: graphql.Int},
            "meses_docencia":    &graphql.ArgumentConfig{Type: graphql.Int},
            "especialidad":      &graphql.ArgumentConfig{Type: graphql.String},
            "numero_horas":      &graphql.ArgumentConfig{Type: graphql.Int},
            "estado":           &graphql.ArgumentConfig{Type: graphql.Boolean},
        },
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			profesor := models.Profesor{
                ID: uint(p.Args["id"].(int)),
            }

            // Optional arguments
            if p.Args["fecha_nacimiento"] != nil {
                profesor.FechaNacimiento = p.Args["fecha_nacimiento"].(time.Time)
            }

            if p.Args["sexo"] != nil {
                profesor.Sexo = p.Args["sexo"].(string)
            }

            if p.Args["pais"] != nil {
                profesor.Pais = p.Args["pais"].(string)
            }

            if p.Args["departamento"] != nil {
                profesor.Departamento = p.Args["departamento"].(string)
            }

            if p.Args["direccion"] != nil {
                profesor.Direccion = p.Args["direccion"].(string)
            }

            if p.Args["celular"] != nil {
                profesor.Celular = p.Args["celular"].(string)
            }

            if p.Args["telefono"] != nil {
                profesor.Telefono = p.Args["telefono"].(string)
            }

            if p.Args["estado_civil"] != nil {
                profesor.EstadoCivil = p.Args["estado_civil"].(string)
            }

            if p.Args["foto"] != nil {
                profesor.Foto = p.Args["foto"].(string)
            }

            if p.Args["condicion_laboral"] != nil {
                profesor.CondicionLaboral = p.Args["condicion_laboral"].(string)
            }

            if p.Args["nivel_educativo"] != nil {
                profesor.NivelEducativo = p.Args["nivel_educativo"].(string)
            }

            if p.Args["fecha_ingreso"] != nil {
                profesor.FechaIngreso = p.Args["fecha_ingreso"].(time.Time)
            }

            if p.Args["fecha_retiro"] != nil {
                profesor.FechaRetiro = p.Args["fecha_retiro"].(time.Time)
            }

            if p.Args["anos_docencia"] != nil {
                profesor.AnosDocencia = int32(p.Args["anos_docencia"].(int))
            }

            if p.Args["meses_docencia"] != nil {
                profesor.MesesDocencia = int32(p.Args["meses_docencia"].(int))
            }

            if p.Args["especialidad"] != nil {
                profesor.Especialidad = p.Args["especialidad"].(string)
            }

            if p.Args["numero_horas"] != nil {
                profesor.NumeroHoras = int32(p.Args["numero_horas"].(int))
            }

            if p.Args["estado"] != nil {
                profesor.Estado = p.Args["estado"].(bool)
            }

            // get connection
			db := config.GetConnection()
			defer db.Close()

            if db.First(&profesor).RecordNotFound() {
                return nil, errors.New(fmt.Sprintf("The record with the id %d was not found",profesor.ID))
            }

            // Optional arguments

			// Execute operations
			if err := db.Model(&profesor).Update(profesor).Error; err != nil {
				return nil, err
			}

			return profesor, nil
		},
	}
}

func DeleteProfesorMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.ProfesorType,
		Args: graphql.FieldConfigArgument{
            "id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
        },
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			profesor := models.Profesor{}

			// get connection
			db := config.GetConnection()
			defer db.Close()

            if db.First(&profesor).RecordNotFound() {
                return nil, errors.New(fmt.Sprintf("The record with the id %d was not found",profesor.ID))
            }

			// Execute operations
			if err := db.Delete(&profesor).Error; err != nil {
				return nil, err
			}

			return profesor, nil
		},
	}
}
