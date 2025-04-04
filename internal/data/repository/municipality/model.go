package municipality

import (
	"database/sql"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
)

type municipalityModel struct {
	ID        sql.NullInt64
	Name      sql.NullString
	RegionID  sql.NullInt64
	IsHidden  sql.NullBool
	CreatedAt sql.NullTime
}

func (m *municipalityModel) convert() *entity.Municipality {
	return &entity.Municipality{
		ID:        m.ID.Int64,
		Name:      m.Name.String,
		RegionID:  m.RegionID.Int64,
		IsHidden:  m.IsHidden.Bool,
		CreatedAt: m.CreatedAt.Time,
	}
}

func newMunicipalityModel(i *entity.Municipality) *municipalityModel {
	return &municipalityModel{
		ID:       sql_common.NewNullInt64(i.ID),
		Name:     sql_common.NewNullString(i.Name),
		RegionID: sql_common.NewNullInt64(i.RegionID),
		IsHidden: sql_common.NewNullBool(i.IsHidden),
	}
}

func newMunicipalityModelFromCreateData(data *repository.CreateMunicipalityData) *municipalityModel {
	return &municipalityModel{
		Name:     sql_common.NewNullString(data.Name),
		RegionID: sql_common.NewNullInt64(data.RegionID),
	}
}
