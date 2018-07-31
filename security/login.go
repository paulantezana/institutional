package security

import (
	"crypto/sha256"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/helpers"
	"github.com/paulantezana/institutional/models"
)

// Login es el controlador de login
func Login(c echo.Context) error {
	user := models.Usuario{}
	if err := c.Bind(&user); err != nil {
		return err
	}

	db := config.GetConnection()
	defer db.Close()

	cc := sha256.Sum256([]byte(user.Clave))
	pwd := fmt.Sprintf("%x", cc)

	// Validate user and email
	db.Where("usuario = ? and clave = ?", user.Usuario, pwd).First(&user)
	if user.ID == 0 {
		db.Where("correo = ? and clave = ?", user.Usuario, pwd).First(&user)
	}

	// Response data login
	if user.ID == 0 {
		return c.JSON(http.StatusOK, helpers.Response{
			Success: false,
			Errors:  []string{"El nombre de usuario o contraseña es incorecta"},
		})
	}

	// customize send data
	user.Clave = ""
	user.ClaveAntigua = ""
	user.ClaveRecuperar = ""

	// get token key
	token := GenerateJWT(user)

	// Login success
	return c.JSON(http.StatusOK, helpers.Response{
		Success: true,
		Data:    Token{Token: token},
	})
}
