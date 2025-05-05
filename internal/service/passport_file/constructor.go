package passport_file

import (
	"go.uber.org/fx"
	"municipality_app/internal/common/config"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/domain/service"
)

type ServiceParams struct {
	fx.In

	PassportRepository     repository.PassportRepository
	MunicipalityRepository repository.MunicipalityRepository
	Config                 *config.Config

	PassportFileRepository repository.PassportFileRepository
}

type passportFileService struct {
	ServiceParams
}

func New(params ServiceParams) service.PassportFileService {
	return &passportFileService{
		ServiceParams: params,
	}
}
