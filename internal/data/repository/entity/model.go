package entity

import (
	"database/sql"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
)

type modelEntity struct {
	ID               sql.NullInt64
	Name             sql.NullString
	EntityTemplateID sql.NullInt64
	Description      sql.NullString
}

type Entity struct {
	ID               int64
	Name             string
	EntityTemplateID int64
	Description      string
}

type EntityTemplate struct {
	ID           int64
	Name         string
	EntityTypeID int64
	PartitionID  int64
}

func (m *modelEntity) convert() *entity.Entity {
	return &entity.Entity{
		ID:               m.ID.Int64,
		Name:             m.Name.String,
		EntityTemplateID: m.EntityTemplateID.Int64,
		Description:      m.Description.String,
	}
}

func newModelEntity(i *entity.Entity) *modelEntity {
	return &modelEntity{
		ID:               sql_common.NewNullInt64(i.ID),
		Name:             sql_common.NewNullString(i.Name),
		EntityTemplateID: sql_common.NewNullInt64(i.EntityTemplateID),
		Description:      sql_common.NewNullString(i.Description),
	}
}

func newModelFromCreateData(data *repository.CreateEntityData) *modelEntity {
	return &modelEntity{
		Name:             sql_common.NewNullString(data.Name),
		EntityTemplateID: sql_common.NewNullInt64(data.EntityTemplateID),
		Description:      sql_common.NewNullString(data.Description),
	}
}
