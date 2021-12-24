package middleware

import (
	"awesomeProject/modules/auth"
	"awesomeProject/modules/auth/requests"
	"fmt"
	"regexp"

	"github.com/gofiber/fiber/v2"
	"github.com/sujit-baniya/flash"
)

var rxEmail = regexp.MustCompile(".+@.+\\..+")

func ValidateRegisterPost(redirectTo string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var register requests.Register
		if err := c.BodyParser(&register); err != nil {
			fmt.Println(err)
			return flash.WithError(c, fiber.Map{
				"message": err.Error(),
				"old":     register,
			}).Redirect(redirectTo)
		}
		c.Locals("register", register)
		c.Locals("email", register.Email)
		return c.Next()
	}
}

func ValidateEmail(redirectTo string) func(c *fiber.Ctx) error {

	return func(c *fiber.Ctx) error {
		var email requests.EmailCode
		fmt.Println("123")
		if err := c.BodyParser(&email); err != nil {
			fmt.Println(err)
			return flash.WithError(c, fiber.Map{
				"message": err.Error(),
				"old":     email,
			}).Redirect(redirectTo)
		}
		fmt.Println(email)
		UserEmail := c.Locals("email").(string)
		match := rxEmail.Match([]byte(UserEmail))
		if UserEmail == "" || !match {

			mp := fiber.Map{
				"error":   true,
				"message": "Please enter a valid email address",
			}
			return flash.WithError(c, mp).Redirect(redirectTo)

		}
		loginResponse, err := auth.GetUserWithProfileByEmail(UserEmail)
		if err != nil {
			mp := fiber.Map{
				"message": err.Error(),
			}
			return flash.WithError(c, mp).Redirect(redirectTo)
		}
		c.Locals("code", email.Code)
		c.Locals("login_response", loginResponse)
		return c.Next()
	}

}
