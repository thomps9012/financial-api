package handlers

import (
	"errors"
	"financial-api/methods"
	"financial-api/models"
	"financial-api/responses"

	"github.com/gofiber/fiber/v2"
)

// @id get-users
// @summary get all users
// @description gathers basic request information for all users in the system and the current status of their requests'
// @tags user, reports, admin
// @produce json
// @success 200
// @router /user [get]
func GetAllUsers(c *fiber.Ctx) error {
	all_users, err := models.FindAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(responses.AllUsers(all_users))
}

type UserIDBody struct {
	UserID string `json:"user_id" bson:"user_id" validate:"required"`
}

// @id get-user
// @summary gets one user
// @description gathers basic request information for a specific user in the system and the current status of their requests'
// @param user-info body UserIDBody true "specific user's id information"
// @tags user, reports, admin
// @produce json
// @success 200
// @router /user/detail [get]
func GetOneUser(c *fiber.Ctx) error {
	var mr *methods.MalformedRequest
	user_id_body := new(UserIDBody)
	err := methods.DecodeJSONBody(c, user_id_body)
	if err != nil {
		if errors.As(err, &mr) {
			return c.Status(mr.Status).JSON(responses.MalformedRequest(mr.Status, mr.Msg))
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
		}
	} else {
		c.BodyParser(user_id_body)
	}
	errors := methods.ValidateStruct(*user_id_body)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.MalformedBody(errors))
	}
	user_info, err := models.GetPublicInfo(user_id_body.UserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(responses.OneUser(user_info))
}

// @id delete-user
// @summary deactivates a user
// @description deactivates a user's account and all of their associated requests
// @param user-info body UserIDBody true "specific user's id information"
// @tags user, no-cache, admin
// @produce json
// @success 200
// @router /user/deactivate [delete]
func DeactivateUser(c *fiber.Ctx) error {
	var mr *methods.MalformedRequest
	user_id_body := new(UserIDBody)
	err := methods.DecodeJSONBody(c, user_id_body)
	if err != nil {
		if errors.As(err, &mr) {
			return c.Status(mr.Status).JSON(responses.MalformedRequest(mr.Status, mr.Msg))
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
		}
	} else {
		c.BodyParser(user_id_body)
	}
	errors := methods.ValidateStruct(*user_id_body)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.MalformedBody(errors))
	}
	user := new(models.User)
	user.ID = user_id_body.UserID
	user_info, err := user.Deactivate()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	c.Response().Header.Add("no-cache", "true")
	return c.Status(fiber.StatusOK).JSON(responses.DeactivateUser(user_info))
}

// @id user-mileage
// @summary user mileage
// @description gathers more detailed information on a specific user's mileage requests
// @param user-info body UserIDBody true "specific user's id information"
// @tags user, reports, mileage, admin
// @produce json
// @success 200
// @router /user/mileage [get]
func UserMileage(c *fiber.Ctx) error {
	var mr *methods.MalformedRequest
	user_id_body := new(UserIDBody)
	err := methods.DecodeJSONBody(c, user_id_body)
	if err != nil {
		if errors.As(err, &mr) {
			return c.Status(mr.Status).JSON(responses.MalformedRequest(mr.Status, mr.Msg))
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
		}
	} else {
		c.BodyParser(user_id_body)
	}
	errors := methods.ValidateStruct(*user_id_body)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.MalformedBody(errors))
	}
	user_info, err := models.GetUserMileage(user_id_body.UserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(responses.UserMileage(user_id_body.UserID, user_info))
}

// @id user-petty_cash
// @summary user petty cash
// @description gathers more detailed information on a specific user's petty cash requests
// @param user-info body UserIDBody true "specific user's id information"
// @tags user, reports, petty cash, admin
// @produce json
// @success 200
// @router /user/petty_cash [get]
func UserPettyCash(c *fiber.Ctx) error {
	var mr *methods.MalformedRequest
	user_id_body := new(UserIDBody)
	err := methods.DecodeJSONBody(c, user_id_body)
	if err != nil {
		if errors.As(err, &mr) {
			return c.Status(mr.Status).JSON(responses.MalformedRequest(mr.Status, mr.Msg))
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
		}
	} else {
		c.BodyParser(user_id_body)
	}
	errors := methods.ValidateStruct(*user_id_body)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.MalformedBody(errors))
	}
	user_info, err := models.GetUserPettyCash(user_id_body.UserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(responses.UserPettyCash(user_id_body.UserID, user_info))
}

// @id user-check
// @summary user check requests
// @description gathers more detailed information on a specific user's check requests
// @param user-info body UserIDBody true "specific user's id information"
// @tags user, reports, check, admin
// @produce json
// @success 200
// @router /user/check [get]
func UserCheckRequests(c *fiber.Ctx) error {
	var mr *methods.MalformedRequest
	user_id_body := new(UserIDBody)
	err := methods.DecodeJSONBody(c, user_id_body)
	if err != nil {
		if errors.As(err, &mr) {
			return c.Status(mr.Status).JSON(responses.MalformedRequest(mr.Status, mr.Msg))
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
		}
	} else {
		c.BodyParser(user_id_body)
	}
	errors := methods.ValidateStruct(*user_id_body)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.MalformedBody(errors))
	}
	user_info, err := models.GetUserCheckRequests(user_id_body.UserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(responses.UserCheckRequests(user_id_body.UserID, user_info))
}
