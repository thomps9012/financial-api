package handlers

import (
	"errors"
	"financial-api/methods"
	"financial-api/middleware"
	"financial-api/models"
	"financial-api/responses"

	"github.com/gofiber/fiber/v2"
)

// @id get-me
// @summary get my info
// @description gathers basic request information for a logged in user and the requests' current status
// @tags user, reports
// @produce json
// @success 200 {object} responses.MyInfoRes
// @router /me [get]
func GetMe(c *fiber.Ctx) error {
	user_id, err := middleware.TokenID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.KeyNotFound())
	}
	if user_id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(responses.BadUserID())
	}
	my_info, err := models.GetPublicInfo(user_id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(responses.MyInfo(my_info))
}

// @id get-my-mileage
// @summary get my mileage reqs
// @description gathers more detailed request information for a logged in user's mileage requests
// @tags user, mileage, reports
// @produce json
// @success 200 {object} responses.MileagesRes
// @router /me/mileage [get]
func GetMyMileage(c *fiber.Ctx) error {
	user_id, err := middleware.TokenID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.KeyNotFound())
	}
	if user_id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(responses.BadUserID())
	}
	my_mileage, err := models.GetUserMileageDetail(user_id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(responses.MyMileage(my_mileage))
}

// @id get-my-checks
// @summary get my check reqs
// @description gathers more detailed request information for a logged in user's check requests
// @tags user, check, reports
// @produce json
// @success 200 {object} responses.CheckRequestsRes
// @router /me/check [get]
func GetMyCheckRequests(c *fiber.Ctx) error {
	user_id, err := middleware.TokenID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.KeyNotFound())
	}
	if user_id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(responses.BadUserID())
	}
	my_checks, err := models.GetUserCheckRequestDetail(user_id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(responses.MyCheckRequests(my_checks))
}

// @id get-my-petty-cash
// @summary get my petty cash reqs
// @description gathers more detailed request information for a logged in user's petty cash requests
// @tags user, petty cash, reports
// @produce json
// @success 200 {object} responses.PettyCashRequestsRes
// @router /me/petty_cash [get]
func GetMyPettyCash(c *fiber.Ctx) error {
	user_id, err := middleware.TokenID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.KeyNotFound())
	}
	my_petty_cash, err := models.GetUserPettyCashDetail(user_id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(responses.MyPettyCash(my_petty_cash))
}

// @id add-vehicle
// @summary add a vehicle
// @description adds vehicle information to a logged in user's account
// @param vehicle-info body models.VehicleInput true "user's vehicle information"
// @tags user, no-cache
// @produce json
// @success 200 {object} responses.VehicleRes
// @router /me/vehicle [post]
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
	user_id, err := middleware.TokenID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.KeyNotFound())
	}
	if user_id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(responses.BadUserID())
	}
	user := new(models.User)
	user.ID = user_id
	vehicle, err := user.AddVehicle(vehicle_input.Name, vehicle_input.Description)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	c.Response().Header.Add("no-cache", "true")
	return c.Status(fiber.StatusCreated).JSON(responses.AddVehicle(vehicle))
}

// @id edit-vehicle
// @summary edit a vehicle
// @description edits information about a specific vehicle for a logged in user
// @param vehicle-info body models.Vehicle true "edited vehicle information"
// @tags user, no-cache
// @produce json
// @success 200 {object} responses.VehicleRes
// @router /me/vehicle [put]
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
	user_id, err := middleware.TokenID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.KeyNotFound())
	}
	if user_id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(responses.BadUserID())
	}
	user := new(models.User)
	user.ID = user_id
	vehicle, err := user.EditVehicle(*vehicle_input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	c.Response().Header.Add("no-cache", "true")
	return c.Status(fiber.StatusOK).JSON(responses.EditVehicle(vehicle))
}

// @id remove-vehicle
// @summary remove a vehicle
// @description removes vehicle information from a logged in user's account
// @param vehicle-info body models.Vehicle true "deleted vehicle information"
// @tags user, no-cache
// @produce json
// @success 200 {object} responses.VehicleRes
// @router /me/vehicle [delete]
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
	user_id, err := middleware.TokenID(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.KeyNotFound())
	}
	if user_id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(responses.BadUserID())
	}
	user := new(models.User)
	user.ID = user_id
	vehicle, err := user.RemoveVehicle(vehicle_input.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	c.Response().Header.Add("no-cache", "true")
	return c.Status(fiber.StatusOK).JSON(responses.RemoveVehicle(vehicle))
}
