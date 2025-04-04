package service

import (
	"context"
	"municipality_app/internal/domain/entity"
)

type ObjectToPartitionService interface {
	ActualizeToPartition(ctx context.Context, partitionID int64, objectIDs []int64) ([]entity.ObjectToPartition, error)
	GetToPartition(ctx context.Context, partitionID int64) ([]entity.ObjectToPartition, error)
}
