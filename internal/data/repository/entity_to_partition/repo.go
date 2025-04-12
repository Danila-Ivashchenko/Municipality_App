package entity_to_partition

import (
	"context"
	"database/sql"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/infrastructure/db"
)

type entityToPartitionRepository struct {
	db *sql.DB
}

func New(m db.DataBaseManager) repository.EntityToPartitionRepository {
	repo := &entityToPartitionRepository{
		db: m.GetDB(),
	}
	return repo
}

func (r *entityToPartitionRepository) Create(ctx context.Context, entityToPartition *entity.EntityToPartition) error {
	m := newModel(entityToPartition)

	return r.execQuery(ctx, createQuery, m.EntityID, m.PartitionID)
}

func (r *entityToPartitionRepository) Delete(ctx context.Context, entityToPartition *entity.EntityToPartition) error {
	m := newModel(entityToPartition)
	return r.execQuery(ctx, deleteQuery, m.EntityID, m.PartitionID)
}

func (r *entityToPartitionRepository) GetByEntityID(ctx context.Context, entityID int64) ([]entity.EntityToPartition, error) {
	return r.fetchRowsWithCondition(ctx, "municipality_entity_id = $1", entityID)
}

func (r *entityToPartitionRepository) GetByPartitionID(ctx context.Context, partitionID int64) ([]entity.EntityToPartition, error) {
	return r.fetchRowsWithCondition(ctx, "municipality_passport_partitition_id = $1", partitionID)
}

func (r *entityToPartitionRepository) GetByEntityIDAndPartitionID(ctx context.Context, entityID, partitionID int64) (*entity.EntityToPartition, error) {
	return r.fetchRowWithCondition(ctx, "municipality_entity_id = $1 AND municipality_passport_partitition_id = $2", entityID, partitionID)
}
