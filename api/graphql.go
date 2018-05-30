package api

import (
	"bytes"
	"errors"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/labstack/echo"
)

// GraphQL function start server graphql
func GraphQL(c echo.Context) error {
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
