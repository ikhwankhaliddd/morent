package main

import (
	"github.com/ikhwankhaleed/morent/config"
	"github.com/ikhwankhaleed/morent/internal/app"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	server := app.NewServer(cfg)
	if err := server.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
