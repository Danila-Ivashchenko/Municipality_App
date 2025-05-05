package route_object

import (
	"database/sql"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
)

type model struct {
	ID             sql.NullInt64
	Name           sql.NullString
	RouteID        sql.NullInt64
	OrderNumber    sql.NullInt64
	SourceObjectID sql.NullInt64
	LocationID     sql.NullInt64
}

func (m *model) convert() *entity.RouteObject {
	var (
		sourceObjectID *int64
		locationID     *int64
	)

	if m.SourceObjectID.Valid {
		value := m.SourceObjectID.Int64
		sourceObjectID = &value
	}

	if m.LocationID.Valid {
		value := m.LocationID.Int64
		locationID = &value
	}

	return &entity.RouteObject{
		ID:             m.ID.Int64,
		Name:           m.Name.String,
		RouteID:        m.RouteID.Int64,
		OrderNumber:    int(m.OrderNumber.Int64),
		SourceObjectID: sourceObjectID,
		LocationID:     locationID,
	}
}

func newModel(data *entity.RouteObject) *model {
	if data == nil {
		return nil
	}

	return &model{
		ID:             sql_common.NewNullInt64(data.ID),
		Name:           sql_common.NewNullString(data.Name),
		RouteID:        sql_common.NewNullInt64(data.RouteID),
		OrderNumber:    sql_common.NewNullInt64(int64(data.OrderNumber)),
		SourceObjectID: sql_common.NewNullInt64Ptr(data.SourceObjectID),
		LocationID:     sql_common.NewNullInt64Ptr(data.LocationID),
	}
}
