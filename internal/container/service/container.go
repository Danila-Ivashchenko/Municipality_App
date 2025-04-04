package service

import (
	"go.uber.org/fx"
	"municipality_app/internal/service/chapter"
	"municipality_app/internal/service/municipality"
	"municipality_app/internal/service/object"
	"municipality_app/internal/service/object_attribute"
	"municipality_app/internal/service/object_ex"
	"municipality_app/internal/service/object_template"
	"municipality_app/internal/service/object_to_partition"
	"municipality_app/internal/service/object_type"
	"municipality_app/internal/service/partition"
	"municipality_app/internal/service/passport"
	"municipality_app/internal/service/passport_ex"
	"municipality_app/internal/service/passport_file"
	"municipality_app/internal/service/region"
	"municipality_app/internal/service/user"
	"municipality_app/internal/service/user_auth"
)

var (
	ServiceContainer = fx.Provide(
		user.New,
		user_auth.New,
		region.New,
		municipality.New,
		passport.New,
		chapter.New,
		partition.New,
		passport_ex.New,
		object_type.New,
		object.New,
		object_template.New,
		object_attribute.New,
		object_ex.New,
		object_to_partition.New,
		passport_file.New,
	)
)
