package object_to_partition

import (
	"database/sql"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
)

type model struct {
	ObjectID    sql.NullInt64
	PartitionID sql.NullInt64
}

func (m *model) convert() *entity.ObjectToPartition {
	return &entity.ObjectToPartition{
		ObjectID:    m.ObjectID.Int64,
		PartitionID: m.PartitionID.Int64,
	}
}

func newModel(data *entity.ObjectToPartition) *model {
	if data == nil {
		return nil
	}

	return &model{
		ObjectID:    sql_common.NewNullInt64(data.ObjectID),
		PartitionID: sql_common.NewNullInt64(data.PartitionID),
	}
}
