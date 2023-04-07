package handlers

import (
	"errors"
	"financial-api/methods"
	"financial-api/models"
	"financial-api/responses"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func SetCookies(c *fiber.Ctx, login_res models.LoginRes, complete chan bool) {
	c.Cookie(&fiber.Cookie{
		Name:  "admin",
		Value: strconv.FormatBool(login_res.Admin),
	})
	c.Cookie(&fiber.Cookie{
		Name:  "permissions",
		Value: strings.Join(login_res.Permissions, ", "),
	})
	c.Cookie(&fiber.Cookie{
		Name:  "user_id",
		Value: login_res.UserID,
	})
	complete <- true
}

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
	exists, err := user_login.Exists()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
	}
	user := new(models.User)
	cookies_set := make(chan bool)
	if !exists {
		login_res, err := user.Create(*user_login)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
		}
		go SetCookies(c, login_res, cookies_set)
		complete := <-cookies_set
		if complete {
			return c.Status(fiber.StatusCreated).JSON(responses.NewUser(login_res.Token))
		}
	} else {
		login_res, err := user.Login(*user_login)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
		}
		go SetCookies(c, login_res, cookies_set)
		complete := <-cookies_set
		if complete {
			return c.Status(fiber.StatusOK).JSON(responses.LoggedIn(login_res.Token))
		}
	}
	return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError("error setting cookies"))
}

func Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:  "admin",
		Value: "",
	})
	c.Cookie(&fiber.Cookie{
		Name:  "permissions",
		Value: "",
	})
	c.Cookie(&fiber.Cookie{
		Name:  "user_id",
		Value: "",
	})
	return c.Status(fiber.StatusAccepted).JSON(responses.LoggedOut())
}
