package responses

import (
	"financial-api/methods"
	"financial-api/models"
)

type NilRes struct {
	Message string `json:"message"`
	Data    string `json:"data"`
}

type UserLoginRes struct {
	Message string          `json:"message"`
	Data    models.LoginRes `json:"data"`
}

func NewUser(login_data models.LoginRes) UserLoginRes {
	return UserLoginRes{
		Message: "New user successfully created @ " + methods.TimeNowFormat(),
		Data:    login_data,
	}
}

func LoggedIn(login_data models.LoginRes) UserLoginRes {
	return UserLoginRes{
		Message: "Successfully logged in @ " + methods.TimeNowFormat(),
		Data:    login_data,
	}
}

func LoggedOut() NilRes {
	return NilRes{
		Message: "Successfully logged out @ " + methods.TimeNowFormat(),
		Data:    "null",
	}
}
