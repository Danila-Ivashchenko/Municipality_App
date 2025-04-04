package router

import (
	"go.uber.org/fx"
	"municipality_app/internal/http/router"
)

var (
	RouterContainer = fx.Provide(
		router.New,
	)
)
