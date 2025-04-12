package entity_to_partition

import (
	"database/sql"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
)

type model struct {
	EntityID    sql.NullInt64
	PartitionID sql.NullInt64
}

func (m *model) convert() *entity.EntityToPartition {
	return &entity.EntityToPartition{
		EntityID:    m.EntityID.Int64,
		PartitionID: m.PartitionID.Int64,
	}
}

func newModel(data *entity.EntityToPartition) *model {
	if data == nil {
		return nil
	}

	return &model{
		EntityID:    sql_common.NewNullInt64(data.EntityID),
		PartitionID: sql_common.NewNullInt64(data.PartitionID),
	}
}
