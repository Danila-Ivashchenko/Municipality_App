package entity_type

import (
	"database/sql"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
)

type modelEntityType struct {
	ID   sql.NullInt64
	Name sql.NullString
}

func (m *modelEntityType) convert() *entity.EntityType {
	return &entity.EntityType{
		ID:   m.ID.Int64,
		Name: m.Name.String,
	}
}

func newModelEntityType(i *entity.EntityType) *modelEntityType {
	return &modelEntityType{
		ID:   sql_common.NewNullInt64(i.ID),
		Name: sql_common.NewNullString(i.Name),
	}
}

func newModelFromCreateData(data *repository.CreateEntityType) *modelEntityType {
	return &modelEntityType{
		Name: sql_common.NewNullString(data.Name),
	}
}
