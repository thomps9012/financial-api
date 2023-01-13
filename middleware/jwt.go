package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

type Claims struct {
	id          string
	name        string
	admin       bool
	permissions []Permission
	jwt.StandardClaims
}

func GenerateToken(id string, name string, admin bool, permissions []Permission) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["name"] = name
	claims["admin"] = admin
	claims["permissions"] = permissions
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
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
			"id":          claims["id"].(string),
			"name":        claims["name"].(string),
			"admin":       claims["admin"].(bool),
			"permissions": claims["permissions"].([]interface{}),
		}
		return info, nil
	}
	return nil, fmt.Errorf("Invalid Token")
}
