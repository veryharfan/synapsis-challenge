package main

import (
	"context"
	"os"
	"synapsis-challenge/config"
	"synapsis-challenge/migration"

	log "github.com/sirupsen/logrus"
)

func main() {
	cfg := config.InitConfig(context.Background())

	args := os.Args
	if len(args) < 2 {
		log.Fatal("Missing args. args: [up | rollback]")
	}

	migrationSvc := migration.New(cfg.DB)

	switch args[1] {
	case "up":
		migrationSvc.Up()
	case "rollback":
		migrationSvc.Rollback()
	default:
		log.Fatal("Invalid migration command")
	}
}
