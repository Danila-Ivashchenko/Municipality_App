package container

import (
	"go.uber.org/fx"
	"municipality_app/internal/common/config"
	"municipality_app/internal/common/migrator"
	"municipality_app/internal/common/transactor/sql_transactor"
)

var (
	CommonContainer = fx.Provide(
		config.GetConfig,
		sql_transactor.NewTransactor,
		migrator.NewMigrator,
	)
)
