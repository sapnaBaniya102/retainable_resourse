package controller

import (
	middlewares "awesomeProject/rest/middleware"
	"sync"

	"github.com/gofiber/fiber/v2"
	view "github.com/sujit-baniya/fiber-view"
	"github.com/sujit-baniya/flash"
)

var mu sync.RWMutex

func LoginGet(c *fiber.Ctx) error {
	mu.Lock()
	defer mu.Unlock()
	data := flash.Get(c)
	data["title"] = "Login | "
	return view.Render(c, "login", view.Append(c, data))
}

// 	data := flash.Get(c)
// 	fmt.Println(view.Template())
// 	return view.Render(c, "login", data)
// }

func LoginPost(c *fiber.Ctx) error { //nolint:wsl
	rememberMe := c.Locals("remember_me").(bool)
	err := middlewares.Login(c, rememberMe) //nolint:wsl
	if err != nil {
		return flash.WithError(c, view.Append(c, fiber.Map{
			"error":   true,
			"message": err.Error(),
		})).Redirect("/login#/")
	}
	return c.Redirect("/app")

}

func LogoutPost(c *fiber.Ctx) error { //nolint:nolintlint,wsl
	if middlewares.IsLoggedIn(c) {

	}
	c.Set("X-DNS-Prefetch-Control", "off")
	c.Set("Pragma", "no-cache")
	c.Set("Expires", "Fri, 01 Jan 1990 00:00:00 GMT")
	c.Set("Cache-Control", "no-cache, must-revalidate, no-store, max-age=0, private")
	return c.Redirect("/login")
}
