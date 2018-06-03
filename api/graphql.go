package api

import (
    "github.com/labstack/echo"
    "bytes"
    "github.com/graphql-go/graphql"
    "errors"
    "net/http"
    "log"
)



func GraphQL(c echo.Context) error {


    //httpHandler := handler.New(&handler.Config{
    //    Schema: Schema,
    //    Pretty: true,
    //})

    // Validation token security
    //user := c.Get("user").(*jwt.Token)
    //claims:= user.Claims.(*security.Claim)

    // Request data
    bufBody := new(bytes.Buffer)
    bufBody.ReadFrom(c.Request().Body)
    query := bufBody.String()
    log.Printf(query)

    // Query GraphQL
    result := graphql.Do(graphql.Params{
        Schema:        Schema,
        RequestString: query,
    })
    if len(result.Errors) > 0 {
        return errors.New("GraphQL ERROR")
    }

    // Return result
    return c.JSON(http.StatusOK, result)
}