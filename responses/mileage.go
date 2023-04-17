package responses

import (
	"financial-api/methods"
	"financial-api/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type MileageOverviewRes struct {
	Message string                  `json:"message"`
	Data    models.Mileage_Overview `json:"data"`
}
type MileageOverviewsRes struct {
	Message string                    `json:"message"`
	Data    []models.Mileage_Overview `json:"data"`
}
type MileagesRes struct {
	Message string                   `json:"message"`
	Data    []models.Mileage_Request `json:"data"`
}
type MileageRes struct {
	Message string                 `json:"message"`
	Data    models.Mileage_Request `json:"data"`
}

func CreateMileage(mileage_info models.Mileage_Overview) MileageOverviewRes {
	return MileageOverviewRes{
		Message: "Mileage Request Successfully created @ " + methods.TimeNowFormat(),
		Data:    mileage_info,
	}
}
func MileageDetail(mileage_info models.Mileage_Request) MileageRes {
	return MileageRes{
		Message: "Mileage Request with " + mileage_info.ID + " Found",
		Data:    mileage_info,
	}
}
func EditMileage(mileage_info models.Mileage_Overview) MileageOverviewRes {
	return MileageOverviewRes{
		Message: "Mileage Request " + mileage_info.ID + " has been EDITED",
		Data:    mileage_info,
	}
}
func DeleteMileage(mileage_info models.Mileage_Overview) MileageOverviewRes {
	return MileageOverviewRes{
		Message: "Mileage Request " + mileage_info.ID + " has been DELETED",
		Data:    mileage_info,
	}
}
func ApproveMileage(mileage_info models.Mileage_Overview) MileageOverviewRes {
	return MileageOverviewRes{
		Message: "Mileage Request " + mileage_info.ID + " has been APPROVED",
		Data:    mileage_info,
	}
}
func RejectMileage(mileage_info models.Mileage_Overview) MileageOverviewRes {
	return MileageOverviewRes{
		Message: "Mileage Request " + mileage_info.ID + " has been REJECTED",
		Data:    mileage_info,
	}
}
func MonthlyMileage(month int, year int, mileage_info []models.Mileage_Overview) MileageOverviewsRes {
	return MileageOverviewsRes{
		Message: "Monthly Mileage Report for " + time.Month(month).String() + ", " + strconv.Itoa(year),
		Data:    mileage_info,
	}
}
func MileageVariance() fiber.Map {
	return fiber.Map{
		"status":  "NOT IMPLEMENTED",
		"code":    501,
		"message": "This API Endpoint is still in development",
		"data":    nil,
	}
}
