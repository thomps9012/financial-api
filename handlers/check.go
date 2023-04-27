package handlers

import (
	"errors"
	"financial-api/methods"
	"financial-api/models"
	"financial-api/responses"

	"github.com/gofiber/fiber/v2"
)

// @id add-check
// @summary add check request
// @description creates a check request for a logged in user
// @param check-request-info body models.CheckRequestInput true "new check request information"
// @tags check, no-cache
// @produce json
// @success 201 {object} responses.CheckOverviewRes
// @router /check [post]
func CreateCheckRequest(c *fiber.Ctx) error {
	var mr *methods.MalformedRequest
	request := new(models.CheckRequestInput)
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
	exists, err := request.Exists(user_id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	if exists {
		return c.Status(fiber.StatusBadRequest).JSON(responses.RequestExists("check"))
	}
	res, err := request.CreateCheckRequest(user_id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	c.Response().Header.Add("no-cache", "true")
	return c.Status(fiber.StatusCreated).JSON(responses.NewCheckRequest(res))
}

// @id check-detail
// @summary check request detail
// @description generates detailed information for a specific check request
// @param check-request-id body models.FindCheckInput true "check request id to find"
// @tags check
// @produce json
// @success 200 {object} responses.CheckDetailRes
// @router /check/detail [get]
func CheckRequestDetail(c *fiber.Ctx) error {
	var mr *methods.MalformedRequest
	find_check_input := new(models.FindCheckInput)
	err := methods.DecodeJSONBody(c, find_check_input)
	if err != nil {
		if errors.As(err, &mr) {
			return c.Status(mr.Status).JSON(responses.MalformedRequest(mr.Status, mr.Msg))
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
		}
	} else {
		c.BodyParser(find_check_input)
	}
	errors := methods.ValidateStruct(*find_check_input)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.MalformedBody(errors))
	}
	user_id := c.Cookies("user_id")
	if user_id == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.BadUserID())
	}
	data, err := models.CheckRequestDetail(find_check_input.CheckID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(responses.CheckRequestDetail(data))
}

// @id edit-check
// @summary edit check request
// @description allows a logged in user to edit their pending or reject check requests
// @param check-request-info body models.EditCheckInput true "check request information to update"
// @tags check, no-cache
// @produce json
// @success 200 {object} responses.CheckOverviewRes
// @router /check [put]
func EditCheckRequest(c *fiber.Ctx) error {
	var mr *methods.MalformedRequest
	request := new(models.EditCheckInput)
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
	response, err := request.EditCheckRequest()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	c.Response().Header.Add("no-cache", "true")
	return c.Status(fiber.StatusOK).JSON(responses.EditCheckRequest(response))
}

// @id delete-check
// @summary delete check request
// @description allows a logged in user to delete and removes one of their check requests
// @param check-request-id body models.FindCheckInput true "check request id to delete"
// @tags check, no-cache
// @produce json
// @success 200 {object} responses.CheckOverviewRes
// @router /check [delete]
func DeleteCheckRequest(c *fiber.Ctx) error {
	var mr *methods.MalformedRequest
	find_request_input := new(models.FindCheckInput)
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
	data, err := models.DeleteCheckRequest(find_request_input.CheckID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	err = models.ClearRequestAssociatedActions(find_request_input.CheckID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	c.Response().Header.Add("no-cache", "true")
	return c.Status(fiber.StatusOK).JSON(responses.DeleteCheckRequest(data))
}

// @id approve-check
// @summary approve check request
// @description allows an administrative user to approve a specific check request
// @param request-id body models.ApproveRejectRequest true "check request id to approve"
// @tags check, no-cache, admin
// @produce json
// @success 200 {object} responses.CheckOverviewRes
// @router /check/approve [post]
func ApproveCheckRequest(c *fiber.Ctx) error {
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
	request := new(models.Check_Request)
	request.ID = approve_info.RequestID
	data, err := request.Approve(user_id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	err = models.ClearIncompleteAction("check", request.ID, user_id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	c.Response().Header.Add("no-cache", "true")
	return c.Status(fiber.StatusOK).JSON(responses.ApproveCheckRequest(data))
}

// @id reject-check
// @summary reject check request
// @description allows a logged in administrative user to reject a check request until further edits have been made
// @param request-id body models.ApproveRejectRequest true "check request id to reject"
// @tags check, no-cache, admin
// @produce json
// @success 200 {object} responses.CheckOverviewRes
// @router /check/reject [post]
func RejectCheckRequest(c *fiber.Ctx) error {
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
	request := new(models.Check_Request)
	request.ID = reject_info.RequestID
	data, err := request.Reject(user_id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	err = models.ClearIncompleteAction("check", request.ID, user_id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	c.Response().Header.Add("no-cache", "true")
	return c.Status(fiber.StatusOK).JSON(responses.RejectCheckRequest(data))
}

// @id monthly-check
// @summary monthly check request report
// @description creates a monthly check report for the entire organization for a given month and year
// @param month-year-input body models.MonthlyRequestInput true "month and year for report on organization wide check requests"
// @tags check, reports, admin
// @produce json
// @success 200 {object} responses.CheckOverviewsRes
// @router /check/monthly [get]
func MonthlyCheckRequests(c *fiber.Ctx) error {
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
	response, err := models.MonthlyCheckRequests(int(monthly_request.Month), monthly_request.Year)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(responses.MonthlyCheckRequests(int(monthly_request.Month), monthly_request.Year, response))
}
