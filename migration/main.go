package migration

import (
	"synapsis-challenge/config"

	log "github.com/sirupsen/logrus"
)

func RunMigration(args []string, dbConfig config.DB) {
	migrationSvc := InitMigrationService(dbConfig)

	switch args[0] {
	case "up":
		migrationSvc.Up()
	case "rollback":
		migrationSvc.Rollback()
	default:
		log.Fatal("Invalid migration command")
	}
}

func IsRunMigration(args []string) ([]string, bool) {
	flag := "--migration"

	var flagIndex int
	var found bool
	var migrationArgs []string

	for i, arg := range args {
		if arg == flag {
			flagIndex = i
			found = true
			break
		}
	}

	if found {
		migrationArgs = args[flagIndex+1:]
	}
	return migrationArgs, found
}
