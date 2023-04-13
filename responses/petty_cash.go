package responses

import (
	"financial-api/methods"
	"financial-api/models"
	"strconv"
	"time"
)

type PettyCashRes struct {
	Status  string                    `json:"status"`
	Code    int                       `json:"code"`
	Message string                    `json:"message"`
	Data    models.Petty_Cash_Request `json:"data"`
}
type PettyCashRequestsRes struct {
	Status  string                      `json:"status"`
	Code    int                         `json:"code"`
	Message string                      `json:"message"`
	Data    []models.Petty_Cash_Request `json:"data"`
}
type PettyCashOverviewRes struct {
	Status  string                     `json:"status"`
	Code    int                        `json:"code"`
	Message string                     `json:"message"`
	Data    models.Petty_Cash_Overview `json:"data"`
}
type PettyCashOverviewsRes struct {
	Status  string                       `json:"status"`
	Code    int                          `json:"code"`
	Message string                       `json:"message"`
	Data    []models.Petty_Cash_Overview `json:"data"`
}

func CreatePettyCash(data models.Petty_Cash_Overview) PettyCashOverviewRes {
	return PettyCashOverviewRes{
		Status:  "CREATED",
		Code:    201,
		Message: "Petty Cash Request Successfully created @ " + methods.TimeNowFormat(),
		Data:    data,
	}
}
func MonthlyPettyCash(month int, year int, data []models.Petty_Cash_Overview) PettyCashOverviewsRes {
	return PettyCashOverviewsRes{
		Status:  "OK",
		Code:    200,
		Message: "Monthly Petty Cash Report for " + time.Month(month).String() + ", " + strconv.Itoa(year),
		Data:    data,
	}
}
func PettyCashDetail(data models.Petty_Cash_Request) PettyCashRes {
	return PettyCashRes{
		Status:  "OK",
		Code:    200,
		Message: "Petty Cash Request with ID: " + data.ID + " Found",
		Data:    data,
	}
}
func EditPettyCash(data models.Petty_Cash_Overview) PettyCashOverviewRes {
	return PettyCashOverviewRes{
		Status:  "OK",
		Code:    200,
		Message: "Petty Cash Request with " + data.ID + " has been EDITED",
		Data:    data,
	}
}
func DeletePettyCash(data models.Petty_Cash_Overview) PettyCashOverviewRes {
	return PettyCashOverviewRes{
		Status:  "OK",
		Code:    200,
		Message: "Petty Cash Request with " + data.ID + " has been DELETED",
		Data:    data,
	}
}
func ApprovePettyCash(data models.Petty_Cash_Overview) PettyCashOverviewRes {
	return PettyCashOverviewRes{
		Status:  "OK",
		Code:    200,
		Message: "Petty Cash Request with " + data.ID + " has been APPROVED",
		Data:    data,
	}
}
func RejectPettyCash(data models.Petty_Cash_Overview) PettyCashOverviewRes {
	return PettyCashOverviewRes{
		Status:  "OK",
		Code:    200,
		Message: "Petty Cash Request with " + data.ID + " has been REJECTED",
		Data:    data,
	}
}
