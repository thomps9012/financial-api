package responses

import (
	"financial-api/methods"

	"github.com/gofiber/fiber/v2"
)

func BadJWT() fiber.Map {
	return fiber.Map{
		"status":  "UNAUTHORIZED",
		"code":    fiber.StatusUnauthorized,
		"message": "Missing or incorrect JSON Web Token",
		"data":    nil,
	}
}

func BadUserID() fiber.Map {
	return fiber.Map{
		"status":  "BAD REQUEST",
		"code":    400,
		"message": "You're either not logged in or are using an invalid user id",
		"data":    nil,
	}
}

func KeyNotFound() fiber.Map {
	return fiber.Map{
		"status":  "UNAUTHORIZED",
		"code":    fiber.StatusUnauthorized,
		"message": "You're attempting to access the resource using an incorrect email, expired or invalid token, please check your credentials or contact app_support@norainc.org",
		"data":    nil,
	}
}

func NotAdmin() fiber.Map {
	return fiber.Map{
		"status":  "FORBIDDEN",
		"code":    fiber.StatusForbidden,
		"message": "You're attempting to visit an administrative resource without the proper permissions",
		"data":    nil,
	}
}

func MalformedBody(errors []*methods.ErrorResponse) fiber.Map {
	return fiber.Map{
		"status":  "BAD REQUEST",
		"code":    400,
		"message": "You're request body is invalid",
		"data":    errors,
	}
}

func MalformedRequest(code int, message string) fiber.Map {
	return fiber.Map{
		"status":  "BAD REQUEST",
		"code":    code,
		"message": message,
		"data":    nil,
	}
}

func ServerError(message string) fiber.Map {
	return fiber.Map{
		"status":  "INTERNAL SERVER ERROR",
		"code":    fiber.StatusInternalServerError,
		"message": message,
		"data":    nil,
	}
}

func InvalidEmail() fiber.Map {
	return fiber.Map{
		"status":  "FORBIDDEN",
		"code":    fiber.StatusForbidden,
		"message": "You're attempting to access a protected organization site",
		"data":    nil,
	}
}

func LogError() fiber.Map {
	return fiber.Map{
		"status":  "INTERNAL SERVER ERROR",
		"code":    fiber.StatusInternalServerError,
		"message": "Your error has been logged in the system @ " + methods.TimeNowFormat(),
		"data":    nil,
	}
}
