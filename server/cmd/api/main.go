package main

import (
	"github.com/shojib116/auditflow-api/config"
	"github.com/shojib116/auditflow-api/internal/bootstrap"
	"log"
)

func main() {
	cfg := config.GetConfig()

	app, err := bootstrap.NewApp(cfg)
	if err != nil {
		log.Fatalf("failed to initialize app: %s", err)
	}

	if err := app.Run(); err != nil {
		log.Fatalf("app exited with error: %s", err)
	}
}
