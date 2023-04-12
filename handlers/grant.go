package handlers

import (
	"errors"
	"financial-api/methods"
	"financial-api/models"
	"financial-api/responses"

	"github.com/gofiber/fiber/v2"
)

// @id get-grants
// @summary all grant information
// @description generates basic information for all organization-wide grants
// @tags grants
// @produce json
// @success 200
// @router /grants [get]
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

// @id get-grant
// @summary one grant information
// @description generates basic information for a specified grant
// @param grant-info body models.Grant true "grant id"
// @tags grants
// @produce json
// @success 200
// @router /grants/detail [get]
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

// @id get-grant-check
// @summary one grant's check requests
// @description generates an overview of check requests for a specified grant
// @tags grants, check, reports, admin
// @param grant-info body models.Grant true "grant id for which check request report should be run"
// @produce json
// @success 200
// @router /grants/check [get]
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

// @id get-grant-mileage
// @summary one grant's mileage requests
// @description generates an overview of mileage requests for a specified grant
// @param grant-info body models.Grant true "grant id for which mileage request report should be run"
// @tags grants, mileage, reports, admin
// @produce json
// @success 200
// @router /grants/mileage [get]
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

// @id get-grant-petty_cash
// @summary one grant's petty cash requests
// @description generates an overview of petty cash requests for a specified grant
// @param grant-info body models.Grant true "grant id for which petty cash request report should be run"
// @tags grants, petty cash, reports, admin
// @produce json
// @success 200
// @router /grants/petty_cash [get]
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
