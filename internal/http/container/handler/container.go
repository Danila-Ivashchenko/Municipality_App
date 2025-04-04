package handler

import (
	"go.uber.org/fx"
	"municipality_app/internal/http/handler/chapter"
	"municipality_app/internal/http/handler/municipality"
	"municipality_app/internal/http/handler/object"
	"municipality_app/internal/http/handler/object_type"
	"municipality_app/internal/http/handler/partition"
	"municipality_app/internal/http/handler/passport"
	"municipality_app/internal/http/handler/region"
	"municipality_app/internal/http/handler/user"
)

var (
	HandlerContainer = fx.Provide(
		user.New,
		region.New,
		municipality.New,
		passport.New,
		object_type.New,
		chapter.New,
		partition.New,
		object.New,
	)
)
