package security

import (
    "net/http"
    "encoding/json"
    "fmt"
    "crypto/sha256"
    "log"
    "github.com/paulantezana/institutional/models"
    "github.com/paulantezana/institutional/config"
)

// Login es el controlador de login
func Login(w http.ResponseWriter, r *http.Request) {
    user := models.Usuario{}
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        fmt.Fprintf(w, "Error: %s\n", err)
        return
    }

    db := config.GetConnection()
    defer db.Close()

    c := sha256.Sum256([]byte(user.Clave))
    pwd := fmt.Sprintf("%x", c)

    db.Where("usuario = ? and clave = ?", user.Usuario, pwd).First(&user)
    if user.ID > 0 {
        user.Clave = ""
        token := GenerateJWT(user)

        j, err := json.Marshal(Token{Token: token})
        if err != nil {
            log.Fatalf("Error al convertir el token a json: %s", err)
        }
        w.Header().Set("Content-Type", "application/json; charset=utf-8")
        w.WriteHeader(http.StatusOK)
        w.Write(j)
    } else {
        http.Error(w,"Usuario o contraseña incorrecta",http.StatusUnauthorized)
    }
}