package responses

import (
	"financial-api/methods"
	"financial-api/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreatePettyCash(data models.Petty_Cash_Overview) fiber.Map {
	return fiber.Map{
		"status":  "CREATED",
		"code":    201,
		"message": "Petty Cash Request Successfully created @ " + methods.TimeNowFormat(),
		"data":    data,
	}
}
func MonthlyPettyCash(month int, year int, data []models.Petty_Cash_Overview) fiber.Map {
	return fiber.Map{
		"status":  "OK",
		"code":    200,
		"message": "Monthly Petty Cash Report for " + time.Month(month).String() + ", " + strconv.FormatInt(int64(year), 64),
		"data":    data,
	}
}
func PettyCashDetail(data models.Petty_Cash_Request) fiber.Map {
	return fiber.Map{
		"status":  "OK",
		"code":    200,
		"message": "Petty Cash Request with ID: " + data.ID + " Found",
		"data":    data,
	}
}
func EditPettyCash(data models.Petty_Cash_Overview) fiber.Map {
	return fiber.Map{
		"status":  "OK",
		"code":    200,
		"message": "Petty Cash Request with " + data.ID + " has been EDITED",
		"data":    data,
	}
}
func DeletePettyCash(data models.Petty_Cash_Overview) fiber.Map {
	return fiber.Map{
		"status":  "OK",
		"code":    200,
		"message": "Petty Cash Request with " + data.ID + " has been DELETED",
		"data":    data,
	}
}
func ApprovePettyCash(data models.Petty_Cash_Overview) fiber.Map {
	return fiber.Map{
		"status":  "OK",
		"code":    200,
		"message": "Petty Cash Request with " + data.ID + " has been APPROVED",
		"data":    data,
	}
}
func RejectPettyCash(data models.Petty_Cash_Overview) fiber.Map {
	return fiber.Map{
		"status":  "OK",
		"code":    200,
		"message": "Petty Cash Request with " + data.ID + " has been REJECTED",
		"data":    data,
	}
}
