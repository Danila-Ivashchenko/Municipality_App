package view

import "municipality_app/internal/domain/entity"

type ObjectView struct {
	ID               int64                   `json:"id"`
	Name             string                  `json:"name"`
	ObjectTemplateID int64                   `json:"object_template_id"`
	Location         *ObjectLocationView     `json:"location"`
	Description      string                  `json:"description"`
	Attributes       []ObjectAttributeExView `json:"attributes"`
}

type ObjectLocationView struct {
	ID        int64   `json:"id"`
	Address   string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Geometry  *string `json:"geometry,omitempty"`
}

func NewObjectLocationView(i *entity.Location) *ObjectLocationView {
	if i == nil {
		return nil
	}

	return &ObjectLocationView{
		ID:        i.ID,
		Address:   i.Address,
		Latitude:  i.Latitude,
		Longitude: i.Longitude,
		Geometry:  i.Geometry,
	}
}

func NewObjectView(i *entity.ObjectEx) ObjectView {
	result := ObjectView{
		ID:               i.ID,
		Name:             i.Name,
		ObjectTemplateID: i.ObjectTemplateID,
		Description:      i.Description,
		Location:         NewObjectLocationView(i.LocationID),
	}

	for _, attr := range i.AttributeValues {
		result.Attributes = append(result.Attributes, NewObjectAttributeExView(attr))
	}

	return result
}

func NewObjectViews(data []entity.ObjectEx) []ObjectView {
	result := make([]ObjectView, 0, len(data))

	for _, i := range data {
		result = append(result, NewObjectView(&i))
	}

	return result
}
