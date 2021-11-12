package main

import (
	"log"

	"describe.me/config"
	"describe.me/internal/app"
	// Imports for swaggo/swag
	_ "describe.me/internal/objects/transport"
	_ "describe.me/internal/utils/response"

	"github.com/ilyakaznacheev/cleanenv"
)

func main() {
	// Configuration
	var cfg config.Config

	err := cleanenv.ReadConfig("./config/config.yml", &cfg)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(&cfg)
}
