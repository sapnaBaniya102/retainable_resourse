package database

import (
	"embed"
	"flag"

	"github.com/sujit-baniya/db"
	"github.com/sujit-baniya/db/migrations"
)

//go:embed migrations/*
var migrationFS embed.FS

var (
	m = flag.Bool("migrate", false, "Update db structure")
)

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
func Setup() {
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

}
