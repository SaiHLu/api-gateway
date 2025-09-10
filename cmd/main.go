package main

import (
	"log"

	"github.com/SaiHLu/api-gateway/config"
	_ "github.com/SaiHLu/api-gateway/docs"
	"github.com/SaiHLu/api-gateway/internal/app"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	server := app.NewAppServer(cfg)

	if err := server.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
