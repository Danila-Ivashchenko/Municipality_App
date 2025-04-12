package entity_to_partition

import (
	"go.uber.org/fx"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/domain/service"
)

type ServiceParams struct {
	fx.In

	EntityToPartitionRepository repository.EntityToPartitionRepository
}

type entityToPartitionService struct {
	ServiceParams
}

func New(params ServiceParams) service.EntityToPartitionService {
	return &entityToPartitionService{
		ServiceParams: params,
	}
}
