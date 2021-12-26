package auth

import (
	"awesomeProject/modules/auth/models"

	"github.com/gofiber/fiber/v2"
	"github.com/sujit-baniya/session"
)

var LoggedInBucket models.LoginBucket

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
