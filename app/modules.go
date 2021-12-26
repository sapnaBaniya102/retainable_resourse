package app

import "github.com/gofiber/fiber"

type Server struct {
	*fiber.App
	Name       string `mapstructure:"APP_NAME" yaml:"name" env:"APP_NAME" env-default:"iSend.to"`
	Version    string `mapstructure:"APP_VERSION" yaml:"version" env:"APP_VERSION" env-default:"dev"`
	Mode       string `mapstructure:"APP_MODE" yaml:"mode" env:"APP_MODE" env-default:"app"`
	Env        string `mapstructure:"APP_ENV" yaml:"env" env:"APP_ENV" env-default:"dev"`
	Key        string `mapstructure:"APP_KEY" yaml:"key" env:"APP_KEY" env-default:"abcdefghijklmnopqrstuvwxyz012345"`
	URL        string `mapstructure:"APP_URL" yaml:"url" env:"APP_URL" env-default:"http://localhost:8080"`
	Host       string `mapstructure:"APP_HOST" yaml:"host" env:"APP_HOST" env-default:"localhost"`
	Port       string `mapstructure:"APP_PORT" yaml:"port" env:"APP_PORT" env-default:"8080"`
	Currency   string `mapstructure:"APP_CURRENCY" yaml:"currency" env:"APP_CURRENCY" env-default:"USD"`
	Path       string `mapstructure:"APP_PATH" yaml:"path" env:"APP_PATH"`
	ExecPath   string `mapstructure:"EXEC_PATH" yaml:"exec_path" env:"EXEC_PATH" env-default:"false"`
	UploadPath string `mapstructure:"UPLOAD_PATH" yaml:"upload_path" env:"UPLOAD_PATH" env-default:"uploads"`
	Debug      bool   `mapstructure:"APP_DEBUG" yaml:"debug" env:"APP_DEBUG" env-default:"true"`
}

type View struct {
	Path       string `yaml:"path" env-default:"web/views/app"`
	Extension  string `yaml:"extension" env-default:".html"`
	AssetsPath string `yaml:"assets_path" env-default:"web/dist"`
}
type Cache struct {
	Driver      string `yaml:"driver" env:"CACHE_DRIVER" env-default:"redis"`
	Name        string `yaml:"name" env:"CACHE_NAME" env-default:"redis-cache"`
	Host        string `yaml:"host" env:"CACHE_HOST" env-default:"localhost"`
	Password    string `yaml:"password" env:"CACHE_PASSWORD"`
	DB          string `yaml:"db" env:"CACHE_DB" env-default:"0"`
	Concurrency string `yaml:"concurrency" env:"CACHE_CONCURRENCY" env-default:"100"`
	Port        int    `yaml:"port" env:"CACHE_PORT" env-default:"6379"`
}
type Migrations struct {
	Dir   string `yaml:"dir" env:"MIGRATIONS_DR" env-default:"migrations"`
	Table string `yaml:"table" env:"TABLE" env-default:"migrations"`
}
type Database struct {
	Driver     string `yaml:"driver" env:"DB_DRIVER" env-default:"postgres"`
	Host       string `yaml:"host" env:"DB_HOST" env-default:"localhost"`
	Username   string `yaml:"username" env:"DB_USER" env-default:"postgres"`
	Password   string `yaml:"password" env:"DB_PASS" env-default:"root"`
	DBName     string `yaml:"db_name" env:"DB_NAME" env-default:"test"`
	Port       int    `yaml:"port" env:"DB_PORT" env-default:"5432"`
	MaxOpenCon int    `yaml:"connections" env:"DB_CONNECTIONS" env-default:"100"`
	MaxIdleCon int    `yaml:"idle_connections" env:"DB_IDLE_CONNECTIONS" env-default:"80"`
}
