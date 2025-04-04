package object_to_partition

import (
	"context"
	"municipality_app/internal/domain/entity"
)

func (svc *objectToPartitionService) GetToPartition(ctx context.Context, partitionID int64) ([]entity.ObjectToPartition, error) {
	return svc.ObjectToPartitionRepository.GetByPartitionID(ctx, partitionID)
}
