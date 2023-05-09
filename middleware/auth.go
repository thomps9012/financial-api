package middleware

import (
	"financial-api/config"
	"financial-api/methods"
	"strings"

	res "financial-api/responses"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func Protected() func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey:   methods.Normalize(config.ENV("JWT_SECRET")),
		ErrorHandler: jwtError,
		ContextKey:   "finance_requests",
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusUnauthorized).JSON(res.BadJWT())
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(res.KeyNotFound())
	}
}

func ParseToken(token_string string) (map[string]interface{}, error) {
	token, err := jwt.Parse(token_string, func(token *jwt.Token) (interface{}, error) {
		return methods.Normalize(config.ENV("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		info := map[string]interface{}{
			"user_id":     claims["user_id"].(string),
			"admin":       claims["admin"].(bool),
			"permissions": claims["permissions"].([]interface{}),
		}
		return info, nil
	}
	return nil, err
}

func AdminRoute(c *fiber.Ctx) error {
	token_string := strings.TrimSpace(strings.Split(c.GetReqHeaders()["Authorization"], "Bearer")[1])
	token_info, err := ParseToken(token_string)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(res.BadJWT())
	}
	if token_info["admin"] != true {
		return c.Status(fiber.StatusForbidden).JSON(res.NotAdmin())
	}
	return c.Next()
}

func TokenID(c *fiber.Ctx) (string, error) {
	token_string := strings.TrimSpace(strings.Split(c.GetReqHeaders()["Authorization"], "Bearer")[1])
	token_info, err := ParseToken(token_string)
	if err != nil {
		return "", err
	}
	return token_info["user_id"].(string), nil
}
func TokenAdmin(c *fiber.Ctx) (bool, error) {
	token_string := strings.TrimSpace(strings.Split(c.GetReqHeaders()["Authorization"], "Bearer")[1])
	token_info, err := ParseToken(token_string)
	if err != nil {
		return false, err
	}
	return token_info["admin"].(bool), nil
}
