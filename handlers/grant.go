package handlers

import (
	"errors"
	"financial-api/methods"
	"financial-api/models"
	"financial-api/responses"

	"github.com/gofiber/fiber/v2"
)

func GetAllGrants(c *fiber.Ctx) error {
	user_id := c.Cookies("user_id")
	if user_id == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.BadUserID())
	}
	data, err := models.GetAllGrants()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(responses.AllGrants(data))
}
func GetOneGrant(c *fiber.Ctx) error {
	var mr *methods.MalformedRequest
	request := new(models.Grant)
	err := methods.DecodeJSONBody(c, request)
	if err != nil {
		if errors.As(err, &mr) {
			return c.Status(mr.Status).JSON(responses.MalformedRequest(mr.Status, mr.Msg))
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
		}
	} else {
		c.BodyParser(request)
	}
	errors := methods.ValidateStruct(*request)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.MalformedBody(errors))
	}
	data, err := request.GetOneGrant()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(responses.OneGrant(data))
}
func GrantCheckRequests(c *fiber.Ctx) error {
	var mr *methods.MalformedRequest
	request := new(models.Grant)
	err := methods.DecodeJSONBody(c, request)
	if err != nil {
		if errors.As(err, &mr) {
			return c.Status(mr.Status).JSON(responses.MalformedRequest(mr.Status, mr.Msg))
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
		}
	} else {
		c.BodyParser(request)
	}
	errors := methods.ValidateStruct(*request)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.MalformedBody(errors))
	}
	grant_data, err := request.GetOneGrant()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	data, err := request.GetGrantCheckRequest()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(responses.GrantCheckRequests(grant_data, data))
}
func GrantMileage(c *fiber.Ctx) error {
	var mr *methods.MalformedRequest
	request := new(models.Grant)
	err := methods.DecodeJSONBody(c, request)
	if err != nil {
		if errors.As(err, &mr) {
			return c.Status(mr.Status).JSON(responses.MalformedRequest(mr.Status, mr.Msg))
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
		}
	} else {
		c.BodyParser(request)
	}
	errors := methods.ValidateStruct(*request)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.MalformedBody(errors))
	}
	grant_data, err := request.GetOneGrant()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	data, err := request.GetGrantMileage()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(responses.GrantMileage(grant_data, data))
}
func GrantPettyCash(c *fiber.Ctx) error {
	var mr *methods.MalformedRequest
	request := new(models.Grant)
	err := methods.DecodeJSONBody(c, request)
	if err != nil {
		if errors.As(err, &mr) {
			return c.Status(mr.Status).JSON(responses.MalformedRequest(mr.Status, mr.Msg))
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
		}
	} else {
		c.BodyParser(request)
	}
	errors := methods.ValidateStruct(*request)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.MalformedBody(errors))
	}
	grant_data, err := request.GetOneGrant()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	data, err := request.GetGrantPettyCash()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(responses.GrantPettyCash(grant_data, data))
}
