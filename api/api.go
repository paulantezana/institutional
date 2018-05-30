package api

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/paulantezana/institutional/security"
)

// Start is starting api web service
func Start() {
	e := echo.New()

	// Middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:7070", "http://localhost:3005"},
		AllowMethods: []string{echo.GET, echo.POST},
	}))
	// User routes
	e.POST("/login", security.Login)

	// Group routes protected JWT
	gql := e.Group("/graphql")
	config := middleware.JWTConfig{
		Claims:     &security.Claim{},
		SigningKey: []byte("secret"),
	}
	gql.Use(middleware.JWTWithConfig(config))

	// End point query
	gql.POST("", GraphQL)

	// Init server
	e.Logger.Fatal(e.Start(":7070"))
}
