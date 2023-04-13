package responses

import (
	"financial-api/methods"
	"financial-api/models"
)

type NilRes struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

type UserLoginRes struct {
	Status  string          `json:"status"`
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    models.LoginRes `json:"data"`
}

func NewUser(login_data models.LoginRes) UserLoginRes {
	return UserLoginRes{
		Status:  "CREATED",
		Code:    201,
		Message: "New user successfully created @ " + methods.TimeNowFormat(),
		Data:    login_data,
	}
}

func LoggedIn(login_data models.LoginRes) UserLoginRes {
	return UserLoginRes{
		Status:  "OK",
		Code:    200,
		Message: "Successfully logged in @ " + methods.TimeNowFormat(),
		Data:    login_data,
	}
}

func LoggedOut() NilRes {
	return NilRes{
		Status:  "OK",
		Code:    200,
		Message: "Successfully logged out @ " + methods.TimeNowFormat(),
		Data:    "null",
	}
}
