package object_template

import (
	"database/sql"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
)

type modelObjectTemplate struct {
	ID             sql.NullInt64
	Name           sql.NullString
	ObjectTypeID   sql.NullInt64
	MunicipalityID sql.NullInt64
}

type ObjectTemplate struct {
	ID           int64
	Name         string
	ObjectTypeID int64
	PartitionID  int64
}

func (m *modelObjectTemplate) convert() *entity.ObjectTemplate {
	return &entity.ObjectTemplate{
		ID:             m.ID.Int64,
		Name:           m.Name.String,
		ObjectTypeID:   m.ObjectTypeID.Int64,
		MunicipalityID: m.MunicipalityID.Int64,
	}
}

func newModelObjectTemplate(i *entity.ObjectTemplate) *modelObjectTemplate {
	return &modelObjectTemplate{
		ID:             sql_common.NewNullInt64(i.ID),
		Name:           sql_common.NewNullString(i.Name),
		ObjectTypeID:   sql_common.NewNullInt64(i.ObjectTypeID),
		MunicipalityID: sql_common.NewNullInt64(i.MunicipalityID),
	}
}

func newModelFromCreateData(data *repository.CreateObjectTemplateData) *modelObjectTemplate {
	return &modelObjectTemplate{
		Name:           sql_common.NewNullString(data.Name),
		ObjectTypeID:   sql_common.NewNullInt64(data.ObjectType),
		MunicipalityID: sql_common.NewNullInt64(data.MunicipalityID),
	}
}
