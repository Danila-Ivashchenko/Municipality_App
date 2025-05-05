package route

import (
	"database/sql"
	"encoding/json"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
)

type model struct {
	ID                sql.NullInt64
	PartitionID       sql.NullInt64
	Name              sql.NullString
	Length            sql.NullInt64
	Duration          sql.NullInt64
	Level             sql.NullInt64
	MovementWay       sql.NullString
	Seasonality       sql.NullString
	PersonalEquipment sql.NullString
	Dangerous         sql.NullString
	Rules             sql.NullString
	RouteEquipment    sql.NullString
	Geometry          *[]byte
}

func (m *model) convert() *entity.Route {
	var (
		geometryPtr *entity.Geometry
		geometryVal entity.Geometry
	)

	if m.Geometry != nil {

		err := json.Unmarshal(*m.Geometry, &geometryVal)
		if err == nil {
			geometryPtr = &geometryVal
		}
	}

	return &entity.Route{
		ID:                m.ID.Int64,
		PartitionID:       m.PartitionID.Int64,
		Name:              m.Name.String,
		Length:            m.Length.Int64,
		Duration:          m.Duration.Int64,
		Level:             uint(m.Level.Int64),
		MovementWay:       m.MovementWay.String,
		Seasonality:       m.Seasonality.String,
		PersonalEquipment: m.PersonalEquipment.String,
		Dangerous:         m.Dangerous.String,
		Rules:             m.Rules.String,
		RouteEquipment:    m.RouteEquipment.String,
		Geometry:          geometryPtr,
	}
}

func newModel(data *entity.Route) *model {
	var (
		geometry *[]byte
	)

	if data == nil {
		return nil
	}

	if data.Geometry != nil {
		geometryVal, err := json.Marshal(data.Geometry)
		if err == nil {
			geometry = &geometryVal
		}
	}

	return &model{
		ID:                sql_common.NewNullInt64(data.ID),
		PartitionID:       sql_common.NewNullInt64(data.PartitionID),
		Name:              sql_common.NewNullString(data.Name),
		Length:            sql_common.NewNullInt64(data.Length),
		Duration:          sql_common.NewNullInt64(data.Duration),
		Level:             sql_common.NewNullInt64(int64(data.Level)),
		MovementWay:       sql_common.NewNullString(data.MovementWay),
		Seasonality:       sql_common.NewNullString(data.Seasonality),
		PersonalEquipment: sql_common.NewNullString(data.PersonalEquipment),
		Dangerous:         sql_common.NewNullString(data.Dangerous),
		Rules:             sql_common.NewNullString(data.Rules),
		RouteEquipment:    sql_common.NewNullString(data.RouteEquipment),
		Geometry:          geometry,
	}
}
