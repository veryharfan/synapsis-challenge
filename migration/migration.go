package migration

import (
	"database/sql"
	"errors"
	"synapsis-challenge/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

type MigrationService interface {
	Up() error
	Rollback() error
	Version() (int, bool, error)
}

type migrationService struct {
	driver  database.Driver
	migrate *migrate.Migrate
}

func New(dbConfig config.DB) MigrationService {
	// Connect to the database
	db, err := sql.Open("postgres", dbConfig.ConnUri)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Create a new instance of the PostgreSQL driver
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Failed to create database driver instance: %v", err)
	}

	// Create a new instance of the migrate instance
	m, err := migrate.NewWithDatabaseInstance(
		"file://migration/sql",
		"synapsis",
		driver,
	)
	if err != nil {
		log.Fatalf("Failed to create migrate instance: %v", err)
	}

	return &migrationService{
		driver:  driver,
		migrate: m,
	}

}

func (s *migrationService) Up() error {
	currVersion, _, err := s.Version()
	if err != nil {
		log.Error("Failed get current version err: ", err)
		return err
	}

	log.Infof("Running migration from version: %d", currVersion)
	if err := s.migrate.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Info("No Changes")
			return nil
		}
		log.Error("Failed run migrate err: ", err)
		return err
	}

	currVersion, _, _ = s.Version()
	log.Info("Migration success, current version:", currVersion)
	return nil
}

func (s migrationService) Rollback() error {
	currVersion, _, err := s.Version()
	if err != nil {
		log.Error("Failed get current version err: ", err)
		return err
	}

	log.Infof("Rollingback 1 step from version: %d", currVersion)

	if err := s.migrate.Steps(-1); err != nil {
		log.Errorf("Failed to rollback, err:%v", err)
		return err
	}

	currVersion, _, _ = s.Version()
	log.Infof("Rollback success, current version:%d", currVersion)
	return nil
}

func (s *migrationService) Version() (int, bool, error) {
	currVersion, dirty, err := s.driver.Version()
	if err != nil {
		log.Errorf("Failed to get version:%v", err)
		return 0, false, err
	}
	return currVersion, dirty, nil
}
