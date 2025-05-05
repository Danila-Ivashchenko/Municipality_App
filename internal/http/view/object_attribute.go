package view

import "municipality_app/internal/domain/entity"

type ObjectAttributeView struct {
	ID               int64  `json:"id"`
	Name             string `json:"name"`
	ObjectTemplateID int64  `json:"object_template_id"`
	DefaultValue     string `json:"default_value"`
	ToShow           bool   `json:"to_show"`
}

func NewObjectAttributeView(i entity.ObjectAttribute) ObjectAttributeView {
	return ObjectAttributeView{
		ID:               i.ID,
		Name:             i.Name,
		ObjectTemplateID: i.ObjectTemplateID,
		DefaultValue:     i.DefaultValue,
		ToShow:           i.ToShow,
	}
}

type ObjectAttributeExView struct {
	ID                int64  `json:"id"`
	Name              string `json:"name"`
	ObjectAttributeID int64  `json:"object_attribute_id"`
	ObjectID          int64  `json:"object_id"`
	Value             string `json:"value"`
}

func NewObjectAttributeExView(i entity.ObjectAttributeValueEx) ObjectAttributeExView {
	return ObjectAttributeExView{
		ID:                i.Value.ID,
		Name:              i.Attribute.Name,
		ObjectAttributeID: i.Attribute.ID,
		ObjectID:          i.Value.ObjectID,
		Value:             i.Value.Value,
	}
}
