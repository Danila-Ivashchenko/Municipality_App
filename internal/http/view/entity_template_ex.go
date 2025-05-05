package view

import "municipality_app/internal/domain/entity"

type EntityTemplateExView struct {
	ID             int64           `json:"id"`
	Name           string          `json:"name"`
	MunicipalityID int64           `json:"municipality_id"`
	EntityTypeID   int64           `json:"entity_type_id"`
	EntityType     *EntityTypeView `json:"entity_type"`

	Entities   []EntityView          `json:"entities"`
	Attributes []EntityAttributeView `json:"attributes"`
}

func NewEntityTemplateExView(i *entity.EntityTemplateEx) *EntityTemplateExView {
	result := &EntityTemplateExView{
		ID:             i.Template.ID,
		Name:           i.Template.Name,
		MunicipalityID: i.Template.MunicipalityID,
		EntityTypeID:   i.Template.EntityTypeID,
		EntityType:     NewEntityTypeView(i.EntityType),
	}

	for _, obj := range i.Entities {
		result.Entities = append(result.Entities, NewEntityView(&obj))
	}

	for _, attr := range i.Attributes {
		result.Attributes = append(result.Attributes, NewEntityAttributeView(attr))
	}

	return result
}

type EntityTemplateView struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	EntityType     int64  `json:"entity_type"`
	MunicipalityID int64  `json:"municipality_id"`
}

func NewEntityTemplateView(i *entity.EntityTemplate) *EntityTemplateView {
	if i == nil {
		return nil
	}

	return &EntityTemplateView{
		ID:             i.ID,
		Name:           i.Name,
		EntityType:     i.EntityTypeID,
		MunicipalityID: i.MunicipalityID,
	}
}
