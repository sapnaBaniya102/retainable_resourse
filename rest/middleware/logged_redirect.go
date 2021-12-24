package middleware

import (
	"awesomeProject/modules/auth"

	"github.com/gofiber/fiber/v2"
)

func RedirectOnLoggedInUser(redirectTo string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if auth.IsLoggedIn(c) {
			return c.Redirect(redirectTo)
		}
		return c.Next()
	}
}
