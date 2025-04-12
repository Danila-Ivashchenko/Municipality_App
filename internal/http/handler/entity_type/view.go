package entity_type

import "municipality_app/internal/domain/entity"

type entityTypeView struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func newEntityTypeView(i *entity.EntityType) *entityTypeView {
	if i == nil {
		return nil
	}

	return &entityTypeView{
		ID:   i.ID,
		Name: i.Name,
	}
}

func newEntityTypeViews(entityTypes []entity.EntityType) []entityTypeView {
	result := make([]entityTypeView, 0, len(entityTypes))

	for _, entityType := range entityTypes {
		result = append(result, *newEntityTypeView(&entityType))
	}

	return result
}
