package controller

import (
	"awesomeProject/modules/auth/requests"
	"awesomeProject/util"
	"fmt"

	"github.com/gofiber/fiber/v2"
	view "github.com/sujit-baniya/fiber-view"
	"github.com/sujit-baniya/flash"
)

func RegisterGet(c *fiber.Ctx) error {
	data := flash.Get(c)
	fmt.Println(view.Template())
	return view.Render(c, "register", data)
}

func RegisterPost(c *fiber.Ctx) error {
	register := c.Locals("register").(requests.Register)

	user, err := register.Signup()
	if err != nil {
		return c.JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
	go util.SendEmail(c, user.Email)
	mp := fiber.Map{
		"success": true,
		"message": "success",
	}
	return flash.WithSuccess(c, mp).Redirect("/verify_email")
}
