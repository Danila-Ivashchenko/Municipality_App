package view

import "municipality_app/internal/domain/entity"

type EntityView struct {
	ID               int64                   `json:"id"`
	Name             string                  `json:"name"`
	EntityTemplateID int64                   `json:"entity_template_id"`
	Description      string                  `json:"description"`
	Attributes       []EntityAttributeExView `json:"attributes"`
}

func NewEntityView(i *entity.EntityEx) EntityView {
	result := EntityView{
		ID:               i.ID,
		Name:             i.Name,
		EntityTemplateID: i.EntityTemplateID,
		Description:      i.Description,
	}

	for _, attr := range i.AttributeValues {
		result.Attributes = append(result.Attributes, NewEntityAttributeExView(attr))
	}

	return result
}

func NewEntityViews(data []entity.EntityEx) []EntityView {
	result := make([]EntityView, 0, len(data))

	for _, i := range data {
		result = append(result, NewEntityView(&i))
	}

	return result
}
