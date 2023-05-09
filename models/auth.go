package models

import (
	"financial-api/config"
	"financial-api/methods"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(user_id string, user_name string, user_permissions []string, user_admin bool) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user_id
	claims["admin"] = user_admin
	claims["permissions"] = user_permissions
	claims["exp"] = time.Now().Add(time.Hour * 6).Unix()
	claims["aud"] = user_name
	claims["issuer"] = config.ENV("JWT_ISS")
	jwt, err := token.SignedString(methods.Normalize(config.ENV("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return jwt, nil
}
