package view

import "municipality_app/internal/domain/entity"

type ObjectTemplateExView struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	MunicipalityID int64  `json:"municipality_id"`
	ObjectType     int64  `json:"object_type"`

	Objects    []ObjectView          `json:"objects"`
	Attributes []ObjectAttributeView `json:"attributes"`
}

func NewObjectTemplateExView(i *entity.ObjectTemplateEx) *ObjectTemplateExView {
	result := &ObjectTemplateExView{
		ID:             i.Template.ID,
		Name:           i.Template.Name,
		MunicipalityID: i.Template.MunicipalityID,
		ObjectType:     i.Template.ObjectTypeID,
	}

	for _, obj := range i.Objects {
		result.Objects = append(result.Objects, NewObjectView(&obj))
	}

	for _, attr := range i.Attributes {
		result.Attributes = append(result.Attributes, NewObjectAttributeView(attr))
	}

	return result
}

type ObjectTemplateView struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	ObjectType     int64  `json:"object_type"`
	MunicipalityID int64  `json:"municipality_id"`
}

func NewObjectTemplateView(i *entity.ObjectTemplate) *ObjectTemplateView {
	if i == nil {
		return nil
	}

	return &ObjectTemplateView{
		ID:             i.ID,
		Name:           i.Name,
		ObjectType:     i.ObjectTypeID,
		MunicipalityID: i.MunicipalityID,
	}
}
