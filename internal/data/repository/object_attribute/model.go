package object_attribute

import (
	"database/sql"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
)

type model struct {
	ID               sql.NullInt64
	ObjectTemplateID sql.NullInt64
	Name             sql.NullString
	DefaultValue     sql.NullString
	ToShow           sql.NullBool
}

func (m *model) convert() *entity.ObjectAttribute {
	return &entity.ObjectAttribute{
		ID:               m.ID.Int64,
		ObjectTemplateID: m.ObjectTemplateID.Int64,
		Name:             m.Name.String,
		DefaultValue:     m.DefaultValue.String,
		ToShow:           m.ToShow.Bool,
	}
}

func newModel(i *entity.ObjectAttribute) *model {
	return &model{
		ID:               sql_common.NewNullInt64(i.ID),
		ObjectTemplateID: sql_common.NewNullInt64(i.ObjectTemplateID),
		Name:             sql_common.NewNullString(i.Name),
		DefaultValue:     sql_common.NewNullString(i.DefaultValue),
		ToShow:           sql_common.NewNullBool(i.ToShow),
	}
}
