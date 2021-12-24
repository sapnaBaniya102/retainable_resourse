package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sujit-baniya/session"
)

func IsLoggedIn(c *fiber.Ctx) bool {
	userID := c.Locals("user_id")
	if userID != nil {
		return true
	}
	userID, _ = session.Get(c, "user_id")
	if userID == nil {
		return false
	}
	c.Locals("user_id", userID)
	return true
}
