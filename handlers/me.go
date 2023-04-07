package handlers

import (
	"errors"
	"financial-api/methods"
	"financial-api/models"
	"financial-api/responses"

	"github.com/gofiber/fiber/v2"
)

func GetMe(c *fiber.Ctx) error {
	user_id := c.Cookies("user_id")
	if user_id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(responses.BadUserID())
	}
	my_info, err := models.GetPublicInfo(user_id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	my_mileage, err := models.GetUserMileage(user_id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	my_checks, err := models.GetUserCheckRequests(user_id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	my_petty_cash, err := models.GetUserPettyCash(user_id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	my_info.Mileage_Requests = my_mileage
	my_info.Check_Requests = my_checks
	my_info.Petty_Cash_Requests = my_petty_cash
	return c.Status(fiber.StatusOK).JSON(responses.MyInfo(my_info))
}
func GetMyMileage(c *fiber.Ctx) error {
	user_id := c.Cookies("user_id")
	if user_id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(responses.BadUserID())
	}
	my_mileage, err := models.GetUserMileageDetail(user_id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(responses.MyMileage(my_mileage))
}
func GetMyCheckRequests(c *fiber.Ctx) error {
	user_id := c.Cookies("user_id")
	if user_id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(responses.BadUserID())
	}
	my_checks, err := models.GetUserCheckRequestDetail(user_id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(responses.MyCheckRequests(my_checks))
}
func GetMyPettyCash(c *fiber.Ctx) error {
	user_id := c.Cookies("user_id")
	my_petty_cash, err := models.GetUserPettyCashDetail(user_id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(responses.MyPettyCash(my_petty_cash))
}
func AddVehicle(c *fiber.Ctx) error {
	var mr *methods.MalformedRequest
	vehicle_input := new(models.VehicleInput)
	err := methods.DecodeJSONBody(c, vehicle_input)
	if err != nil {
		if errors.As(err, &mr) {
			return c.Status(mr.Status).JSON(responses.MalformedRequest(mr.Status, mr.Msg))
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
		}
	} else {
		if err := c.BodyParser(vehicle_input); err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(responses.MalformedRequest(422, err.Error()))
		}
	}
	errors := methods.ValidateStruct(*vehicle_input)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.MalformedBody(errors))
	}
	user_id := c.Cookies("user_id")
	if user_id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(responses.BadUserID())
	}
	user := new(models.User)
	user.ID = user_id
	vehicle, err := user.AddVehicle(vehicle_input.Name, vehicle_input.Description)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	return c.Status(fiber.StatusCreated).JSON(responses.AddVehicle(vehicle))
}
func EditVehicle(c *fiber.Ctx) error {
	var mr *methods.MalformedRequest
	vehicle_input := new(models.Vehicle)
	err := methods.DecodeJSONBody(c, vehicle_input)
	if err != nil {
		if errors.As(err, &mr) {
			return c.Status(mr.Status).JSON(responses.MalformedRequest(mr.Status, mr.Msg))
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
		}
	} else {
		c.BodyParser(vehicle_input)
	}
	errors := methods.ValidateStruct(*vehicle_input)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.MalformedBody(errors))
	}
	user_id := c.Cookies("user_id")
	if user_id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(responses.BadUserID())
	}
	user := new(models.User)
	user.ID = user_id
	vehicle, err := user.EditVehicle(*vehicle_input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(responses.EditVehicle(vehicle))
}
func RemoveVehicle(c *fiber.Ctx) error {
	var mr *methods.MalformedRequest
	vehicle_input := new(models.Vehicle)
	err := methods.DecodeJSONBody(c, vehicle_input)
	if err != nil {
		if errors.As(err, &mr) {
			return c.Status(mr.Status).JSON(responses.MalformedRequest(mr.Status, mr.Msg))
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
		}
	} else {
		c.BodyParser(vehicle_input)
	}
	errors := methods.ValidateStruct(*vehicle_input)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.MalformedBody(errors))
	}
	user_id := c.Cookies("user_id")
	if user_id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(responses.BadUserID())
	}
	user := new(models.User)
	user.ID = user_id
	vehicle, err := user.RemoveVehicle(vehicle_input.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(responses.RemoveVehicle(vehicle))
}
