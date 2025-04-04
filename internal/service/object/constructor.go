package object

import (
	"go.uber.org/fx"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/domain/service"
)

type ServiceParams struct {
	fx.In

	ObjectRepository       repository.ObjectRepository
	LocationRepository     repository.LocationRepository
	ObjectAttributeService service.ObjectAttributeService
}

type objectService struct {
	ServiceParams
}

func New(params ServiceParams) service.ObjectService {
	return &objectService{
		ServiceParams: params,
	}
}
