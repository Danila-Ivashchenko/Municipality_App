package region

import (
	"go.uber.org/fx"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/domain/service"
)

type ServiceParams struct {
	fx.In

	RegionRepository repository.RegionRepository
}

type regionService struct {
	ServiceParams
}

func New(params ServiceParams) service.RegionService {
	return &regionService{
		ServiceParams: params,
	}
}
