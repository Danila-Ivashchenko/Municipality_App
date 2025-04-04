package object_attribute_value

import (
	"database/sql"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
)

type model struct {
	ID                sql.NullInt64
	ObjectAttributeID sql.NullInt64
	ObjectID          sql.NullInt64
	Value             sql.NullString
}

func (m *model) convert() *entity.ObjectAttributeValue {
	return &entity.ObjectAttributeValue{
		ID:                m.ID.Int64,
		ObjectAttributeID: m.ObjectAttributeID.Int64,
		ObjectID:          m.ObjectID.Int64,
		Value:             m.Value.String,
	}
}

func newModel(i *entity.ObjectAttributeValue) *model {
	return &model{
		ID:                sql_common.NewNullInt64(i.ID),
		ObjectAttributeID: sql_common.NewNullInt64(i.ObjectAttributeID),
		ObjectID:          sql_common.NewNullInt64(i.ObjectID),
		Value:             sql_common.NewNullString(i.Value),
	}
}
