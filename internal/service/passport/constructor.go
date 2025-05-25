package passport

import (
	"go.uber.org/fx"
	"municipality_app/internal/common/transactor"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/domain/service"
)

type ServiceParams struct {
	fx.In

	PassportRepository     repository.PassportRepository
	MunicipalityRepository repository.MunicipalityRepository

	Transactor transactor.Transactor
}

type passportService struct {
	ServiceParams
}

func New(params ServiceParams) service.PassportService {
	return &passportService{
		ServiceParams: params,
	}
}
