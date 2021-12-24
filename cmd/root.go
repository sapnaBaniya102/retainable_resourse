package cmd

import (
	"awesomeProject/app"
	"embed"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/goccy/go-json"

	"awesomeProject/rest/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/sujit-baniya/db"
	"github.com/sujit-baniya/db/migrations"
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
	if *m {
		return
	}

	initMiddlewareAndRoutes(srv)
	host := fmt.Sprintf("%s:%s", app.App.Server.Host, app.App.Server.Port)
	log.Fatal(srv.Listen(host))

}

func initServer() *fiber.App {
	app.App.Engine = fiber.New(fiber.Config{
		Views:             view.Template(),
		ReduceMemoryUsage: true,
		AppName:           app.App.Server.Name,
		JSONDecoder:       json.Unmarshal,
		JSONEncoder:       json.Marshal,
	})
	return app.App.Engine
}

func initMiddlewareAndRoutes(srv *fiber.App) {
	initStatic(srv)
	routes.InitLanding(srv)

}
func initServices(migrationFS embed.FS, migrate bool) *fiber.App {
	err := initPreServices()
	if err != nil {
		panic(err)
	}
	initMigration(migrationFS)
	if migrate {
		return nil
	}
	srv := initServer()
	return srv
}

func initPreServices() error {
	view.Default(view.Config{
		Path:      app.App.View.Path,
		Extension: app.App.View.Extension,
	})
	return db.Default(db.Config{
		Driver:   app.App.Database.Driver,
		Host:     app.App.Database.Host,
		Username: app.App.Database.Username,
		Password: app.App.Database.Password,
		DBName:   app.App.Database.DBName,
		Port:     app.App.Database.Port,
	})
}

func initMigration(migrationFS embed.FS) {
	migrations.NewMigration(migrations.Config{
		EmbeddedFS: migrationFS,
		Dir:        app.App.Migrations.Dir,
		TableName:  app.App.Migrations.Table,
	})
	if *m {
		migrations.MigrationMain()
		return
	}
}

func initStatic(srv *fiber.App) {
	srv.Static("/", app.App.View.AssetsPath, fiber.Static{
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
