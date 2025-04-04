package municipality

import (
	"go.uber.org/fx"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/domain/service"
)

type ServiceParams struct {
	fx.In

	RegionRepository       repository.RegionRepository
	MunicipalityRepository repository.MunicipalityRepository
}

type municipalityService struct {
	ServiceParams
}

func New(params ServiceParams) service.MunicipalityService {
	return &municipalityService{
		ServiceParams: params,
	}
}
