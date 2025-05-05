package view

import "municipality_app/internal/domain/entity"

type ObjectTypeView struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func NewObjectTypeView(i *entity.ObjectType) *ObjectTypeView {
	if i == nil {
		return nil
	}

	return &ObjectTypeView{
		ID:   i.ID,
		Name: i.Name,
	}
}
