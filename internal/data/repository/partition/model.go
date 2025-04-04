package partition

import (
	"database/sql"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
)

type modelPartition struct {
	ID          sql.NullInt64
	ChapterID   sql.NullInt64
	Name        sql.NullString
	Description sql.NullString
	Text        sql.NullString
	OrderNumber sql.NullInt64
}

func newModelPartition(i *entity.Partition) *modelPartition {
	return &modelPartition{
		ID:          sql_common.NewNullInt64(i.ID),
		ChapterID:   sql_common.NewNullInt64(i.ChapterID),
		Name:        sql_common.NewNullString(i.Name),
		Description: sql_common.NewNullString(i.Description),
		Text:        sql_common.NewNullString(i.Text),
		OrderNumber: sql_common.NewNullInt64(int64(i.OrderNumber)),
	}
}

func newModelFromCreateData(i *repository.CreatePartitionData) *modelPartition {
	return &modelPartition{
		ChapterID:   sql_common.NewNullInt64(i.ChapterID),
		Name:        sql_common.NewNullString(i.Name),
		Description: sql_common.NewNullString(i.Description),
		Text:        sql_common.NewNullString(i.Text),
		OrderNumber: sql_common.NewNullInt64(int64(i.OrderNumber)),
	}
}

func (m *modelPartition) convert() *entity.Partition {
	return &entity.Partition{
		ID:          m.ID.Int64,
		ChapterID:   m.ChapterID.Int64,
		Name:        m.Name.String,
		Description: m.Description.String,
		Text:        m.Text.String,
		OrderNumber: uint(m.OrderNumber.Int64),
	}
}
