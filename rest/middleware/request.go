package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func ValidateRequest(c *fiber.Ctx) error {

	return c.Next()
}
