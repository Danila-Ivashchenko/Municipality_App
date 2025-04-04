package middleware

import (
	"go.uber.org/fx"
	"municipality_app/internal/http/middleware/auth"
	"municipality_app/internal/http/middleware/mun"
	"municipality_app/internal/http/middleware/municipality_passport"
)

var (
	MiddlewareContainer = fx.Provide(
		auth.New,
		mun.New,
		municipality_passport.New,
	)
)
