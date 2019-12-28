package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

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

//ExtractToken for split token
func ExtractToken(token string) string {
	if len(strings.Split(token, " ")) == 2 {
		token = strings.Split(token, " ")[1]
		return token
	}
	return ""
}

//VerifyToken middleware
func VerifyToken(c *gin.Context) error {
	header := c.Request.Header.Get("Authorization")
	tokenString := ExtractToken(header)
	token, err := jwt.Parse(tokenString, func(tkn *jwt.Token) (interface{}, error) {
		if _, ok := tkn.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Token no valido: %v", tkn.Header["alg"])
		}
		return []byte("MySecretSeed"), nil
	})
	if err != nil {
		return err
	}
	fmt.Println("Result", token)
	return nil
}
