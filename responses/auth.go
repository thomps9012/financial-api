package responses

import (
	"financial-api/methods"
)

type NilRes struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"string"`
	Data    string `json:"data"`
}

func NewUser(token string) NilRes {
	return NilRes{
		Status:  "CREATED",
		Code:    201,
		Message: "New user successfully created @ " + methods.TimeNowFormat(),
		Data:    token,
	}
}

func LoggedIn(token string) NilRes {
	return NilRes{
		Status:  "OK",
		Code:    200,
		Message: "Successfully logged in @ " + methods.TimeNowFormat(),
		Data:    token,
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
