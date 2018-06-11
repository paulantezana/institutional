package mutations

import (
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models"
	"time"
)

func CreateFilialMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.FilialType,
		Args: graphql.FieldConfigArgument{
			"nombre":                         &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"pais":                           &graphql.ArgumentConfig{Type: graphql.String},
			"departamento":                   &graphql.ArgumentConfig{Type: graphql.String},
			"provincia":                      &graphql.ArgumentConfig{Type: graphql.String},
			"distrito":                       &graphql.ArgumentConfig{Type: graphql.String},
			"centro_poblado":                 &graphql.ArgumentConfig{Type: graphql.String},
			"direccion_sede":                 &graphql.ArgumentConfig{Type: graphql.String},
			"direccion_matricula":            &graphql.ArgumentConfig{Type: graphql.String},
			"codigo_modular":                 &graphql.ArgumentConfig{Type: graphql.String},
			"tipo_gestion":                   &graphql.ArgumentConfig{Type: graphql.String},
			"mension":                        &graphql.ArgumentConfig{Type: graphql.String},
			"resolucion_autorizacion_tipo":   &graphql.ArgumentConfig{Type: graphql.String},
			"resolucion_autorizacion_numero": &graphql.ArgumentConfig{Type: graphql.String},
			"resolucion_autorizacion_fecha":  &graphql.ArgumentConfig{Type: graphql.DateTime},
			"resolucion_revalidacion_tipo":   &graphql.ArgumentConfig{Type: graphql.String},
			"resolucion_revalidacion_numero": &graphql.ArgumentConfig{Type: graphql.String},
			"resolucion_revalidacion_fecha":  &graphql.ArgumentConfig{Type: graphql.DateTime},
			"instituto_id":                   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			"estado":                         &graphql.ArgumentConfig{Type: graphql.Boolean},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			filial := models.Filial{
				Nombre:      p.Args["nombre"].(string),
				InstitutoID: uint(p.Args["instituto_id"].(int)),
			}

			if p.Args["pais"] != nil {
				filial.Pais = p.Args["pais"].(string)
			}
			if p.Args["departamento"] != nil {
				filial.Departamento = p.Args["departamento"].(string)
			}
			if p.Args["provincia"] != nil {
				filial.Provincia = p.Args["provincia"].(string)
			}
			if p.Args["distrito"] != nil {
				filial.Distrito = p.Args["distrito"].(string)
			}
			if p.Args["centro_poblado"] != nil {
				filial.CentroPoblado = p.Args["centro_poblado"].(string)
			}
			if p.Args["direccion_sede"] != nil {
				filial.DireccionSede = p.Args["direccion_sede"].(string)
			}
			if p.Args["direccion_matricula"] != nil {
				filial.DireccionMatricula = p.Args["direccion_matricula"].(string)
			}
			if p.Args["codigo_modular"] != nil {
				filial.CodigoModular = p.Args["codigo_modular"].(string)
			}
			if p.Args["tipo_gestion"] != nil {
				filial.TipoGestion = p.Args["tipo_gestion"].(string)
			}
			if p.Args["mension"] != nil {
				filial.Mension = p.Args["mension"].(string)
			}
			if p.Args["resolucion_autorizacion_tipo"] != nil {
				filial.ResolucionAutorizacionTipo = p.Args["resolucion_autorizacion_tipo"].(string)
			}
			if p.Args["resolucion_autorizacion_numero"] != nil {
				filial.ResolucionAutorizacionNumero = p.Args["resolucion_autorizacion_numero"].(string)
			}
			if p.Args["resolucion_autorizacion_fecha"] != nil {
				filial.ResolucionAutorizacionFecha = p.Args["resolucion_autorizacion_fecha"].(time.Time)
			}
			if p.Args["resolucion_revalidacion_tipo"] != nil {
				filial.ResolucionRevalidacionTipo = p.Args["resolucion_revalidacion_tipo"].(string)
			}
			if p.Args["resolucion_revalidacion_numero"] != nil {
				filial.ResolucionRevalidacionNumero = p.Args["resolucion_revalidacion_numero"].(string)
			}
			if p.Args["resolucion_revalidacion_fecha"] != nil {
				filial.ResolucionRevalidacionFecha = p.Args["resolucion_revalidacion_fecha"].(time.Time)
			}
			if p.Args["estado"] != nil {
				filial.Estado = p.Args["estado"].(bool)
			}

			// get connection
			db := config.GetConnection()
			defer db.Close()

			// Execute operations
			if err := db.Create(&filial).Error; err != nil {
				return nil, err
			}

			return filial, nil
		},
	}
}

func UpdateFilialMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.FilialType,
		Args: graphql.FieldConfigArgument{
			"id":                             &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			"nombre":                         &graphql.ArgumentConfig{Type: graphql.String},
			"pais":                           &graphql.ArgumentConfig{Type: graphql.String},
			"departamento":                   &graphql.ArgumentConfig{Type: graphql.String},
			"provincia":                      &graphql.ArgumentConfig{Type: graphql.String},
			"distrito":                       &graphql.ArgumentConfig{Type: graphql.String},
			"centro_poblado":                 &graphql.ArgumentConfig{Type: graphql.String},
			"direccion_sede":                 &graphql.ArgumentConfig{Type: graphql.String},
			"direccion_matricula":            &graphql.ArgumentConfig{Type: graphql.String},
			"codigo_modular":                 &graphql.ArgumentConfig{Type: graphql.String},
			"tipo_gestion":                   &graphql.ArgumentConfig{Type: graphql.String},
			"mension":                        &graphql.ArgumentConfig{Type: graphql.String},
			"resolucion_autorizacion_tipo":   &graphql.ArgumentConfig{Type: graphql.String},
			"resolucion_autorizacion_numero": &graphql.ArgumentConfig{Type: graphql.String},
			"resolucion_autorizacion_fecha":  &graphql.ArgumentConfig{Type: graphql.DateTime},
			"resolucion_revalidacion_tipo":   &graphql.ArgumentConfig{Type: graphql.String},
			"resolucion_revalidacion_numero": &graphql.ArgumentConfig{Type: graphql.String},
			"resolucion_revalidacion_fecha":  &graphql.ArgumentConfig{Type: graphql.DateTime},
			"instituto_id":                   &graphql.ArgumentConfig{Type: graphql.Int},
			"estado":                         &graphql.ArgumentConfig{Type: graphql.Boolean},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			filial := models.Filial{
				ID: uint(p.Args["id"].(int)),
			}

			// get connection
			db := config.GetConnection()
			defer db.Close()

			if db.First(&filial).RecordNotFound() {
				return nil, errors.New(fmt.Sprintf("The record with the id %d was not found", filial.ID))
			}

			// optional arguments
			if p.Args["nombre"] != nil {
				filial.Nombre = p.Args["nombre"].(string)
			}
			if p.Args["pais"] != nil {
				filial.Pais = p.Args["pais"].(string)
			}
			if p.Args["departamento"] != nil {
				filial.Departamento = p.Args["departamento"].(string)
			}
			if p.Args["provincia"] != nil {
				filial.Provincia = p.Args["provincia"].(string)
			}
			if p.Args["distrito"] != nil {
				filial.Distrito = p.Args["distrito"].(string)
			}
			if p.Args["centro_poblado"] != nil {
				filial.CentroPoblado = p.Args["centro_poblado"].(string)
			}
			if p.Args["direccion_sede"] != nil {
				filial.DireccionSede = p.Args["direccion_sede"].(string)
			}
			if p.Args["direccion_matricula"] != nil {
				filial.DireccionMatricula = p.Args["direccion_matricula"].(string)
			}
			if p.Args["codigo_modular"] != nil {
				filial.CodigoModular = p.Args["codigo_modular"].(string)
			}
			if p.Args["tipo_gestion"] != nil {
				filial.TipoGestion = p.Args["tipo_gestion"].(string)
			}
			if p.Args["mension"] != nil {
				filial.Mension = p.Args["mension"].(string)
			}
			if p.Args["resolucion_autorizacion_tipo"] != nil {
				filial.ResolucionAutorizacionTipo = p.Args["resolucion_autorizacion_tipo"].(string)
			}
			if p.Args["resolucion_autorizacion_numero"] != nil {
				filial.ResolucionAutorizacionNumero = p.Args["resolucion_autorizacion_numero"].(string)
			}
			if p.Args["resolucion_autorizacion_fecha"] != nil {
				filial.ResolucionAutorizacionFecha = p.Args["resolucion_autorizacion_fecha"].(time.Time)
			}
			if p.Args["resolucion_revalidacion_tipo"] != nil {
				filial.ResolucionRevalidacionTipo = p.Args["resolucion_revalidacion_tipo"].(string)
			}
			if p.Args["resolucion_revalidacion_numero"] != nil {
				filial.ResolucionRevalidacionNumero = p.Args["resolucion_revalidacion_numero"].(string)
			}
			if p.Args["resolucion_revalidacion_fecha"] != nil {
				filial.ResolucionRevalidacionFecha = p.Args["resolucion_revalidacion_fecha"].(time.Time)
			}
			if p.Args["estado"] != nil {
				filial.Estado = p.Args["estado"].(bool)
			}
			if p.Args["estado"] != nil {
				filial.InstitutoID = uint(p.Args["instituto_id"].(int))
			}

			// Execute operations
			if err := db.Model(&filial).Update(filial).Error; err != nil {
				return nil, err
			}

			return filial, nil
		},
	}
}

func DeleteFilialMutation() *graphql.Field {
	return &graphql.Field{
		Type: models.FilialType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			filial := models.Filial{
				ID: uint(p.Args["id"].(int)),
			}

			// get connection
			db := config.GetConnection()
			defer db.Close()

			if db.First(&filial).RecordNotFound() {
				return nil, errors.New(fmt.Sprintf("The record with the id %d was not found", filial.ID))
			}

			// Execute operations
			if err := db.Delete(&filial).Error; err != nil {
				return nil, err
			}

			return filial, nil
		},
	}
}
