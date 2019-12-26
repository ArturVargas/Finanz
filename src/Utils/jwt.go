package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

//CreateToken use jwt
func CreateToken(userID uint, userName, userEmail string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userID"] = userID
	claims["email"] = userEmail
	claims["name"] = userName
	claims["exp"] = time.Now().Add(time.Hour * 3).Unix() // El token expira en 3hrs
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("MySecretSeed")) // El Secret debe ir en las variables de entorno
}
