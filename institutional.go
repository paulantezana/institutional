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
)

func main() {
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

    // Config server
    server := &http.Server{
        Addr: ":7070",
        Handler: mux,
        ReadTimeout: 10 * time.Second,
        WriteTimeout: 10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }

    // listening server
    fmt.Println("listening... http://localhost:7070")
    log.Fatal(server.ListenAndServe())
}