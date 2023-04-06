package responses

import (
	"financial-api/methods"

	"github.com/gofiber/fiber/v2"
)

func NewUser(token string) fiber.Map {
	return fiber.Map{
		"status":  "CREATED",
		"code":    201,
		"message": "New user successfully created @ " + methods.TimeNowFormat(),
		"data":    token,
	}
}

func LoggedIn(token string) fiber.Map {
	return fiber.Map{
		"status":  "OK",
		"code":    200,
		"message": "Successfully logged in @ " + methods.TimeNowFormat(),
		"data":    token,
	}
}

func LoggedOut() fiber.Map {
	return fiber.Map{
		"status":  "OK",
		"code":    200,
		"message": "Successfully logged out @ " + methods.TimeNowFormat(),
		"data":    nil,
	}
}
