package cmd

import (
	"awesomeProject/app"
	"embed"
	"encoding/json"
	"flag"
	"time"

	"github.com/gofiber/fiber/v2"
	view "github.com/sujit-baniya/fiber-view"
)

var (
	m = flag.Bool("migrate", false, "Update db structure")
	c = flag.String("config", "config.yml", "Read Config")
)

func Execute(migrationFS embed.FS) {
	flag.Parse()
	app.ParseConfig(*c)
	srv := initServices(migrationFS, *m)

	initMiddlewareAndRoutes(srv)

}
func initServer() *fiber.App {
	app.App.Engine = fiber.New(fiber.Config{
		Views:             view.Template(),
		ViewsLayout:       "layouts/landing",
		ReduceMemoryUsage: true,
		AppName:           app.App.Server.Name,
		JSONDecoder:       json.Unmarshal,
		JSONEncoder:       json.Marshal,
	})
}

func initMiddlewareAndRoutes(srv *fiber.App) {
	initStatic(srv)

}
func initServices(migrationFS embed.FS, migrate bool) *fiber.App {

	srv := initServer()

	return srv
}

func initStatic(srv *fiber.App) {
	srv.Static("/", app.App.View.AssetsPath+"/landing", fiber.Static{
		Compress:      true,
		ByteRange:     true,
		CacheDuration: 24 * time.Hour,
	})
	srv.Static("/app/", app.App.View.AssetsPath, fiber.Static{
		Compress:      true,
		ByteRange:     true,
		CacheDuration: 24 * time.Hour,
	})
}
