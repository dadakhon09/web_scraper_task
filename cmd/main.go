package main

import (
	"github.com/dadakhon09/web_scraper_task/api"
	"github.com/dadakhon09/web_scraper_task/config"
	"github.com/dadakhon09/web_scraper_task/pkg/logger"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "task")

	server := api.New(&api.RouterOptions{
		Log:      log,
		Cfg:      &cfg,
	})

	server.Run(cfg.HttpPort)

}
