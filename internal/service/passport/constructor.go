package passport

import (
	"go.uber.org/fx"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/domain/service"
)

type ServiceParams struct {
	fx.In

	PassportRepository     repository.PassportRepository
	MunicipalityRepository repository.MunicipalityRepository
}

type passportService struct {
	ServiceParams
}

func New(params ServiceParams) service.PassportService {
	return &passportService{
		ServiceParams: params,
	}
}
