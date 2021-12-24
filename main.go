package main

import (
	"awesomeProject/cmd"
	"embed"

	_ "github.com/lib/pq"
)

//go:embed migrations/*
var migrationFS embed.FS

func main() {
	cmd.Execute(migrationFS)
}
