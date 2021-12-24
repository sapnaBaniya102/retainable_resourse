package app

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sujit-baniya/config"
)

type Config struct {
	Engine *fiber.App

	Cache      Cache      `json:"cache" yaml:"cache"`
	Database   Database   `json:"database" yaml:"database"`
	Migrations Migrations `json:"migrations" yaml:"migrations"`
	View       View       `json:"view" yaml:"view"`
	Server     Server     `json:"server" yaml:"server"`
}

var App *Config

func ParseConfig(conf string) {
	var cfg Config
	if err := config.ReadConfig(conf, &cfg); err != nil {
		log.Fatalln(err)
	}
	App = &cfg
}
