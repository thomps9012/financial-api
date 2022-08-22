package middleware

import (
	"fmt"
	"log"
	// "os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// var jwtKey = []byte(os.Getenv("JWT_KEY"))
var jwtKey = []byte("15cbe07571a61d35917e62a42f5e76c7")

type Claims struct {
	id   string
	role string
	jwt.StandardClaims
}

func GenerateToken(id string, role string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Fatal("Error while signing JWT")
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var info = map[string]interface{}{
			"id":   claims["id"].(string),
			"role": claims["role"].(string),
		}
		return info, nil
	}
	return nil, fmt.Errorf("Invalid Token")
}
