package entity_attribute

import (
	"go.uber.org/fx"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/domain/service"
)

type ServiceParams struct {
	fx.In

	EntityAttributeRepository      repository.EntityAttributeRepository
	EntityAttributeValueRepository repository.EntityAttributeValueRepository
}

type entityAttributeService struct {
	ServiceParams
}

func New(params ServiceParams) service.EntityAttributeService {
	return &entityAttributeService{
		ServiceParams: params,
	}
}
