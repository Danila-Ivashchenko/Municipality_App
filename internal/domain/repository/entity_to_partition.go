package repository

import (
	"context"
	"municipality_app/internal/domain/entity"
)

type EntityToPartitionRepository interface {
	Create(ctx context.Context, EntityToPartition *entity.EntityToPartition) error
	Delete(ctx context.Context, EntityToPartition *entity.EntityToPartition) error

	GetByEntityID(ctx context.Context, EntityID int64) ([]entity.EntityToPartition, error)
	GetByPartitionID(ctx context.Context, partitionID int64) ([]entity.EntityToPartition, error)
	GetByEntityIDAndPartitionID(ctx context.Context, EntityID, partitionID int64) (*entity.EntityToPartition, error)
}
