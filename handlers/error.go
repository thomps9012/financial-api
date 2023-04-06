package handlers

import (
	"financial-api/responses"

	"github.com/gofiber/fiber/v2"
)

func LogError(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).JSON(responses.LogError())
}
