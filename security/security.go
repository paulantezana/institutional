package security

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/paulantezana/institutional/models"
	"log"
	"time"
    "net/http"
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
	result, err := token.SignedString([]byte("secret"))
	if err != nil {
		log.Fatal("No se pudo firmar el token")
	}
	return result
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
    (*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3005")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

