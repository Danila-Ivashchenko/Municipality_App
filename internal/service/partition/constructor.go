package partition

import (
	"go.uber.org/fx"
	"municipality_app/internal/common/transactor"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/domain/service"
)

type ServiceParams struct {
	fx.In

	PartitionRepository      repository.PartitionRepository
	ObjectToPartitionService service.ObjectToPartitionService
	ObjectService            service.ObjectService
	ObjectTemplateService    service.ObjectTemplateService

	EntityToPartitionService service.EntityToPartitionService
	EntityService            service.EntityService
	EntityTemplateService    service.EntityTemplateService
	RouteService             service.RouteService

	Transactor transactor.Transactor
}

type partitionService struct {
	ServiceParams
}

func New(params ServiceParams) service.PartitionService {
	return &partitionService{
		ServiceParams: params,
	}
}
