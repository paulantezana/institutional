package controller

import (
	"crypto/sha256"
	"fmt"
	"github.com/labstack/echo"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/helpers"
	"github.com/paulantezana/institutional/security"
	"net/http"
    "github.com/paulantezana/institutional/models"
)

func Login(c echo.Context) error {
	// Get data json
	user := models.Usuario{}
	if err := c.Bind(&user); err != nil {
		return err
	}

	// Get connection database
	db := config.GetConnection()
	defer db.Close()

	// Execute operations
	cc := sha256.Sum256([]byte(user.Clave))
	pwd := fmt.Sprintf("%x", cc)

	db.Where("usuario=? and clave =?", user.Usuario, pwd).First(&user)
	if user.ID > 0 {
		user.Clave = ""
		token := security.GenerateJWT(user)

		return c.JSON(http.StatusOK, security.Token{
			Token: token,
		})
	} else {
		//return c.JSON(http.StatusOK, helpers.Message{
		//	Message: "Invalid username or password",
		//})
		return echo.ErrUnauthorized
	}
}

func GetUsuario(c echo.Context) error {
	// Get connection database
	db := config.GetConnection()
	defer db.Close()

	// Execute operations
	usuarios := make([]models.Usuario, 0)
	if err := db.Find(&usuarios).Error; err != nil {
		return err
	}

	// Return final data
	return c.JSON(http.StatusOK, usuarios)
}

func CreateUsuario(c echo.Context) error {
	// Get data json
	usuario := models.Usuario{}
	if err := c.Bind(&usuario); err != nil {
		return err
	}

	// Get connection database
	db := config.GetConnection()
	defer db.Close()

	// Execute operations
	if usuario.Clave != usuario.ConfirmarClave {
		return c.JSON(http.StatusOK, helpers.Message{
			Title:   "Login",
			Message: "Las contraseñas no coinciden",
		})
	}

	cc := sha256.Sum256([]byte(usuario.Clave))
	pwd := fmt.Sprintf("%x", cc)
	usuario.Clave = pwd

	if err := db.Create(&usuario).Error; err != nil {
		return err
	}

	// Return final data
	return c.JSON(http.StatusCreated, helpers.Message{
		ID:      usuario.ID,
		Title:   "Usuario",
		Message: "Usuario Guardado",
	})
}

func UpdateUsuario(c echo.Context) error {
	// Get data json
	usuario := models.Usuario{}
	if err := c.Bind(&usuario); err != nil {
		return err
	}

	// Get connection database
	db := config.GetConnection()
	defer db.Close()

	// Execute operations
	if err := db.Model(&usuario).Updates(usuario).Error; err != nil {
		return err
	}

	// Return final data
	return c.JSON(http.StatusOK, usuario)
}

func DeleteUsuario(c echo.Context) error {
	// Get data json
	usuario := models.Usuario{}
	if err := c.Bind(&usuario); err != nil {
		return err
	}

	// Get connection database
	db := config.GetConnection()
	defer db.Close()

	// Execute operations
	if err := db.Delete(&usuario).Error; err != nil {
		return err
	}

	// Return final data
	return c.JSON(http.StatusNoContent, usuario)
}
