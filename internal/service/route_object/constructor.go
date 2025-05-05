package route_object

import (
	"go.uber.org/fx"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/domain/service"
)

type ServiceParams struct {
	fx.In

	RouteObjectRepository repository.RouteObjectRepository
	ObjectService         service.ObjectService
	LocationRepository    repository.LocationRepository
}

type routeObjectService struct {
	ServiceParams
}

func New(params ServiceParams) service.RouteObjectService {
	return &routeObjectService{
		ServiceParams: params,
	}
}
