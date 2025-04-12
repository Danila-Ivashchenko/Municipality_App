package entity_template

import (
	"database/sql"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
)

type modelEntityTemplate struct {
	ID             sql.NullInt64
	Name           sql.NullString
	EntityTypeID   sql.NullInt64
	MunicipalityID sql.NullInt64
}

type EntityTemplate struct {
	ID           int64
	Name         string
	EntityTypeID int64
	PartitionID  int64
}

func (m *modelEntityTemplate) convert() *entity.EntityTemplate {
	return &entity.EntityTemplate{
		ID:             m.ID.Int64,
		Name:           m.Name.String,
		EntityTypeID:   m.EntityTypeID.Int64,
		MunicipalityID: m.MunicipalityID.Int64,
	}
}

func newModelEntityTemplate(i *entity.EntityTemplate) *modelEntityTemplate {
	return &modelEntityTemplate{
		ID:             sql_common.NewNullInt64(i.ID),
		Name:           sql_common.NewNullString(i.Name),
		EntityTypeID:   sql_common.NewNullInt64(i.EntityTypeID),
		MunicipalityID: sql_common.NewNullInt64(i.MunicipalityID),
	}
}

func newModelFromCreateData(data *repository.CreateEntityTemplateData) *modelEntityTemplate {
	return &modelEntityTemplate{
		Name:           sql_common.NewNullString(data.Name),
		EntityTypeID:   sql_common.NewNullInt64(data.EntityType),
		MunicipalityID: sql_common.NewNullInt64(data.MunicipalityID),
	}
}
