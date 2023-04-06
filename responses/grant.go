package responses

import (
	"financial-api/models"

	"github.com/gofiber/fiber/v2"
)

func AllGrants(data []models.Grant) fiber.Map {
	return fiber.Map{
		"status":  "OK",
		"code":    200,
		"message": "All Grants Data",
		"data":    data,
	}
}
func OneGrant(data models.Grant) fiber.Map {
	return fiber.Map{
		"status":  "OK",
		"code":    200,
		"message": "Grant Data Found for Grant: " + data.Name,
		"data":    data,
	}
}
func GrantCheckRequests(grant models.Grant, data []models.Check_Request_Overview) fiber.Map {
	return fiber.Map{
		"status":  "OK",
		"code":    200,
		"message": "Check Requests for Grant " + grant.Name + " Found",
		"data":    data,
	}
}
func GrantMileage(grant models.Grant, data []models.Mileage_Overview) fiber.Map {
	return fiber.Map{
		"status":  "OK",
		"code":    200,
		"message": "Mileage Requests for Grant " + grant.Name + " Found",
		"data":    data,
	}
}
func GrantPettyCash(grant models.Grant, data []models.Petty_Cash_Overview) fiber.Map {
	return fiber.Map{
		"status":  "OK",
		"code":    200,
		"message": "Petty Cash Requests for Grant " + grant.Name + " Found",
		"data":    data,
	}
}
