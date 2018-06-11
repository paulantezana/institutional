package main

import (
    "github.com/graphql-go/graphql"
    "github.com/paulantezana/institutional/mutations"
    "github.com/paulantezana/institutional/queries"
    "github.com/graphql-go/handler"
    "net/http"
    "fmt"
    "time"
    "github.com/labstack/gommon/log"
    "github.com/mnmtanish/go-graphiql"
    "os"
)

func main() {
    // Migration database
    //fmt.Println("Init migration")
    //migrations.Migrate()
    //fmt.Println("Finalize migration")

    // root mutation
    rootMutation := graphql.NewObject(graphql.ObjectConfig{
        Name: "RootMutation",
        Fields: mutations.GetRootFields(),
    })

    //root query
    //we just define a trivial example here, since root query is required.
    rootQuery := graphql.NewObject(graphql.ObjectConfig{
        Name: "RootQuery",
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
        Schema: &schema,
        Pretty: true,
        GraphiQL: true,
    })

    // Create new server mux
    mux := http.NewServeMux()

    // Static file server
    fs := http.FileServer(http.Dir("public"))
    mux.Handle("/",fs)

    // Handle functions
    mux.Handle("/graphql", h)
    mux.HandleFunc("/graphiql",graphiql.ServeGraphiQL)


    // Custom port
    port := os.Getenv("PORT")
    if port == "" {
        port = "7070"
    }
    // Config server
    server := &http.Server{
        Addr: ":" + port,
        Handler: mux,
        ReadTimeout: 10 * time.Second,
        WriteTimeout: 10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }

    // listening server
    fmt.Println("listening... http://localhost:7070")
    log.Fatal(server.ListenAndServe())
}