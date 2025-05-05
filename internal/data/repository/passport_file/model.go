package passport_file

import (
	"database/sql"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
)

type passportFileModel struct {
	ID         sql.NullInt64
	Path       sql.NullString
	PassportID sql.NullInt64
	FileName   sql.NullString
	CreateAt   sql.NullTime
}

func (m *passportFileModel) convert() *entity.PassportFile {
	return &entity.PassportFile{
		ID:         m.ID.Int64,
		FileName:   m.FileName.String,
		PassportID: m.PassportID.Int64,
		Path:       m.Path.String,
		CreateAt:   m.CreateAt.Time,
	}
}

func newPassportFileModel(i *entity.PassportFile) *passportFileModel {
	return &passportFileModel{
		ID:         sql_common.NewNullInt64(i.ID),
		FileName:   sql_common.NewNullString(i.FileName),
		PassportID: sql_common.NewNullInt64(i.PassportID),
		Path:       sql_common.NewNullString(i.Path),
		CreateAt:   sql_common.NewNullTime(i.CreateAt),
	}
}
