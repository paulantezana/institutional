package security

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/paulantezana/institutional/models"
	"log"
	"time"
    "github.com/paulantezana/institutional/config"
)

// GenerateJWT generate token custom claims
func GenerateJWT(user models.Usuario) string {
	// Set custom claims
	claims := &Claim{
		user,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "paul",
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	result, err := token.SignedString([]byte(config.GetConfig().Server.Key))
	if err != nil {
		log.Fatal("No se pudo firmar el token")
	}
	return result
}
