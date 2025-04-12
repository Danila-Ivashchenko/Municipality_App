package service

import (
	"context"
	"municipality_app/internal/domain/entity"
)

type EntityToPartitionService interface {
	ActualizeToPartition(ctx context.Context, partitionID int64, EntityIDs []int64) ([]entity.EntityToPartition, error)
	GetToPartition(ctx context.Context, partitionID int64) ([]entity.EntityToPartition, error)
}
