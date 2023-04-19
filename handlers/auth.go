package handlers

import (
	"errors"
	"financial-api/methods"
	"financial-api/models"
	"financial-api/responses"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func SetCookies(c *fiber.Ctx, login_res models.LoginRes, complete chan bool) {
	admin_cookie := new(fiber.Cookie)
	admin_cookie.Name = "admin"
	admin_cookie.Value = strconv.FormatBool(login_res.Admin)
	admin_cookie.Expires = time.Now().Add(time.Hour * 12)
	c.Cookie(admin_cookie)
	permissions_cookie := new(fiber.Cookie)
	permissions_cookie.Name = "permissions"
	permissions_cookie.Value = strings.Join(login_res.Permissions, ", ")
	permissions_cookie.Expires = time.Now().Add(time.Hour * 12)
	c.Cookie(permissions_cookie)
	id_cookie := new(fiber.Cookie)
	id_cookie.Name = "user_id"
	id_cookie.Value = login_res.UserID
	id_cookie.Expires = time.Now().Add(time.Hour * 12)
	c.Cookie(id_cookie)
	complete <- true
}

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
		fmt.Println("create userID", c.Cookies("user_id"))
		fmt.Println("create admin", c.Cookies("admin"))
		if complete {
			return c.Status(fiber.StatusCreated).JSON(responses.NewUser(login_res))
		}
	} else {
		login_res, err := user.Login(*user_login)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError(err.Error()))
		}
		go SetCookies(c, login_res, cookies_set)
		complete := <-cookies_set
		fmt.Println("userID", c.Cookies("user_id"))
		fmt.Println("admin", c.Cookies("admin"))
		fmt.Println("permissions", c.Cookies("permissions"))
		if complete {
			return c.Status(fiber.StatusOK).JSON(responses.LoggedIn(login_res))
		}
	}
	c.Response().Header.Add("no-cache", "true")
	return c.Status(fiber.StatusInternalServerError).JSON(responses.ServerError("error setting cookies"))
}

// @id logout-user
// @summary logout user
// @description logs a user out and clears all server side cookies associated with their session
// @tags user, no-cache, auth
// @produce json
// @success 200 {object} responses.NilRes
// @router /auth/logout [post]
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
	c.Response().Header.Add("no-cache", "true")
	return c.Status(fiber.StatusOK).JSON(responses.LoggedOut())
}
