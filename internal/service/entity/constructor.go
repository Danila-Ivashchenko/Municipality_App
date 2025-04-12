package entity

import (
	"go.uber.org/fx"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/domain/service"
)

type ServiceParams struct {
	fx.In

	EntityRepository       repository.EntityRepository
	LocationRepository     repository.LocationRepository
	EntityAttributeService service.EntityAttributeService
}

type entityService struct {
	ServiceParams
}

func New(params ServiceParams) service.EntityService {
	return &entityService{
		ServiceParams: params,
	}
}
