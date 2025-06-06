package route

import (
	"go.uber.org/fx"
	"municipality_app/internal/common/transactor"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/domain/service"
)

type ServiceParams struct {
	fx.In

	RouteObjectService service.RouteObjectService
	RouteRepository    repository.RouteRepository

	Transactor transactor.Transactor
}

type routeService struct {
	ServiceParams
}

func New(params ServiceParams) service.RouteService {
	return &routeService{
		ServiceParams: params,
	}
}
