package security

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/paulantezana/institutional/config"
	"github.com/paulantezana/institutional/models"
    "github.com/paulantezana/institutional/helpers"
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
		user.ClaveAntigua =""
		user.ClaveRecuperar =""
		token := GenerateJWT(user)
		j, err := json.Marshal(helpers.Response{
            Success: true,
		    Data: Token{Token:token},
        })
		if err != nil {
			log.Fatalf("Error al convertir el token a json: %s", err)
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	} else {
	    j, err := json.Marshal(helpers.Response{
	        Success: false,
	        Errors: []string{"El nombre de usuario o contraseña es incorecta"},
        })
		if err != nil {
			log.Fatalf("Error al convertir el token a json: %s", err)
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}
