package security

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/paulantezana/institutional/models"
)

type Claim struct {
	Usuario models.Usuario `json:"usuario"`
	jwt.StandardClaims
}
