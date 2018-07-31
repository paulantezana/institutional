package security

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models"
	"log"
	"time"
)

type Claim struct {
	Usuario models.Usuario `json:"usuario"`
	jwt.StandardClaims
}

// Token send json
type Token struct {
	Token string `json:"token"`
}

// GenerateJWT generate token custom claims
func GenerateJWT(user models.Usuario) string {
	// Set custom claims
	claims := &Claim{
		user,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "paulantezana",
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
