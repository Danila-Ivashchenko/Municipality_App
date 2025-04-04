package chapter

import (
	"database/sql"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
)

type Chapter struct {
	ID          int64
	PassportID  int64
	Name        string
	Description string
	Text        string
	OrderNumber uint
}

type modelChapter struct {
	ID          sql.NullInt64
	PassportID  sql.NullInt64
	Name        sql.NullString
	Description sql.NullString
	Text        sql.NullString
	OrderNumber sql.NullInt64
}

func newModelChapter(i *entity.Chapter) *modelChapter {
	return &modelChapter{
		ID:          sql_common.NewNullInt64(i.ID),
		PassportID:  sql_common.NewNullInt64(i.PassportID),
		Name:        sql_common.NewNullString(i.Name),
		Description: sql_common.NewNullString(i.Description),
		Text:        sql_common.NewNullString(i.Text),
		OrderNumber: sql_common.NewNullInt64(int64(i.OrderNumber)),
	}
}

func newModelFromCreateData(i *repository.CreateChapterData) *modelChapter {
	return &modelChapter{
		PassportID:  sql_common.NewNullInt64(i.PassportID),
		Name:        sql_common.NewNullString(i.Name),
		Description: sql_common.NewNullString(i.Description),
		Text:        sql_common.NewNullString(i.Text),
		OrderNumber: sql_common.NewNullInt64(int64(i.OrderNumber)),
	}
}

func (m *modelChapter) convert() *entity.Chapter {
	return &entity.Chapter{
		ID:          m.ID.Int64,
		PassportID:  m.PassportID.Int64,
		Name:        m.Name.String,
		Description: m.Description.String,
		Text:        m.Text.String,
		OrderNumber: uint(m.OrderNumber.Int64),
	}
}
