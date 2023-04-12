package responses

import (
	"financial-api/methods"
	"financial-api/models"

	"github.com/gofiber/fiber/v2"
)

func BadJWT() NilRes {
	return NilRes{
		Status:  "UNAUTHORIZED",
		Code:    fiber.StatusUnauthorized,
		Message: "Missing or incorrect JSON Web Token",
		Data:    "null",
	}
}

func BadUserID() NilRes {
	return NilRes{
		Status:  "BAD REQUEST",
		Code:    400,
		Message: "You're either not logged in or are using an invalid user id",
		Data:    "null",
	}
}

func KeyNotFound() NilRes {
	return NilRes{
		Status:  "UNAUTHORIZED",
		Code:    fiber.StatusUnauthorized,
		Message: "You're attempting to access the resource using an incorrect email, expired or invalid token, please check your credentials or contact app_support@norainc.org",
		Data:    "null",
	}
}

func NotAdmin() NilRes {
	return NilRes{
		Status:  "FORBIDDEN",
		Code:    fiber.StatusForbidden,
		Message: "You're attempting to visit an administrative resource without the proper permissions",
		Data:    "null",
	}
}

type MalformedBodyRes struct {
	Status  string                   `json:"status"`
	Code    int                      `json:"code"`
	Message string                   `json:"string"`
	Data    []*methods.ErrorResponse `json:"data"`
}

func MalformedBody(errors []*methods.ErrorResponse) MalformedBodyRes {
	return MalformedBodyRes{
		Status:  "BAD REQUEST",
		Code:    400,
		Message: "You're request body is invalid",
		Data:    errors,
	}
}

func MalformedRequest(code int, message string) NilRes {
	return NilRes{
		Status:  "BAD REQUEST",
		Code:    code,
		Message: message,
		Data:    "null",
	}
}

func ServerError(message string) NilRes {
	return NilRes{
		Status:  "INTERNAL SERVER ERROR",
		Code:    fiber.StatusInternalServerError,
		Message: message,
		Data:    "null",
	}
}

func InvalidEmail() NilRes {
	return NilRes{
		Status:  "FORBIDDEN",
		Code:    fiber.StatusForbidden,
		Message: "You're attempting to access a protected organization site",
		Data:    "null",
	}
}

type ErrorLogRes struct {
	Status  string                  `json:"status"`
	Code    int                     `json:"code"`
	Message string                  `json:"string"`
	Data    models.ErrorLogOverview `json:"data"`
}

func LogError(data models.ErrorLogOverview) ErrorLogRes {
	return ErrorLogRes{
		Status:  "INTERNAL SERVER ERROR",
		Code:    fiber.StatusInternalServerError,
		Message: "Your error has been logged in the system @ " + methods.TimeNowFormat(),
		Data:    data,
	}
}
