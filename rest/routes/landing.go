package routes

import (
	controllers "awesomeProject/rest/controller"
	middlewares "awesomeProject/rest/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func InitLanding(r fiber.Router) {
	LandingRoutes(r)
	AuthRoutes(r)
}

func LandingRoutes(app fiber.Router) {
	app.Use(cors.New())
	app.Use(middlewares.ValidateRequest)

	app.Get("/index",
		controllers.RegisterGet)
	app.Get("/register",
		middlewares.RedirectOnLoggedInUser("/app"),
		controllers.RegisterGet,
	)

	app.Get("/verify_email",
		middlewares.RedirectOnLoggedInUser("/app"),
		controllers.EmailCode,
	)
	app.Get("/success", func(c *fiber.Ctx) error {
		return c.SendString("success")
	})

}
