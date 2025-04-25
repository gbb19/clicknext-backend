package migration

import (
	"clicknext-backend/internal/config"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type Migrator struct {
	migrate *migrate.Migrate
}

func NewMigrator(cfg *config.DatabaseConfig) (*Migrator, error) {
	dsn := getConnectionString(cfg)
	log.Println("Connecting with DSN:", dsn)
	m, err := migrate.New("file://migrations", dsn)
	if err != nil {
		return nil, err
	}

	return &Migrator{
		migrate: m,
	}, nil
}

func (m *Migrator) Up() error {
	if err := m.migrate.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}

func (m *Migrator) Down() error {
	if err := m.migrate.Down(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}

func (m *Migrator) Steps(n int) error {
	if err := m.migrate.Steps(n); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}

func (m *Migrator) DropDatabase() error {
	m.migrate.Drop()
	return nil
}

func getConnectionString(cfg *config.DatabaseConfig) string {
	return "postgres://" + cfg.Username + ":" + cfg.Password + "@" + "localhost" + ":" + cfg.Port + "/" + cfg.DBName + "?sslmode=disable"
}
