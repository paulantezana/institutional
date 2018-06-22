package main

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/labstack/gommon/log"
	"github.com/mnmtanish/go-graphiql"
	"github.com/paulantezana/institutional/migrations"
	"github.com/paulantezana/institutional/mutations"
	"github.com/paulantezana/institutional/queries"
	"net/http"
	"os"
    "github.com/paulantezana/institutional/security"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
    "flag"
    "github.com/paulantezana/institutional/config"
)

func main() {
	// Migration database
    var migrate string
    flag.StringVar(&migrate, "migrate", "no", "Genera la migración a la BD")
    flag.Parse()
    if migrate == "yes" {
        fmt.Println("Init migration")
        migrations.Migrate()
        fmt.Println("Finalize migration")
    }

	// root mutation
	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name:   "RootMutation",
		Fields: mutations.GetRootFields(),
	})

	//root query
	//we just define a trivial example here, since root query is required.
	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name:   "RootQuery",
		Fields: queries.GetRootFields(),
	})

	// define schema
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})

	if err != nil {
		panic(err)
	}

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	// Create new server mux
	router := mux.NewRouter()

	//Static file server
    router.PathPrefix("/").Handler(http.FileServer(http.Dir("public")))

	// Handle functions
	router.HandleFunc("/login",security.Login).Methods("POST")
	router.Handle("/graphql",security.HandleValidate(h))
	router.HandleFunc("/graphiql", graphiql.ServeGraphiQL)

	// Custom port
	port := os.Getenv("PORT")
	if port == "" {
		port = config.GetConfig().Server.Port
	}

	// listening server
    fmt.Println("======================================================")
	fmt.Println("Software developer paul antezana - architecture frontend")
    fmt.Println("=> http server started on http://localhost:" + port)
    fmt.Println("======================================================")

	// Config coors server listener
	log.Fatal(http.ListenAndServe(":" + port, handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST"}),
		handlers.AllowedOrigins([]string{"*"}))(router),
	))
}
