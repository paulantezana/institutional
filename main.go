package main

import (
	"flag"
	"fmt"
	"github.com/graphql-go/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/paulantezana/institutional/api"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models"
	"github.com/paulantezana/institutional/schemas"
	"github.com/paulantezana/institutional/security"
	"os"
)

func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera la migración a la BD")
	flag.Parse()
	//if migrate == "yes" {
	fmt.Println("Init migration")
	Migrate()
	fmt.Println("Finalize migration")
	//}

	// Create new server mux
	e := echo.New()

	// COR
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		//AllowHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"},
		AllowMethods: []string{echo.GET, echo.POST},
	}))

	// Custom port
	port := os.Getenv("PORT")
	if port == "" {
		port = config.GetConfig().Server.Port
	}

	// GRAPH QL ========================================================================
	h := handler.New(&handler.Config{
		Schema:   &schemas.Schema,
		Pretty:   true,
		GraphiQL: true,
	})

	gq := e.Group("/graphql")
	gq.POST("", echo.WrapHandler(h))

	// API REST ==========================================================================
	ar := e.Group("")
	ar.Use(middleware.Logger())
	ar.Use(middleware.Recover())

	ar.POST("/login", security.Login)
	ar.POST("/forgout/serach", api.ForgoutSearch)
	ar.POST("/forgout/validate", api.ForgoutValidate)
	ar.POST("/forgout/change", api.ForgoutChange)
	ar.POST("/register", api.RegisterUser)

	// STATIC FILES =======================================================================
	st := e.Group("")
	st.Static("/", "public")

	// Start Server
	e.Logger.Fatal(e.Start(":" + port))
}

func Migrate() {
	db := config.GetConnection()
	defer db.Close()

	db.AutoMigrate(
		&models.Core{},
		&models.CoreMenu{},
		&models.CoreModulo{},
		&models.CorePermit{},
		&models.CoreRol{},
		&models.CoreSubMenu{},

		&models.Instituto{},
		&models.Filial{},
		&models.Carrera{},
		&models.Semestre{},
		&models.Modulo{},
		&models.Unidad{},

		&models.Usuario{},
		&models.Alumno{},
		&models.Profesor{},

		&models.Monto{},
		&models.Pago{},
		&models.Beca{},
		&models.Matricula{},

		&models.Empresa{},
		&models.Practica{},
		&models.PracticaNota{},
		&models.Representante{},

		&models.Personal{},

		&models.Nota{},

		&models.CategoriaLibro{},
		&models.Libro{},
		&models.Prestamo{},
	)

	//db.Model(&models.Filial{}).AddForeignKey("instituto_id", "institutos(id)", "RESTRICT", "RESTRICT")
	//db.Model(&models.Carrera{}).AddForeignKey("filial_id", "filiales(id)", "RESTRICT", "RESTRICT")
	//db.Model(&models.Semestre{}).AddForeignKey("carrera_id", "carreras(id)", "RESTRICT", "RESTRICT")
	//db.Model(&models.Modulo{}).AddForeignKey("semestre_id", "semestres(id)", "RESTRICT", "RESTRICT")
	//db.Model(&models.Unidad{}).AddForeignKey("modulo_id", "modulos(id)", "RESTRICT", "RESTRICT")
	//
	//db.Model(&models.CorePermit{}).AddForeignKey("core_sub_menu_id","core_modulos(id)","RESTRICT", "RESTRICT")
	//db.Model(&models.CoreSubMenu{}).AddForeignKey("core_menu_id","core_menus(id)","RESTRICT", "RESTRICT")
	//db.Model(&models.CoreMenu{}).AddForeignKey("core_rol_id","core_roles(id)","RESTRICT", "RESTRICT")
	//db.Model(&models.CoreModulo{}).AddForeignKey("core_rol_id","core_roles(id)","RESTRICT", "RESTRICT")
	//
	//db.Model(&models.Usuario{}).AddForeignKey("core_rol_id","core_roles(id)","RESTRICT", "RESTRICT")
}
