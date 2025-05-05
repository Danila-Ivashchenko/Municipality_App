package repository

import (
	"go.uber.org/fx"
	"municipality_app/internal/data/repository/chapter"
	"municipality_app/internal/data/repository/entity"
	"municipality_app/internal/data/repository/entity_attribute"
	"municipality_app/internal/data/repository/entity_attribute_value"
	"municipality_app/internal/data/repository/entity_template"
	"municipality_app/internal/data/repository/entity_to_partition"
	"municipality_app/internal/data/repository/entity_type"
	"municipality_app/internal/data/repository/location"
	"municipality_app/internal/data/repository/municipality"
	"municipality_app/internal/data/repository/object"
	"municipality_app/internal/data/repository/object_attribute"
	"municipality_app/internal/data/repository/object_attribute_value"
	"municipality_app/internal/data/repository/object_template"
	"municipality_app/internal/data/repository/object_to_partition"
	"municipality_app/internal/data/repository/object_type"
	"municipality_app/internal/data/repository/partition"
	"municipality_app/internal/data/repository/passport"
	"municipality_app/internal/data/repository/passport_file"
	"municipality_app/internal/data/repository/region"
	"municipality_app/internal/data/repository/route"
	"municipality_app/internal/data/repository/route_object"
	"municipality_app/internal/data/repository/user"
	"municipality_app/internal/data/repository/user_auth_token"
)

var (
	RepositoryContainer = fx.Provide(
		user.New,
		user_auth_token.New,
		region.New,
		municipality.New,
		passport.New,
		partition.New,
		chapter.New,
		object_type.New,
		object_template.New,
		object.New,
		location.New,
		object_attribute.New,
		object_attribute_value.New,
		object_to_partition.New,

		entity_type.New,
		entity_template.New,
		entity.New,
		entity_attribute.New,
		entity_attribute_value.New,
		entity_to_partition.New,

		route.New,
		route_object.New,

		passport_file.New,
	)
)
