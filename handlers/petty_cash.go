package handlers

import (
	"errors"
	"financial-api/methods"
	"financial-api/models"
	"financial-api/responses"

	"github.com/gofiber/fiber/v2"
)

func CreatePettyCash(c *fiber.Ctx) error {
	var mr *methods.MalformedRequest
	request := new(models.PettyCashInput)
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
	user_id := c.Cookies("user_id")
	if user_id == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.BadUserID())
	}
	res, err := request.CreatePettyCash(user_id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	c.Response().Header.Add("no-cache", "true")
	return c.Status(fiber.StatusCreated).JSON(responses.CreatePettyCash(res))
}
func PettyCashDetail(c *fiber.Ctx) error {
	var mr *methods.MalformedRequest
	find_request_input := new(models.FindPettyCashInput)
	err := methods.DecodeJSONBody(c, find_request_input)
	if err != nil {
		if errors.As(err, &mr) {
			return c.Status(mr.Status).JSON(responses.MalformedRequest(mr.Status, mr.Msg))
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
		}
	} else {
		c.BodyParser(find_request_input)
	}
	errors := methods.ValidateStruct(*find_request_input)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.MalformedBody(errors))
	}
	user_id := c.Cookies("user_id")
	if user_id == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.BadUserID())
	}
	data, err := models.PettyCashDetails(find_request_input.PettyCashID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(responses.PettyCashDetail(data))
}
func EditPettyCash(c *fiber.Ctx) error {
	var mr *methods.MalformedRequest
	request := new(models.EditPettyCash)
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
	user_id := c.Cookies("user_id")
	if user_id == "" || user_id != request.User_ID {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.BadUserID())
	}
	response, err := request.EditPettyCash()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	c.Response().Header.Add("no-cache", "true")
	return c.Status(fiber.StatusOK).JSON(responses.EditPettyCash(response))
}
func DeletePettyCash(c *fiber.Ctx) error {
	var mr *methods.MalformedRequest
	find_request_input := new(models.FindPettyCashInput)
	err := methods.DecodeJSONBody(c, find_request_input)
	if err != nil {
		if errors.As(err, &mr) {
			return c.Status(mr.Status).JSON(responses.MalformedRequest(mr.Status, mr.Msg))
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
		}
	} else {
		c.BodyParser(find_request_input)
	}
	errors := methods.ValidateStruct(*find_request_input)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.MalformedBody(errors))
	}
	user_id := c.Cookies("user_id")
	if user_id == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.BadUserID())
	}
	data, err := models.DeletePettyCash(find_request_input.PettyCashID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	c.Response().Header.Add("no-cache", "true")
	return c.Status(fiber.StatusOK).JSON(responses.DeletePettyCash(data))
}
func ApprovePettyCash(c *fiber.Ctx) error {
	var mr *methods.MalformedRequest
	approve_info := new(models.ApproveRejectRequest)
	err := methods.DecodeJSONBody(c, approve_info)
	if err != nil {
		if errors.As(err, &mr) {
			return c.Status(mr.Status).JSON(responses.MalformedRequest(mr.Status, mr.Msg))
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
		}
	} else {
		c.BodyParser(approve_info)
	}
	errors := methods.ValidateStruct(*approve_info)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.MalformedBody(errors))
	}
	user_id := c.Cookies("user_id")
	if user_id == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.BadUserID())
	}
	request := new(models.Petty_Cash_Request)
	request.ID = approve_info.RequestID
	data, err := request.Approve(user_id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	c.Response().Header.Add("no-cache", "true")
	return c.Status(fiber.StatusOK).JSON(responses.ApprovePettyCash(data))
}
func RejectPettyCash(c *fiber.Ctx) error {
	var mr *methods.MalformedRequest
	reject_info := new(models.ApproveRejectRequest)
	err := methods.DecodeJSONBody(c, reject_info)
	if err != nil {
		if errors.As(err, &mr) {
			return c.Status(mr.Status).JSON(responses.MalformedRequest(mr.Status, mr.Msg))
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
		}
	} else {
		c.BodyParser(reject_info)
	}
	errors := methods.ValidateStruct(*reject_info)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.MalformedBody(errors))
	}
	user_id := c.Cookies("user_id")
	if user_id == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.BadUserID())
	}
	request := new(models.Petty_Cash_Request)
	request.ID = reject_info.RequestID
	data, err := request.Reject(user_id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	c.Response().Header.Add("no-cache", "true")
	return c.Status(fiber.StatusOK).JSON(responses.RejectPettyCash(data))
}
func MonthlyPettyCash(c *fiber.Ctx) error {
	var mr *methods.MalformedRequest
	monthly_request := new(models.MonthlyRequestInput)
	err := methods.DecodeJSONBody(c, monthly_request)
	if err != nil {
		if errors.As(err, &mr) {
			return c.Status(mr.Status).JSON(responses.MalformedRequest(mr.Status, mr.Msg))
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
		}
	} else {
		c.BodyParser(monthly_request)
	}
	errors := methods.ValidateStruct(*monthly_request)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.MalformedBody(errors))
	}
	response, err := models.MonthlyPettyCash(int(monthly_request.Month), monthly_request.Year)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(responses.MonthlyPettyCash(int(monthly_request.Month), monthly_request.Year, response))
}
