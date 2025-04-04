package container

import (
	"go.uber.org/fx"
	"municipality_app/internal/infrastructure/pgsql"
)

var (
	InfrastructureContainer = fx.Provide(
		pgsql.NewPsqlManager,
	)
)
