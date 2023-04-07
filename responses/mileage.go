package responses

import (
	"financial-api/methods"
	"financial-api/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateMileage(mileage_info models.Mileage_Overview) fiber.Map {
	return fiber.Map{
		"status":  "CREATED",
		"code":    201,
		"message": "Mileage Request Successfully created @ " + methods.TimeNowFormat(),
		"data":    mileage_info,
	}
}
func MonthlyMileage(month int, year int, mileage_info []models.Mileage_Overview) fiber.Map {
	return fiber.Map{
		"status":  "OK",
		"code":    200,
		"message": "Monthly Mileage Report for " + time.Month(month).String() + ", " + strconv.Itoa(year),
		"data":    mileage_info,
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
func MileageDetail(mileage_info models.Mileage_Request) fiber.Map {
	return fiber.Map{
		"status":  "OK",
		"code":    200,
		"message": "Mileage Request with " + mileage_info.ID + " Found",
		"data":    mileage_info,
	}
}
func EditMileage(mileage_info models.Mileage_Overview) fiber.Map {
	return fiber.Map{
		"status":  "OK",
		"code":    200,
		"message": "Mileage Request " + mileage_info.ID + " has been EDITED",
		"data":    mileage_info,
	}
}
func DeleteMileage(mileage_info models.Mileage_Overview) fiber.Map {
	return fiber.Map{
		"status":  "OK",
		"code":    200,
		"message": "Mileage Request " + mileage_info.ID + " has been DELETED",
		"data":    mileage_info,
	}
}
func ApproveMileage(mileage_info models.Mileage_Overview) fiber.Map {
	return fiber.Map{
		"status":  "OK",
		"code":    200,
		"message": "Mileage Request " + mileage_info.ID + " has been APPROVED",
		"data":    mileage_info,
	}
}
func RejectMileage(mileage_info models.Mileage_Overview) fiber.Map {
	return fiber.Map{
		"status":  "OK",
		"code":    200,
		"message": "Mileage Request " + mileage_info.ID + " has been REJECTED",
		"data":    mileage_info,
	}
}
