package repository

import (
	"context"
	"municipality_app/internal/domain/entity"
)

type ObjectToPartitionRepository interface {
	Create(ctx context.Context, objectToPartition *entity.ObjectToPartition) error
	Delete(ctx context.Context, objectToPartition *entity.ObjectToPartition) error

	GetByObjectID(ctx context.Context, objectID int64) ([]entity.ObjectToPartition, error)
	GetByPartitionID(ctx context.Context, partitionID int64) ([]entity.ObjectToPartition, error)
	GetByObjectIDAndPartitionID(ctx context.Context, objectID, partitionID int64) (*entity.ObjectToPartition, error)
}
