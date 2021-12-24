package routes

import (
	controllers "awesomeProject/rest/controller"
	middlewares "awesomeProject/rest/middleware"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(r fiber.Router) {
	r.Post("/register",
		middlewares.ValidateRegisterPost("/register"),
		controllers.RegisterPost,
	)

	r.Post("/verify_email",
		func(c *fiber.Ctx) error {
			return c.SendString("success")
		},
		middlewares.ValidateEmail("/verify_email"),
		controllers.VerifyRegisteredEmail,
	)

}
