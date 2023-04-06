package responses

import (
	"financial-api/methods"
	"financial-api/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func NewCheckRequest(data models.Check_Request_Overview) fiber.Map {
	return fiber.Map{
		"status":  "CREATED",
		"code":    201,
		"message": "Check Request Successfully created @ " + methods.TimeNowFormat(),
		"data":    data,
	}
}
func MonthlyCheckRequests(month int, year int, data []models.Check_Request_Overview) fiber.Map {
	return fiber.Map{
		"status":  "OK",
		"code":    200,
		"message": "Monthly Check Request Report for " + time.Month(month).String() + ", " + strconv.FormatInt(int64(year), 64),
		"data":    data,
	}
}
func CheckRequestDetail(data models.Check_Request) fiber.Map {
	return fiber.Map{
		"status":  "OK",
		"code":    200,
		"message": "Check Request with ID: " + data.ID + " Found",
		"data":    data,
	}
}
func EditCheckRequest(data models.Check_Request_Overview) fiber.Map {
	return fiber.Map{
		"status":  "OK",
		"code":    200,
		"message": "Check Request with " + data.ID + " has been EDITED",
		"data":    data,
	}
}
func DeleteCheckRequest(data models.Check_Request_Overview) fiber.Map {
	return fiber.Map{
		"status":  "OK",
		"code":    200,
		"message": "Check Request with " + data.ID + " has been DELETED",
		"data":    data,
	}
}
func ApproveCheckRequest(data models.Check_Request_Overview) fiber.Map {
	return fiber.Map{
		"status":  "OK",
		"code":    200,
		"message": "Check Request with " + data.ID + " has been APPROVED",
		"data":    data,
	}
}
func RejectCheckRequest(data models.Check_Request_Overview) fiber.Map {
	return fiber.Map{
		"status":  "OK",
		"code":    200,
		"message": "Check Request with " + data.ID + " has been REJECTED",
		"data":    data,
	}
}
