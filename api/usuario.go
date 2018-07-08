package api

import (
    "net/http"
    "github.com/paulantezana/institutional/models"
    "encoding/json"
    "fmt"
    "github.com/paulantezana/institutional/config"
    "github.com/paulantezana/institutional/helpers"
    "log"
    "math/rand"
    "time"
    "bytes"
    "html/template"
    "crypto/sha256"
)

func ForgoutSearch(w http.ResponseWriter, r *http.Request) {
    xUser := models.Usuario{}
    user := models.Usuario{}
    err := json.NewDecoder(r.Body).Decode(&xUser)
    if err != nil {
        fmt.Fprintf(w, "Error: %s\n", err)
        return
    }
    // Create response
    res := helpers.Response{}

    // get connection
    db := config.GetConnection()
    defer db.Close()

    // Validations
    if err := db.Where("correo = ?", xUser.Correo).First(&user).Error; err != nil {
        res = helpers.Response{
            Errors: []string{"Tu búsqueda no arrojó ningún resultado. Vuelve a intentarlo con otros datos."},
            Success: false,
        }
    }else {
        //Generate key validation
        key := (int)(rand.Float32() * 10000000)
        user.ClaveRecuperar = fmt.Sprint(key)
        user.FechaRecuperacionClave = time.Now()

        //Update database
        if err := db.Model(&user).Update(user).Error; err != nil {
            return
        }

        // SEND EMAIL get html template
        t, err := template.ParseFiles("public/mail.html")
        if err != nil{
          log.Panic(err)
        }
        // SEND EMAIL new buffer
        buf := new(bytes.Buffer)
        err = t.Execute(buf, user)
        if err != nil{
          log.Panic(err)
        }

        // SEND EMAIL
        err = config.SendEmail(user.Correo, fmt.Sprint(key) + " es el código de recuperación de tu cuenta en SINST",buf.String())
        if err != nil {
          log.Panic(err)
        }

        // Response success api service
        user.ClaveRecuperar = ""
        user.Clave = ""
        res = helpers.Response{
            Success: true,
            Data: user,
            Errors: []string{},
        }
    }

    j, err := json.Marshal(res)
    if err != nil {
        log.Fatalf("Error al convertir el response a json: %s", err)
    }
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.WriteHeader(http.StatusOK)
    w.Write(j)
}

func ForgoutValidate(w http.ResponseWriter, r *http.Request) {
    xUser := models.Usuario{}
    user := models.Usuario{}
    err := json.NewDecoder(r.Body).Decode(&xUser)
    if err != nil {
        fmt.Fprintf(w, "Error: %s\n", err)
        return
    }
    // Create response
    res := helpers.Response{}

    // get connection
    db := config.GetConnection()
    defer db.Close()

    // Validations
    if err := db.Where("id = ? AND clave_recuperar = ?", xUser.ID, xUser.ClaveRecuperar).First(&user).Error; err != nil {
        res = helpers.Response{
            Errors: []string{"El número que ingresaste no coincide con tu código. Vuelve a intentarlo"},
            Success: false,
        }
    }else {
        user.ClaveRecuperar = ""
        user.Clave = ""

        // Validate expiration
        if time.Now().Year() != user.FechaRecuperacionClave.Year() || time.Now().Month() != user.FechaRecuperacionClave.Month() {
            res = helpers.Response{
                Success: false,
                Errors: []string{"La clave ya expiró"},
            }
        }else {
            if time.Now().Day() - user.FechaRecuperacionClave.Day() > 7 {
                res = helpers.Response{
                    Success: false,
                    Errors: []string{"La clave ya expiró"},
                }
            }else {
                // Response success api service
                res = helpers.Response{
                    Success: true,
                    Data: user,
                    Errors: []string{},
                }
            }
        }


    }

    j, err := json.Marshal(res)
    if err != nil {
        log.Fatalf("Error al convertir el response a json: %s", err)
    }
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.WriteHeader(http.StatusOK)
    w.Write(j)
}

func ForgoutChange(w http.ResponseWriter, r *http.Request) {
    xUser := models.Usuario{}
    err := json.NewDecoder(r.Body).Decode(&xUser)
    if err != nil {
        fmt.Fprintf(w, "Error: %s\n", err)
        return
    }

    // Create response
    user := models.Usuario{}
    res := helpers.Response{}

    // get connection
    db := config.GetConnection()
    defer db.Close()

    if err := db.Where("id = ?", xUser.ID).First(&user).Error; err != nil {
        fmt.Fprintf(w, "Error: usuario no encontrado %s\n", err)
        return
    }

    // Encrypted old password
    tt := sha256.Sum256([]byte(xUser.Clave))
    ttt := fmt.Sprintf("%x", tt)

    // Update
    user.Clave = ttt
    user.ClaveRecuperar = ""
    user.FechaModificacionClave = time.Now()
    if err := db.Model(&user).Update(user).Error; err != nil {
        fmt.Fprintf(w, "Error: no se pudo actualizar los datos %s\n", err)
        return
    }

    res = helpers.Response{
        Success: true,
        Data: user,
        Errors: []string{},
    }

    j, err := json.Marshal(res)
    if err != nil {
        log.Fatalf("Error al convertir el response a json: %s", err)
    }
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.WriteHeader(http.StatusOK)
    w.Write(j)
}
