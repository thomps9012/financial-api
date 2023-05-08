package handlers

import (
	"errors"
	"financial-api/methods"
	"financial-api/models"
	"financial-api/responses"
	"regexp"

	"github.com/gofiber/fiber/v2"
)

// @id login-user
// @summary login user
// @description either logs a user in or creates an account depending on previous use of the application
// @tags user, no-cache, auth
// @param login-info body models.UserLogin true "user's account information"
// @produce json
// @success 201 {object} responses.UserLoginRes
// @success 200 {object} responses.UserLoginRes
// @router /auth/login [post]
func Login(c *fiber.Ctx) error {
	var mr *methods.MalformedRequest
	user_login := new(models.UserLogin)
	err := methods.DecodeJSONBody(c, user_login)
	if err != nil {
		if errors.As(err, &mr) {
			return c.Status(mr.Status).JSON(responses.MalformedRequest(mr.Status, mr.Msg))
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
		}
	} else {
		c.BodyParser(user_login)
	}
	errors := methods.ValidateStruct(*user_login)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.MalformedBody(errors))
	}
	exists, exists_err := user_login.Exists()
	if exists_err != nil {
		return c.Status(exists_err.Status).JSON(responses.ServerError(exists_err.Error()))
	}
	valid_email, email_err := regexp.Match(`\w+@norainc.org`, []byte(user_login.Email))
	if !valid_email || email_err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.InvalidEmail())
	}
	user := new(models.User)
	if !exists {
		login_res, create_err := user.Create(*user_login)
		if create_err != nil {
			return c.Status(create_err.Status).JSON(responses.ServerError(create_err.Error()))
		}
		c.Response().Header.Add("no-cache", "true")
		return c.Status(fiber.StatusCreated).JSON(responses.NewUser(login_res))
	} else {
		login_res, login_err := user.Login(*user_login)
		if login_err != nil {
			return c.Status(login_err.Status).JSON(responses.ServerError(login_err.Error()))
		}
		c.Response().Header.Add("no-cache", "true")
		return c.Status(fiber.StatusOK).JSON(responses.LoggedIn(login_res))
	}
}

// @id logout-user
// @summary logout user
// @description logs a user out
// @tags user, no-cache, auth
// @produce json
// @success 200 {object} responses.NilRes
// @router /auth/logout [post]
func Logout(c *fiber.Ctx) error {
	c.Response().Header.Add("no-cache", "true")
	return c.Status(fiber.StatusOK).JSON(responses.LoggedOut())
}
