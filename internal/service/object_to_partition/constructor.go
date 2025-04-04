package object_to_partition

import (
	"go.uber.org/fx"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/domain/service"
)

type ServiceParams struct {
	fx.In

	ObjectToPartitionRepository repository.ObjectToPartitionRepository
}

type objectToPartitionService struct {
	ServiceParams
}

func New(params ServiceParams) service.ObjectToPartitionService {
	return &objectToPartitionService{
		ServiceParams: params,
	}
}
