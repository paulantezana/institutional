package api

import (
    "html/template"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
    "github.com/paulantezana/institutional/controller"
    "github.com/paulantezana/institutional/security"
    "io"
    "net/http"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
    templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

    // Add global methods if data is a map
    if viewContext, isMap := data.(map[string]interface{}); isMap {
        viewContext["reverse"] = c.Echo().Reverse
    }

    return t.templates.ExecuteTemplate(w, name, data)
}




// Start is starting api web service
func Start() {
	e := echo.New()

    renderer := &TemplateRenderer{
        templates: template.Must(template.ParseGlob("../public/*.html")),
    }
    e.Renderer = renderer

    // Named route "foobar"
    e.GET("/something", func(c echo.Context) error {
        return c.Render(http.StatusOK, "../something.html", map[string]interface{}{
            "name": "Dolly!",
        })
    }).Name = "foobar"

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

	// Middleware
	//e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	//	AllowOrigins: []string{"http://localhost:7070", "http://localhost:3005"},
	//	AllowMethods: []string{echo.GET, echo.POST},
	//}))
    e.POST("/login", controller.Login)

	// Page template
	e.Static("/","../public")

    // Restricted group
    api := e.Group("/api")

    // Configure middleware with the custom claims type
    config := middleware.JWTConfig{
    	Claims:     &security.Claim{},
    	SigningKey: []byte("secret"),
    }
    api.Use(middleware.JWTWithConfig(config))

    // Routes =================================================== Usuario ===============================
    api.GET("/usuario", controller.GetUsuario)
    e.POST("/usuario", controller.CreateUsuario)
    e.PUT("/usuario", controller.UpdateUsuario)
    e.DELETE("/usuario", controller.DeleteUsuario)

	// Routes =================================================== Carrera ===============================
	e.GET("/carrera",controller.GetCarreras)
    e.POST("/carrera", controller.CreateCarrera)
    e.PUT("/carrera", controller.UpdateCarrera)
    e.DELETE("/carrera", controller.DeleteCarrera)

	// Init server
	e.Logger.Fatal(e.Start(":7070"))
}
