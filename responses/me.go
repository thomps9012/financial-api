package responses

import (
	"financial-api/models"

	"github.com/gofiber/fiber/v2"
)

func MyInfo(my_info models.PublicInfo) fiber.Map {
	return fiber.Map{
		"status":  "OK",
		"code":    200,
		"message": "Found the following request and user info for " + my_info.Name,
		"data":    my_info,
	}
}
func MyMileage(requests []models.Mileage_Request) fiber.Map {
	return fiber.Map{
		"status":  "OK",
		"code":    200,
		"message": "Your mileage request information",
		"data":    requests,
	}
}
func MyCheckRequests(requests []models.Check_Request) fiber.Map {
	return fiber.Map{
		"status":  "OK",
		"code":    200,
		"message": "Your check request information",
		"data":    requests,
	}
}
func MyPettyCash(requests []models.Petty_Cash_Request) fiber.Map {
	return fiber.Map{
		"status":  "OK",
		"code":    200,
		"message": "You petty cash requests",
		"data":    requests,
	}
}
func AddVehicle(vehicle models.Vehicle) fiber.Map {
	return fiber.Map{
		"status":  "CREATED",
		"code":    201,
		"message": "New vehicle added to your account",
		"data":    vehicle,
	}
}
func EditVehicle(vehicle models.Vehicle) fiber.Map {
	return fiber.Map{
		"status":  "OK",
		"code":    200,
		"message": "Vehicle " + vehicle.ID + " successfully saved",
		"data":    vehicle,
	}
}
func RemoveVehicle(vehicle models.Vehicle) fiber.Map {
	return fiber.Map{
		"status":  "OK",
		"code":    200,
		"message": "Vehicle " + vehicle.Name + " removed from account",
		"data":    vehicle,
	}
}
