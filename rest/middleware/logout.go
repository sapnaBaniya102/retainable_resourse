package middleware

import (
	"awesomeProject/modules/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/sujit-baniya/session"
)

func Logout(c *fiber.Ctx) error {
	c.ClearCookie()
	id, err := session.ID(c)
	if err != nil {
		auth.LoggedInBucket.Remove(id)
	}

	return session.DeleteWithDestroy(c, "user_id", "user")
}
