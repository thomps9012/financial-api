package handlers

import (
	"errors"
	"financial-api/methods"
	"financial-api/models"
	"financial-api/responses"

	"github.com/gofiber/fiber/v2"
)

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
	return c.Status(fiber.StatusOK).JSON(responses.DeactivateUser(user_info))
}
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
