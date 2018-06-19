package security

import (
    "net/http"
    "github.com/dgrijalva/jwt-go"
    "fmt"
    "log"
    "github.com/paulantezana/institutional/config"
)

func HandleValidate(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        //setupResponse(&w, r) // Restriction coors
        tokenString := r.Header.Get("Authorization")[7:]

        token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
            }
            return []byte(config.GetConfig().Server.Key), nil
        })

        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            log.Printf("JWT Authenticated OK (app: %s)", claims["app"])

            next.ServeHTTP(w, r)
        }
    })
}