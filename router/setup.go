package router

import "github.com/gofiber/fiber/v2"

var USER fiber.Router

// SetUpRoutes setup all the routes
func SetUpRoutes(app *fiber.App) {
	api := app.Group("/api")
	USER = api.Group("/user")
	SetupUserRoutes()
}
