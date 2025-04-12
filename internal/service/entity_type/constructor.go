package entity_type

import (
	"go.uber.org/fx"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/domain/service"
)

type ServiceParams struct {
	fx.In

	EntityTypeRepository repository.EntityTypeRepository
}

type entityTypeService struct {
	ServiceParams
}

func New(params ServiceParams) service.EntityTypeService {
	return &entityTypeService{
		ServiceParams: params,
	}
}
