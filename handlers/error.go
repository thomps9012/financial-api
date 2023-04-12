package handlers

import (
	"errors"
	"financial-api/methods"
	"financial-api/models"
	"financial-api/responses"

	"github.com/gofiber/fiber/v2"
)

// @id log-error
// @summary logs errors
// @description logs client side errors to allow for faster issue remediation
// @param error-info body models.ErrorLog true "error log information"
// @tags errors
// @produce json
// @success 200 {object} responses.ErrorLogRes
// @router /error [post]
func LogError(c *fiber.Ctx) error {
	var mr *methods.MalformedRequest
	request := new(models.ErrorLog)
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
	res, err := request.Save()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	c.Response().Header.Add("no-cache", "true")
	return c.Status(fiber.StatusCreated).JSON(responses.LogError(res))
}
