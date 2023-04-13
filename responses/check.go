package responses

import (
	"financial-api/methods"
	"financial-api/models"
	"strconv"
	"time"
)

type CheckOverviewsRes struct {
	Status  string                          `json:"status"`
	Code    int                             `json:"code"`
	Message string                          `json:"message"`
	Data    []models.Check_Request_Overview `json:"data"`
}
type CheckOverviewRes struct {
	Status  string                        `json:"status"`
	Code    int                           `json:"code"`
	Message string                        `json:"message"`
	Data    models.Check_Request_Overview `json:"data"`
}
type CheckDetailRes struct {
	Status  string               `json:"status"`
	Code    int                  `json:"code"`
	Message string               `json:"message"`
	Data    models.Check_Request `json:"data"`
}
type CheckRequestsRes struct {
	Status  string                 `json:"status"`
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    []models.Check_Request `json:"data"`
}

func NewCheckRequest(data models.Check_Request_Overview) CheckOverviewRes {
	return CheckOverviewRes{
		Status:  "CREATED",
		Code:    201,
		Message: "Check Request Successfully created @ " + methods.TimeNowFormat(),
		Data:    data,
	}
}
func CheckRequestDetail(data models.Check_Request) CheckDetailRes {
	return CheckDetailRes{
		Status:  "OK",
		Code:    200,
		Message: "Check Request with ID: " + data.ID + " Found",
		Data:    data,
	}
}
func EditCheckRequest(data models.Check_Request_Overview) CheckOverviewRes {
	return CheckOverviewRes{
		Status:  "OK",
		Code:    200,
		Message: "Check Request with " + data.ID + " has been EDITED",
		Data:    data,
	}
}
func DeleteCheckRequest(data models.Check_Request_Overview) CheckOverviewRes {
	return CheckOverviewRes{
		Status:  "OK",
		Code:    200,
		Message: "Check Request with " + data.ID + " has been DELETED",
		Data:    data,
	}
}
func ApproveCheckRequest(data models.Check_Request_Overview) CheckOverviewRes {
	return CheckOverviewRes{
		Status:  "OK",
		Code:    200,
		Message: "Check Request with " + data.ID + " has been APPROVED",
		Data:    data,
	}
}
func RejectCheckRequest(data models.Check_Request_Overview) CheckOverviewRes {
	return CheckOverviewRes{
		Status:  "OK",
		Code:    200,
		Message: "Check Request with " + data.ID + " has been REJECTED",
		Data:    data,
	}
}
func MonthlyCheckRequests(month int, year int, data []models.Check_Request_Overview) CheckOverviewsRes {
	return CheckOverviewsRes{
		Status:  "OK",
		Code:    200,
		Message: "Monthly Check Request Report for " + time.Month(month).String() + ", " + strconv.Itoa(year),
		Data:    data,
	}
}
