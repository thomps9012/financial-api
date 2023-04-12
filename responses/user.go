package responses

import (
	"financial-api/models"

	"github.com/gofiber/fiber/v2"
)

func AllUsers(user_info []models.PublicInfo) fiber.Map {
	return fiber.Map{
		"status":  "OK",
		"code":    200,
		"message": "All user data",
		"data":    user_info,
	}
}
func OneUser(user_info models.PublicInfo) fiber.Map {
	return fiber.Map{
		"status":  "OK",
		"code":    200,
		"message": "Overview information for " + user_info.ID,
		"data":    user_info,
	}
}
func DeactivateUser(user_info models.PublicInfo) fiber.Map {
	return fiber.Map{
		"status":  "OK",
		"code":    200,
		"message": "User " + user_info.Name + "'s account has been deactivated",
		"data":    nil,
	}
}
func UserMileage(user_id string, user_info []models.Mileage_Overview) fiber.Map {
	return fiber.Map{
		"status":  "OK",
		"code":    200,
		"message": "Mileage Requests for User " + user_id,
		"data":    user_info,
	}
}
func UserPettyCash(user_id string, user_info []models.Petty_Cash_Overview) fiber.Map {
	return fiber.Map{
		"status":  "OK",
		"code":    200,
		"message": "Petty Cash Requests for User " + user_id,
		"data":    user_info,
	}
}
func UserCheckRequests(user_id string, user_info []models.Check_Request_Overview) fiber.Map {
	return fiber.Map{
		"status":  "OK",
		"code":    200,
		"message": "Check Requests for User " + user_id,
		"data":    user_info,
	}
}
