package responses

import (
	"financial-api/methods"
	"financial-api/models"
)

func BadJWT() NilRes {
	return NilRes{
		Message: "Missing or incorrect JSON Web Token",
		Data:    "null",
	}
}

func BadUserID() NilRes {
	return NilRes{
		Message: "You're either not logged in or are using an invalid user id",
		Data:    "null",
	}
}

func KeyNotFound() NilRes {
	return NilRes{
		Message: "You're attempting to access the resource using an incorrect email, expired or invalid token, please check your credentials or contact app_support@norainc.org",
		Data:    "null",
	}
}

func NotAdmin() NilRes {
	return NilRes{
		Message: "You're attempting to visit an administrative resource without the proper permissions",
		Data:    "null",
	}
}

func RequestExists(request_type string) NilRes {
	return NilRes{
		Message: "You're attempting to create a duplicate " + request_type + " request. Please edit an existing one or make the appropriate changes.",
		Data:    "null",
	}
}

type MalformedBodyRes struct {
	Message string                   `json:"string"`
	Data    []*methods.ErrorResponse `json:"data"`
}

func MalformedBody(errors []*methods.ErrorResponse) MalformedBodyRes {
	return MalformedBodyRes{
		Message: "You're request body is invalid",
		Data:    errors,
	}
}

func MalformedRequest(code int, message string) NilRes {
	return NilRes{
		Message: message,
		Data:    "null",
	}
}

func ServerError(message string) NilRes {
	return NilRes{
		Message: message,
		Data:    "null",
	}
}

func InvalidEmail() NilRes {
	return NilRes{
		Message: "You're attempting to access a protected organization site",
		Data:    "null",
	}
}

type ErrorLogRes struct {
	Message string                  `json:"string"`
	Data    models.ErrorLogOverview `json:"data"`
}

func LogError(data models.ErrorLogOverview) ErrorLogRes {
	return ErrorLogRes{
		Message: "Your error has been logged in the system @ " + methods.TimeNowFormat(),
		Data:    data,
	}
}
