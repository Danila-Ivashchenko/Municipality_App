package object_attribute

import (
	"go.uber.org/fx"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/domain/service"
)

type ServiceParams struct {
	fx.In

	ObjectAttributeRepository      repository.ObjectAttributeRepository
	ObjectAttributeValueRepository repository.ObjectAttributeValueRepository
}

type objectAttributeService struct {
	ServiceParams
}

func New(params ServiceParams) service.ObjectAttributeService {
	return &objectAttributeService{
		ServiceParams: params,
	}
}
