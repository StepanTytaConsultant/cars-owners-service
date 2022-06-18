package migrate

import (
	"github.com/pkg/errors"
	migrate "github.com/rubenv/sql-migrate"

	"github.com/cars-owners-service/assets"
	"github.com/cars-owners-service/internal/config"
)

const (
	Up   = "up"
	Down = "down"
)

// migrateUp Migrates database up
func migrateUp(cfg config.Config) (int, error) {
	rawDB, err := cfg.DB().DB()
	if err != nil {
		return 0, errors.Wrap(err, "failed to get raw db")
	}

	applied, err := migrate.Exec(rawDB, cfg.Driver(), assets.Migrations, migrate.Up)

	if err != nil {
		return 0, errors.Wrap(err, "failed to apply migrations")
	}

	return applied, nil
}

// migrateDown Migrates database down
func migrateDown(cfg config.Config) (int, error) {
	rawDB, err := cfg.DB().DB()
	if err != nil {
		return 0, errors.Wrap(err, "failed to get raw db")
	}

	applied, err := migrate.Exec(rawDB, cfg.Driver(), assets.Migrations, migrate.Down)
	if err != nil {
		return 0, errors.Wrap(err, "failed to apply migrations")
	}
	return applied, nil
}

func Migrate(cfg config.Config, direction string) error {
	migrator := migrateUp
	switch direction {
	case Down:
		migrator = migrateDown
	}
	migrationsCount, err := migrator(cfg)
	if err != nil {
		return err
	}
	cfg.Logging().WithField("applied", migrationsCount).Info("Migrations applied")
	return nil
}
