package entity_attribute_value

import (
	"database/sql"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
)

type model struct {
	ID                sql.NullInt64
	EntityAttributeID sql.NullInt64
	EntityID          sql.NullInt64
	Value             sql.NullString
}

func (m *model) convert() *entity.EntityAttributeValue {
	return &entity.EntityAttributeValue{
		ID:                m.ID.Int64,
		EntityAttributeID: m.EntityAttributeID.Int64,
		EntityID:          m.EntityID.Int64,
		Value:             m.Value.String,
	}
}

func newModel(i *entity.EntityAttributeValue) *model {
	return &model{
		ID:                sql_common.NewNullInt64(i.ID),
		EntityAttributeID: sql_common.NewNullInt64(i.EntityAttributeID),
		EntityID:          sql_common.NewNullInt64(i.EntityID),
		Value:             sql_common.NewNullString(i.Value),
	}
}
