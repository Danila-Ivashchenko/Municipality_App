package entity_attribute

import (
	"database/sql"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
)

type model struct {
	ID               sql.NullInt64
	EntityTemplateID sql.NullInt64
	Name             sql.NullString
	DefaultValue     sql.NullString
	ToShow           sql.NullBool
}

func (m *model) convert() *entity.EntityAttribute {
	return &entity.EntityAttribute{
		ID:               m.ID.Int64,
		EntityTemplateID: m.EntityTemplateID.Int64,
		Name:             m.Name.String,
		DefaultValue:     m.DefaultValue.String,
		ToShow:           m.ToShow.Bool,
	}
}

func newModel(i *entity.EntityAttribute) *model {
	return &model{
		ID:               sql_common.NewNullInt64(i.ID),
		EntityTemplateID: sql_common.NewNullInt64(i.EntityTemplateID),
		Name:             sql_common.NewNullString(i.Name),
		DefaultValue:     sql_common.NewNullString(i.DefaultValue),
		ToShow:           sql_common.NewNullBool(i.ToShow),
	}
}
