package entity_to_partition

import (
	"context"
	"municipality_app/internal/domain/entity"
)

func (svc *entityToPartitionService) GetToPartition(ctx context.Context, partitionID int64) ([]entity.EntityToPartition, error) {
	return svc.EntityToPartitionRepository.GetByPartitionID(ctx, partitionID)
}
