package handlers

import (
	"errors"
	"financial-api/methods"
	"financial-api/models"
	"financial-api/responses"

	"github.com/gofiber/fiber/v2"
)

func CreateMileage(c *fiber.Ctx) error {
	var mr *methods.MalformedRequest
	mileage_request := new(models.MileageInput)
	err := methods.DecodeJSONBody(c, mileage_request)
	if err != nil {
		if errors.As(err, &mr) {
			return c.Status(mr.Status).JSON(responses.MalformedRequest(mr.Status, mr.Msg))
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
		}
	} else {
		c.BodyParser(mileage_request)
	}
	errors := methods.ValidateStruct(*mileage_request)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.MalformedBody(errors))
	}
	user_id := c.Cookies("user_id")
	if user_id == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.BadUserID())
	}
	res, err := mileage_request.CreateMileage(user_id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	return c.Status(fiber.StatusCreated).JSON(responses.CreateMileage(res))
}
func MileageVariance(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).JSON(responses.MileageVariance())
}
func MileageDetail(c *fiber.Ctx) error {
	var mr *methods.MalformedRequest
	find_mileage_input := new(models.FindMileageInput)
	err := methods.DecodeJSONBody(c, find_mileage_input)
	if err != nil {
		if errors.As(err, &mr) {
			return c.Status(mr.Status).JSON(responses.MalformedRequest(mr.Status, mr.Msg))
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
		}
	} else {
		c.BodyParser(find_mileage_input)
	}
	errors := methods.ValidateStruct(*find_mileage_input)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.MalformedBody(errors))
	}
	user_id := c.Cookies("user_id")
	if user_id == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.BadUserID())
	}
	data, err := models.MileageDetail(find_mileage_input.MileageID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(responses.MileageDetail(data))
}
func EditMileage(c *fiber.Ctx) error {
	var mr *methods.MalformedRequest
	mileage_request := new(models.EditMileageInput)
	err := methods.DecodeJSONBody(c, mileage_request)
	if err != nil {
		if errors.As(err, &mr) {
			return c.Status(mr.Status).JSON(responses.MalformedRequest(mr.Status, mr.Msg))
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
		}
	} else {
		c.BodyParser(mileage_request)
	}
	errors := methods.ValidateStruct(*mileage_request)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.MalformedBody(errors))
	}
	user_id := c.Cookies("user_id")
	if user_id == "" || user_id != mileage_request.User_ID {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.BadUserID())
	}
	response, err := mileage_request.EditMileage()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(responses.EditMileage(response))
}
func DeleteMileage(c *fiber.Ctx) error {
	var mr *methods.MalformedRequest
	find_mileage_input := new(models.FindMileageInput)
	err := methods.DecodeJSONBody(c, find_mileage_input)
	if err != nil {
		if errors.As(err, &mr) {
			return c.Status(mr.Status).JSON(responses.MalformedRequest(mr.Status, mr.Msg))
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
		}
	} else {
		c.BodyParser(find_mileage_input)
	}
	errors := methods.ValidateStruct(*find_mileage_input)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.MalformedBody(errors))
	}
	user_id := c.Cookies("user_id")
	if user_id == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.BadUserID())
	}
	data, err := models.DeleteMileage(find_mileage_input.MileageID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(responses.DeleteMileage(data))
}
func ApproveMileage(c *fiber.Ctx) error {
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
	mileage := new(models.Mileage_Request)
	mileage.ID = approve_info.RequestID
	data, err := mileage.Approve(user_id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(responses.ApproveMileage(data))
}
func RejectMileage(c *fiber.Ctx) error {
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
	mileage := new(models.Mileage_Request)
	mileage.ID = reject_info.RequestID
	data, err := mileage.Reject(user_id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(responses.RejectMileage(data))
}
func MonthlyMileage(c *fiber.Ctx) error {
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
	response, err := models.MonthlyMileage(int(monthly_request.Month), monthly_request.Year)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(responses.MonthlyMileage(int(monthly_request.Month), monthly_request.Year, response))
}
