package location

import (
	"database/sql"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
)

type modelLocation struct {
	ID        sql.NullInt64
	Address   sql.NullString
	Latitude  sql.NullFloat64
	Longitude sql.NullFloat64
	Geometry  sql.NullString
}

type Location struct {
	ID        int64
	Address   string
	Latitude  float64
	Longitude float64
	Geometry  *string
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

func (m *modelLocation) convert() *entity.Location {
	return &entity.Location{
		ID:        m.ID.Int64,
		Address:   m.Address.String,
		Latitude:  m.Latitude.Float64,
		Longitude: m.Longitude.Float64,
		Geometry:  nil,
	}
}

func newModelObject(i *entity.Location) *modelLocation {
	return &modelLocation{
		ID:        sql_common.NewNullInt64(i.ID),
		Address:   sql_common.NewNullString(i.Address),
		Latitude:  sql_common.NewNullFloat64(i.Latitude),
		Longitude: sql_common.NewNullFloat64(i.Longitude),
		Geometry:  sql_common.NewNullStringPtr(i.Geometry),
	}
}

func newModelFromCreateData(data *repository.CreateLocationData) *modelLocation {
	return &modelLocation{
		Address:   sql_common.NewNullStringPtr(data.Address),
		Latitude:  sql_common.NewNullFloat64Ptr(data.Latitude),
		Longitude: sql_common.NewNullFloat64Ptr(data.Longitude),
		Geometry:  sql_common.NewNullStringPtr(data.Geometry),
	}
}
