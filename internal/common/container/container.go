package container

import (
	"go.uber.org/fx"
	"municipality_app/internal/common/config"
)

var (
	CommonContainer = fx.Provide(
		config.GetConfig,
	)
)
