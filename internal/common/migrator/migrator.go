package migrator

import (
	"database/sql"
	migrate "github.com/rubenv/sql-migrate"
	"log/slog"
	"municipality_app/internal/common/config"
	"municipality_app/internal/infrastructure/db"
)

type BaseMigrator struct {
	migrationsPath string
	db             *sql.DB
}

func NewMigrator(cfg *config.Config, m db.DataBaseManager) *BaseMigrator {
	return &BaseMigrator{
		migrationsPath: cfg.GetMigrationsPath(),
		db:             m.GetDB(),
	}
}

func (m *BaseMigrator) Up() error {
	migrations := &migrate.FileMigrationSource{
		Dir: m.migrationsPath,
	}

	n, err := migrate.Exec(m.db, "postgres", migrations, migrate.Up)
	if err != nil {
		return err
	}

	slog.Info("Applied %d migrations!", n)
	return nil
}
