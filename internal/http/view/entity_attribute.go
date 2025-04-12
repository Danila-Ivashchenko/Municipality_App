package view

import "municipality_app/internal/domain/entity"

type EntityAttributeView struct {
	ID               int64  `json:"id"`
	Name             string `json:"name"`
	EntityTemplateID int64  `json:"entity_template_id"`
	DefaultValue     string `json:"default_value"`
}

func NewEntityAttributeView(i entity.EntityAttribute) EntityAttributeView {
	return EntityAttributeView{
		ID:               i.ID,
		Name:             i.Name,
		EntityTemplateID: i.EntityTemplateID,
		DefaultValue:     i.DefaultValue,
	}
}

type EntityAttributeExView struct {
	ID                int64  `json:"id"`
	Name              string `json:"name"`
	ObjectAttributeID int64  `json:"entity_attribute_id"`
	ObjectID          int64  `json:"entity_id"`
	Value             string `json:"value"`
}

func NewEntityAttributeExView(i entity.EntityAttributeValueEx) EntityAttributeExView {
	return EntityAttributeExView{
		ID:                i.Value.ID,
		Name:              i.Attribute.Name,
		ObjectAttributeID: i.Attribute.ID,
		ObjectID:          i.Value.EntityID,
		Value:             i.Value.Value,
	}
}
