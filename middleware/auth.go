package middleware

import (
	"financial-api/config"
	"financial-api/methods"
	"fmt"
	"strconv"

	res "financial-api/responses"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

type Permission string

const (
	EMPLOYEE     Permission = "EMPLOYEE"
	MANAGER      Permission = "MANAGER"
	SUPERVISOR   Permission = "SUPERVISOR"
	EXECUTIVE    Permission = "EXECUTIVE"
	FINANCE_TEAM Permission = "FINANCE_TEAM"
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

func AdminRoute(c *fiber.Ctx) error {
	admin_cookie := c.Cookies("admin")
	fmt.Println(admin_cookie)
	admin_status, err := strconv.ParseBool(admin_cookie)
	if err != nil || !admin_status {
		return c.Status(fiber.StatusForbidden).JSON(res.NotAdmin())
	}
	return c.Next()
}
