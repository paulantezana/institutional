package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/mnmtanish/go-graphiql"
	"net/http"
	"os"
    "github.com/paulantezana/institutional/security"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
    "flag"
    "github.com/paulantezana/institutional/config"
    "github.com/paulantezana/institutional/models"
    "github.com/paulantezana/institutional/schema"
    "github.com/paulantezana/institutional/api"
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
    router := mux.NewRouter()

    // Custom port
    port := os.Getenv("PORT")
    if port == "" {
        port = config.GetConfig().Server.Port
    }

    router.HandleFunc("/login",security.Login).Methods("POST")

    router.HandleFunc("/forgout/serach",api.ForgoutSearch).Methods("POST")
    router.HandleFunc("/forgout/validate",api.ForgoutValidate).Methods("POST")
    router.HandleFunc("/forgout/change",api.ForgoutChange).Methods("POST")

    router.Handle("/graphql", schema.GraphQL())            // GraphQL Server

    router.HandleFunc("/graphiql", graphiql.ServeGraphiQL) // GraphiQL Server
    router.PathPrefix("/").Handler(http.FileServer(http.Dir("public"))) //Static file server

    fmt.Println("=> http server started on http://localhost:" + port) // listening server message

    // Config coors server listener
    log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(
        handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
        handlers.AllowedMethods([]string{"GET", "POST"}),
        handlers.AllowedOrigins([]string{"*"}))(router),
    ))
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
