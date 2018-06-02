package api

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
    "github.com/paulantezana/institutional/controller"
    "github.com/paulantezana/institutional/security"
)


// Start is starting api web service
func Start() {
	e := echo.New()

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // Este midelware indica desde que 
    // ips o nombres de dominio se consumiran la api
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:7070", "http://localhost:3005", "http://localhost"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
    }))
    
    // Login del usuario de donde se optiene la clave token de seguridad
    // para despues hacer toda las peticiones a la api
    e.POST("/login", controller.Login)

    // API: con proteccion de autenticacion mediante de tokens de seguridad
    api := e.Group("/api")

    // Configuracion del middleware con los claims personalizados
    config := middleware.JWTConfig{
    	Claims:     &security.Claim{},
    	SigningKey: []byte("secret"),
    }
    api.Use(middleware.JWTWithConfig(config))

    // ==================================================================================================
    // Toda las rutas del api
    // Routes =================================================== Usuario ===============================
    api.GET("/usuario", controller.GetUsuario)
    api.POST("/usuario", controller.CreateUsuario)
    api.PUT("/usuario", controller.UpdateUsuario)
    api.DELETE("/usuario", controller.DeleteUsuario)

    // Routes =================================================== Carrera ===============================
	api.GET("/instituto",controller.GetCarreras)
    api.POST("/instituto", controller.CreateCarrera)
    api.PUT("/instituto", controller.UpdateCarrera)
    api.DELETE("/instituto", controller.DeleteCarrera)

    // Routes =================================================== Carrera ===============================
	api.GET("/filial",controller.GetCarreras)
    api.POST("/filial", controller.CreateCarrera)
    api.PUT("/filial", controller.UpdateCarrera)
    api.DELETE("/filial", controller.DeleteCarrera)

	// Routes =================================================== Carrera ===============================
	api.GET("/carrera",controller.GetCarreras)
    api.POST("/carrera", controller.CreateCarrera)
    api.PUT("/carrera", controller.UpdateCarrera)
    api.DELETE("/carrera", controller.DeleteCarrera)

    // Routes =================================================== Semestre ===============================
	api.GET("/semestre",controller.GetCarreras)
    api.POST("/semestre", controller.CreateCarrera)
    api.PUT("/semestre", controller.UpdateCarrera)
    api.DELETE("/semestre", controller.DeleteCarrera)

    // Routes =================================================== Modulo ===============================
	api.GET("/modulo",controller.GetCarreras)
    api.POST("/modulo", controller.CreateCarrera)
    api.PUT("/modulo", controller.UpdateCarrera)
    api.DELETE("/modulo", controller.DeleteCarrera)

    // Routes =================================================== Unidad ===============================
	api.GET("/unidad",controller.GetCarreras)
    api.POST("/unidad", controller.CreateCarrera)
    api.PUT("/unidad", controller.UpdateCarrera)
    api.DELETE("/unidad", controller.DeleteCarrera)

    // Init server
	e.Logger.Fatal(e.Start(":7070"))
}
