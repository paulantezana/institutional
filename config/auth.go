package config
//
//import (
//    "crypto/rsa"
//    "io/ioutil"
//    "github.com/labstack/gommon/log"
//    "github.com/dgrijalva/jwt-go"
//    "github.com/paulantezana/institutional/models"
//    "time"
//)
//
//var (
//    privateKey *rsa.PrivateKey
//    PublicKey *rsa.PublicKey
//)
//
//
//func init()  {
//    privateBytes, err := ioutil.ReadFile("../keys/private.rsa")
//    if err != nil{
//        log.Fatal("No se pudo leer el archivo privado")
//    }
//
//    publicBytes, err := ioutil.ReadFile("../keys/public.rsa")
//    if err != nil{
//        log.Fatal("No se pudo leer el archivo público")
//    }
//
//    privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
//    if err != nil{
//        log.Fatal("No se pudo hacer el parse a privatekey")
//    }
//
//    PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
//    if err != nil{
//        log.Fatal("No se pudo hacer el parse a PublicKey")
//    }
//}
//
//// GenerateJWT Generate token client
//func GenerateJWT(user models.Usuario) string  {
//    // Set custom claims
//    claims := &Claim{
//        user,
//        jwt.StandardClaims{
//            ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
//            Issuer:    "paul",
//        },
//    }
//
//    // Create token with claims
//    token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
//
//    // Generate encoded token and send it as response.
//    result, err := token.SignedString(privateKey)
//    if err != nil {
//        log.Fatal("No se pudo firmar el token")
//    }
//    return result
//}