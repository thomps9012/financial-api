package handlers

import (
	"errors"
	"financial-api/methods"
	"financial-api/models"
	"financial-api/responses"

	"github.com/gofiber/fiber/v2"
)

// @id add-mileage
// @summary add mileage request
// @description creates a mileage request for a logged in user
// @param mileage-request-info body models.MileageInput true "new mileage request information"
// @tags mileage, no-cache
// @produce json
// @success 201 {object} responses.MileageOverviewRes
// @router /mileage [post]
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
	c.Response().Header.Add("no-cache", "true")
	return c.Status(fiber.StatusCreated).JSON(responses.CreateMileage(res))
}
func MileageVariance(c *fiber.Ctx) error {
	c.Response().Header.Add("no-cache", "true")
	return c.Status(fiber.StatusNotImplemented).JSON(responses.MileageVariance())
}

// @id mileage-detail
// @summary mileage request detail
// @description generates detailed information for a specific mileage request
// @param mileage-request-id body models.FindMileageInput true "mileage request id to find"
// @tags mileage
// @produce json
// @success 200 {object} responses.MileageRes
// @router /mileage/detail [get]
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

// @id edit-mileage
// @summary edit mileage request
// @description allows a logged in user to edit their pending or reject mileage requests
// @param mileage-request-info body models.EditMileageInput true "mileage request information to update"
// @tags mileage, no-cache
// @produce json
// @success 200 {object} responses.MileageOverviewRes
// @router /mileage [put]
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
	c.Response().Header.Add("no-cache", "true")
	return c.Status(fiber.StatusOK).JSON(responses.EditMileage(response))
}

// @id delete-mileage
// @summary delete mileage request
// @description allows a logged in user to delete and removes one of their mileage requests
// @param mileage-request-id body models.FindMileageInput true "mileage request id to delete"
// @tags mileage, no-cache
// @produce json
// @success 200 {object} responses.MileageOverviewRes
// @router /mileage [delete]
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
	err = models.ClearRequestAssociatedActions(find_mileage_input.MileageID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	c.Response().Header.Add("no-cache", "true")
	return c.Status(fiber.StatusOK).JSON(responses.DeleteMileage(data))
}

// @id approve-mileage
// @summary approve mileage request
// @description allows an administrative user to approve a specific mileage request
// @param request-id body models.ApproveRejectRequest true "mileage request id to approve"
// @tags mileage, no-cache, admin
// @produce json
// @success 200 {object} responses.MileageOverviewRes
// @router /mileage/approve [post]
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
	err = models.ClearIncompleteAction("mileage", mileage.ID, user_id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	c.Response().Header.Add("no-cache", "true")
	return c.Status(fiber.StatusOK).JSON(responses.ApproveMileage(data))
}

// @id reject-mileage
// @summary reject mileage request
// @description allows a logged in administrative user to reject a mileage request until further edits have been made
// @param request-id body models.ApproveRejectRequest true "mileage request id to reject"
// @tags mileage, no-cache, admin
// @produce json
// @success 200 {object} responses.MileageOverviewRes
// @router /mileage/reject [post]
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
	err = models.ClearIncompleteAction("mileage", mileage.ID, user_id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	c.Response().Header.Add("no-cache", "true")
	return c.Status(fiber.StatusOK).JSON(responses.RejectMileage(data))
}

// @id monthly-mileage
// @summary monthly mileage report
// @description creates a monthly mileage report for the entire organization for a given month and year
// @param month-year-input body models.MonthlyRequestInput true "month and year for report on organization wide mileage requests"
// @tags mileage, reports, admin
// @produce json
// @success 200 {object} responses.MileageOverviewsRes
// @router /mileage/monthly [get]
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
