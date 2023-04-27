package handlers

import (
	"errors"
	"financial-api/methods"
	"financial-api/models"
	"financial-api/responses"

	"github.com/gofiber/fiber/v2"
)

// @id add-petty_cash
// @summary add petty cash request
// @description creates a petty cash request for a logged in user
// @param petty-cash-request-info body models.PettyCashInput true "new petty cash request information"
// @tags petty cash, no-cache
// @produce json
// @success 201 {object} responses.PettyCashOverviewRes
// @router /petty_cash [post]
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
	exists, err := request.Exists(user_id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	if exists {
		return c.Status(fiber.StatusBadRequest).JSON(responses.RequestExists("petty cash"))
	}
	res, err := request.CreatePettyCash(user_id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	c.Response().Header.Add("no-cache", "true")
	return c.Status(fiber.StatusCreated).JSON(responses.CreatePettyCash(res))
}

// @id petty_cash-detail
// @summary petty cash request detail
// @description generates detailed information for a specific petty cash request
// @param petty-cash-request-id body models.FindPettyCashInput true "petty cash request id to find"
// @tags petty cash
// @produce json
// @success 200 {object} responses.PettyCashRes
// @router /petty_cash/detail [get]
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

// @id edit-petty_cash
// @summary edit petty cash request
// @description allows a logged in user to edit their pending or reject petty cash requests
// @param petty-cash-request-info body models.EditPettyCash true "petty cash request information to update"
// @tags petty cash, no-cache
// @produce json
// @success 200 {object} responses.PettyCashOverviewRes
// @router /petty_cash [put]
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

// @id delete-petty_cash
// @summary delete petty cash request
// @description allows a logged in user to delete and removes one of their petty cash requests
// @param petty-cash-request-id body models.FindPettyCashInput true "petty cash request id to delete"
// @tags petty cash, no-cache
// @produce json
// @success 200 {object} responses.PettyCashOverviewRes
// @router /petty_cash [delete]
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
	err = models.ClearRequestAssociatedActions(find_request_input.PettyCashID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	c.Response().Header.Add("no-cache", "true")
	return c.Status(fiber.StatusOK).JSON(responses.DeletePettyCash(data))
}

// @id approve-petty_cash
// @summary approve petty cash request
// @description allows an administrative user to approve a specific petty cash request
// @param request-id body models.ApproveRejectRequest true "petty cash request id to approve"
// @tags petty cash, no-cache, admin
// @produce json
// @success 200 {object} responses.PettyCashOverviewRes
// @router /petty_cash/approve [post]
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
	err = models.ClearIncompleteAction("petty_cash", request.ID, user_id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	c.Response().Header.Add("no-cache", "true")
	return c.Status(fiber.StatusOK).JSON(responses.ApprovePettyCash(data))
}

// @id reject-petty_cash
// @summary reject petty cash request
// @description allows a logged in administrative user to reject a petty cash request until further edits have been made
// @param request-id body models.ApproveRejectRequest true "petty cash request id to reject"
// @tags petty cash, no-cache, admin
// @produce json
// @success 200 {object} responses.PettyCashOverviewRes
// @router /petty_cash/reject [post]
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
	err = models.ClearIncompleteAction("petty_cash", request.ID, user_id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	c.Response().Header.Add("no-cache", "true")
	return c.Status(fiber.StatusOK).JSON(responses.RejectPettyCash(data))
}

// @id monthly-petty_cash
// @summary monthly petty cash request report
// @description creates a monthly petty cash report for the entire organization for a given month and year
// @param month-year-input body models.MonthlyRequestInput true "month and year for report on organization wide petty cash requests"
// @tags petty cash, reports, admin
// @produce json
// @success 200 {object} responses.PettyCashOverviewsRes
// @router /petty_cash/monthly [get]
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
