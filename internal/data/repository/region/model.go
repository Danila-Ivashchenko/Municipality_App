package region

import (
	"database/sql"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
)

type regionModel struct {
	ID   sql.NullInt64
	Name sql.NullString
	Code sql.NullString
}

func (m *regionModel) convert() *entity.Region {
	return &entity.Region{
		ID:   m.ID.Int64,
		Name: m.Name.String,
		Code: m.Code.String,
	}
}

func newRegionModel(data *repository.CreateRegionData) *regionModel {
	if data == nil {
		return nil
	}

	return &regionModel{
		Name: sql_common.NewNullString(data.Name),
		Code: sql_common.NewNullString(data.Code),
	}
}
