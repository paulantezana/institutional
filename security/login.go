package security

import (
	"crypto/sha256"
	"fmt"
	"github.com/labstack/echo"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models/institucional"
	"net/http"
)

// Login login user in api
func Login(c echo.Context) error {
	user := institucional.Usuario{}
	if err := c.Bind(&user); err != nil {
		return err
	}

	db := config.GetConnection()
	defer db.Close()

	cc := sha256.Sum256([]byte(user.Clave))
	pwd := fmt.Sprintf("%x", cc)

	db.Where("username=? and pass =?", user.Usuario, pwd).First(&user)
	if user.ID > 0 {
		user.Clave = ""
		token := GenerateJWT(user)

		return c.JSON(http.StatusOK, Token{
			Token: token,
		})
	} else {
		//return c.JSON(http.StatusOK, helpers.Message{
		//	Message: "Invalid username or password",
		//})
		return echo.ErrUnauthorized
	}
}
