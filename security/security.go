package security

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/paulantezana/institutional/models/institucional"
	"log"
	"time"
)

// GenerateJWT generate token custom claims
func GenerateJWT(user institucional.Usuario) string {
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
	result, err := token.SignedString([]byte("secret"))
	if err != nil {
		log.Fatal("No se pudo firmar el token")
	}
	return result
}
