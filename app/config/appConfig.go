package config

import (
	"carApp/app/logging"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	AppConfig *AppConfig
}

// function provider
func New(log logging.ILogger) IConfig {

	// load with godotenv
	if err := godotenv.Load(".env"); err != nil {
		log.LogConsole().Fatal("Error loading .env file")
	}

	cfg := &Config{
		AppConfig: &AppConfig{
			App: &App{
				Port: string(os.Getenv("APP_PORT")),
			},
			Database: &Database{
				Host:     os.Getenv("DATABASE_HOST"),
				Port:     os.Getenv("DATABASE_PORT"),
				User:     os.Getenv("DATABASE_USER"),
				Password: os.Getenv("DATABASE_PASSWORD"),
				Name:     os.Getenv("DATABASE_NAME"),
			},
			Jaeger: &Jaeger{
				Host: os.Getenv("JAEGER_HOST"),
				Port: os.Getenv("JAEGER_PORT"),
			},
		},
	}

	log.LogConsole().Info("success load env")
	return cfg
}

func (c *Config) Config() *AppConfig {
	return c.AppConfig
}

type AppConfig struct {
	App      *App      `json:"app,omitempty"`
	Database *Database `json:"database,omitempty"`
	Jaeger   *Jaeger   `json:"jaeger,omitempty"`
}

type App struct {
	Port string `json:"port,omitempty"`
}

type Database struct {
	Host     string `json:"host,omitempty"`
	Port     string `json:"port,omitempty"`
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
	Name     string `json:"name,omitempty"`
}

type Jaeger struct {
	Host string `json:"host,omitempty"`
	Port string `json:"port,omitempty"`
}
