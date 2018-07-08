package api

import (
    "net/http"
    "github.com/paulantezana/institutional/models"
    "fmt"
    "github.com/paulantezana/institutional/config"
    "github.com/paulantezana/institutional/helpers"
    "log"
    "math/rand"
    "time"
    "bytes"
    "html/template"
    "crypto/sha256"
    "github.com/labstack/echo"
    "crypto/md5"
)

func ForgoutSearch(c echo.Context) error {
    user := models.Usuario{}
    if err := c.Bind(&user); err != nil {
        return err
    }

    // Get connection
    db := config.GetConnection()
    defer db.Close()

    // Validations
    if err := db.Where("correo = ?", user.Correo).First(&user).Error; err != nil {
        return c.JSON(http.StatusOK,helpers.Response{
            Errors: []string{"Tu búsqueda no arrojó ningún resultado. Vuelve a intentarlo con otros datos."},
            Success: false,
        })
    }

    // Generate key validation
    key := (int)(rand.Float32() * 10000000)
    user.ClaveRecuperar = fmt.Sprint(key)
    user.FechaRecuperacionClave = time.Now()

    // Update database
    if err := db.Model(&user).Update(user).Error; err != nil {
        return c.JSON(http.StatusInternalServerError,err)
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
    return c.JSON(http.StatusOK,helpers.Response{
        Success: true,
        Data: user,
        Errors: []string{},
    })
}

func ForgoutValidate(c echo.Context) error {
    user := models.Usuario{}
    if err := c.Bind(&user); err != nil {
        return err
    }

    // get connection
    db := config.GetConnection()
    defer db.Close()

    // Validations
    if err := db.Where("id = ? AND clave_recuperar = ?", user.ID, user.ClaveRecuperar).First(&user).Error; err != nil {
        return c.JSON(http.StatusOK,helpers.Response{
            Errors: []string{"El número que ingresaste no coincide con tu código. Vuelve a intentarlo"},
            Success: false,
        })
    }

    // Validate expiration
    if
        time.Now().Year() != user.FechaRecuperacionClave.Year() ||
        time.Now().Month() != user.FechaRecuperacionClave.Month() ||
        time.Now().Day() - user.FechaRecuperacionClave.Day() > 7{
            return c.JSON(http.StatusOK,helpers.Response{
                Success: false,
                Errors: []string{"La clave ya expiró"},
            })
    }

    user.ClaveRecuperar = ""
    user.Clave = ""

    return c.JSON(http.StatusOK,helpers.Response{
        Success: true,
        Data: user,
        Errors: []string{},
    })
}

func ForgoutChange(c echo.Context) error {
    xUser := models.Usuario{}
    user := models.Usuario{}
    if err := c.Bind(&xUser); err != nil {
        return err
    }

    // get connection
    db := config.GetConnection()
    defer db.Close()

    if err := db.Where("id = ?", xUser.ID).First(&user).Error; err != nil {
        return err
    }

    // Encrypted old password
    tt := sha256.Sum256([]byte(xUser.Clave))
    ttt := fmt.Sprintf("%x", tt)

    // Update
    user.Clave = ttt
    user.ClaveRecuperar = ""
    user.FechaModificacionClave = time.Now()
    if err := db.Model(&user).Update(user).Error; err != nil {
        return err
    }

    return c.JSON(http.StatusOK,helpers.Response{
        Success: true,
        Data: user,
    })
}

func RegisterUser(c echo.Context) error {
    user := models.Usuario{}
    if err := c.Bind(&user); err != nil {
        return err
    }

    // get connection
    db := config.GetConnection()
    defer db.Close()

    if user.Clave != user.ConfirmarClave {
        c.JSON(http.StatusOK, helpers.Response{
            Success: false,
            Errors: []string{"Las contraseñas no coinciden"},
        })
    }

    cc := sha256.Sum256([]byte(user.Clave))
    pwd := fmt.Sprintf("%x", cc)
    user.Clave = pwd

    picmd5 := md5.Sum([]byte(user.Correo))
    picstr := fmt.Sprintf("%x", picmd5)
    user.Avatar = "https://gravatar.com/avatar/" + picstr + "?s=100"

    if err := db.Create(&user).Error; err != nil {
        return  err
    }

    return c.JSON(http.StatusOK,helpers.Response{
        Success: true,
        Data: user,
    })
}