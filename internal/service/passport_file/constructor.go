package passport_file

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

type passportFileService struct {
	ServiceParams
}

func New(params ServiceParams) service.PassportFileService {
	return &passportFileService{
		ServiceParams: params,
	}
}
