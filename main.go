package main

import (
	"awesomeProject/router"
	"awesomeProject/util"
	"embed"
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/lib/pq"
	"github.com/sujit-baniya/db"
	"github.com/sujit-baniya/db/migrations"
	view "github.com/sujit-baniya/fiber-view"
	"github.com/sujit-baniya/flash"
)

func CreateServer() *fiber.App {
	view.Default(view.Config{
		Path:      "./web/views",
		Extension: ".html",
		Global:    []string{"auth"},
	})
	app := fiber.New(fiber.Config{
		Views: view.Template(),
	})
	return app

}

//go:embed migrations/*
var migrationFS embed.FS
var (
	m = flag.Bool("migrate", false, "Update db structure")
)

func main() {
	flag.Parse()
	err := db.Default(db.Config{
		Driver:   "postgres",
		Host:     "localhost",
		Username: "postgres",
		Password: "root",
		DBName:   "gotest",
		Port:     5432,
	})
	if err != nil {
		panic(err)
	}
	initMigration()
	app := CreateServer()
	app.Static("/", "./web/dist")
	app.Get("/", func(c *fiber.Ctx) error {
		data := flash.Get(c)
		return view.Render(c, "login", data)
	})
	app.Get("/register", func(c *fiber.Ctx) error {
		data := flash.Get(c)
		return view.Render(c, "app/register", data)
	})
	app.Post("/register", util.Validation)
	app.Get("/validation", func(c *fiber.Ctx) error {
		data1 := flash.Get(c)
		return view.Render(c, "app/validation", data1)
	})
	app.Post("/validation", util.CodeVerification)
	app.Get("/success", func(c *fiber.Ctx) error {
		return c.SendString("success")
	})
	app.Use(cors.New())

	router.SetUpRoutes(app)
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})
	log.Fatal(app.Listen(":8001"))
}
func initMigration() {
	migrations.NewMigration(migrations.Config{
		EmbeddedFS: migrationFS,
		Dir:        "migrations",
		TableName:  "test_migration",
	})
	if *m {
		migrations.MigrationMain()
		return
	}
}
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
