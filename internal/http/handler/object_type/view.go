package object_type

import "municipality_app/internal/domain/entity"

type objectTypeView struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func newObjectTypeView(i *entity.ObjectType) *objectTypeView {
	if i == nil {
		return nil
	}

	return &objectTypeView{
		ID:   i.ID,
		Name: i.Name,
	}
}

func newObjectTypeViews(objectTypes []entity.ObjectType) []objectTypeView {
	result := make([]objectTypeView, 0, len(objectTypes))

	for _, objectType := range objectTypes {
		result = append(result, *newObjectTypeView(&objectType))
	}

	return result
}
