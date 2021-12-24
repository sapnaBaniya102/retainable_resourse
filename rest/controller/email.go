package controller

import (
	"awesomeProject/modules/auth/requests"
	"crypto/rand"
	"fmt"

	"github.com/gofiber/fiber/v2"
	view "github.com/sujit-baniya/fiber-view"
	"github.com/sujit-baniya/flash"
)

func EmailCode(c *fiber.Ctx) error {
	data := flash.Get(c)
	fmt.Println(view.Template())
	return view.Render(c, "validation", data)
}

func VerifyRegisteredEmail(c *fiber.Ctx) error {
	EmailCode := c.Locals("code").(requests.EmailCode)
	var p, _ = rand.Prime(rand.Reader, 16)

	if EmailCode.Code.String() == p.String() {
		mp := fiber.Map{
			"success": true,
			"message": "success",
		}

		return flash.WithSuccess(c, mp).Redirect("/success")
	}
	error := fiber.Map{
		"error":   true,
		"message": "error",
	}
	return flash.WithError(c, error).Redirect("/verify_email")
}
