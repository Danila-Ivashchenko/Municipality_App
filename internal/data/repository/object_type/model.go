package object_type

import (
	"database/sql"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
)

type modelObjectType struct {
	ID   sql.NullInt64
	Name sql.NullString
}

func (m *modelObjectType) convert() *entity.ObjectType {
	return &entity.ObjectType{
		ID:   m.ID.Int64,
		Name: m.Name.String,
	}
}

func newModelObjectType(i *entity.ObjectType) *modelObjectType {
	return &modelObjectType{
		ID:   sql_common.NewNullInt64(i.ID),
		Name: sql_common.NewNullString(i.Name),
	}
}

func newModelFromCreateData(data *repository.CreateObjectType) *modelObjectType {
	return &modelObjectType{
		Name: sql_common.NewNullString(data.Name),
	}
}
