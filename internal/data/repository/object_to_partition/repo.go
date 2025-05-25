package object_to_partition

import (
	"context"
	"municipality_app/internal/common/sql_handler"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/infrastructure/db"
)

type objectToPartitionRepository struct {
	handler sql_handler.Handler
}

func New(m db.DataBaseManager) repository.ObjectToPartitionRepository {
	repo := &objectToPartitionRepository{
		handler: sql_handler.NewHandler(m.GetDB()),
	}
	return repo
}

func (r *objectToPartitionRepository) Create(ctx context.Context, objectToPartition *entity.ObjectToPartition) error {
	m := newModel(objectToPartition)

	return r.execQuery(ctx, createQuery, m.ObjectID, m.PartitionID)
}

func (r *objectToPartitionRepository) Delete(ctx context.Context, objectToPartition *entity.ObjectToPartition) error {
	m := newModel(objectToPartition)
	return r.execQuery(ctx, deleteQuery, m.ObjectID, m.PartitionID)
}

func (r *objectToPartitionRepository) GetByObjectID(ctx context.Context, objectID int64) ([]entity.ObjectToPartition, error) {
	return r.fetchRowsWithCondition(ctx, "municipality_object_id = $1", objectID)
}

func (r *objectToPartitionRepository) GetByPartitionID(ctx context.Context, partitionID int64) ([]entity.ObjectToPartition, error) {
	return r.fetchRowsWithCondition(ctx, "municipality_passport_partitition_id = $1", partitionID)
}

func (r *objectToPartitionRepository) GetByObjectIDAndPartitionID(ctx context.Context, objectID, partitionID int64) (*entity.ObjectToPartition, error) {
	return r.fetchRowWithCondition(ctx, "municipality_object_id = $1 AND municipality_passport_partitition_id = $2", objectID, partitionID)
}
