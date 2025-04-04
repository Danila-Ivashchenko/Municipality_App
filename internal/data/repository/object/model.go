package object

import (
	"database/sql"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
)

type modelObject struct {
	ID               sql.NullInt64
	Name             sql.NullString
	ObjectTemplateID sql.NullInt64
	LocationID       sql.NullInt64
	Description      sql.NullString
}

type Object struct {
	ID               int64
	Name             string
	ObjectTemplateID int64
	LocationID       int64
	Description      string
}

type ObjectTemplate struct {
	ID           int64
	Name         string
	ObjectTypeID int64
	PartitionID  int64
}

func (m *modelObject) convert() *entity.Object {
	var (
		locationIDPtr *int64
	)

	if m.LocationID.Valid {
		locationIDPtr = &m.LocationID.Int64
	}

	return &entity.Object{
		ID:               m.ID.Int64,
		Name:             m.Name.String,
		ObjectTemplateID: m.ObjectTemplateID.Int64,
		LocationID:       locationIDPtr,
		Description:      m.Description.String,
	}
}

func newModelObject(i *entity.Object) *modelObject {
	return &modelObject{
		ID:               sql_common.NewNullInt64(i.ID),
		Name:             sql_common.NewNullString(i.Name),
		ObjectTemplateID: sql_common.NewNullInt64(i.ObjectTemplateID),
		LocationID:       sql_common.NewNullInt64Ptr(i.LocationID),
		Description:      sql_common.NewNullString(i.Description),
	}
}

func newModelFromCreateData(data *repository.CreateObjectData) *modelObject {
	return &modelObject{
		Name:             sql_common.NewNullString(data.Name),
		ObjectTemplateID: sql_common.NewNullInt64(data.ObjectTemplateID),
		LocationID:       sql_common.NewNullInt64Ptr(data.LocationID),
		Description:      sql_common.NewNullString(data.Description),
	}
}
