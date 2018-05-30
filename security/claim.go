package security

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/paulantezana/institutional/models/institucional"
)

type Claim struct {
	Usuario institucional.Usuario `json:"usuario"`
	jwt.StandardClaims
}
