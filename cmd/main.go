package main

import (
	"log"
	"log/slog"

	"github.com/MarcosMorelli/medication-api/internal/config"
	"github.com/MarcosMorelli/medication-api/internal/database"
	"github.com/MarcosMorelli/medication-api/internal/logger"
	"github.com/MarcosMorelli/medication-api/internal/server"
)

func main() {
	config := config.NewConfig()
	logger.Init(config)
	database.Init(config)

	slog.Info("Init", "config", config)

	server := server.NewServer(config.Port)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
