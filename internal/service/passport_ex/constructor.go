package passport_ex

import (
	"go.uber.org/fx"
	"municipality_app/internal/domain/service"
)

type ServiceParams struct {
	fx.In

	PassportService          service.PassportService
	ChapterService           service.ChapterService
	PartitionService         service.PartitionService
	ObjectExService          service.ObjectExService
	ObjectService            service.ObjectService
	ObjectToPartitionService service.ObjectToPartitionService
}

type passportExService struct {
	ServiceParams
}

func New(params ServiceParams) service.PassportExService {
	return &passportExService{
		ServiceParams: params,
	}
}
