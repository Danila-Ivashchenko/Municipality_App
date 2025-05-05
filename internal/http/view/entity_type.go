package view

import "municipality_app/internal/domain/entity"

type EntityTypeView struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func NewEntityTypeView(i *entity.EntityType) *EntityTypeView {
	if i == nil {
		return nil
	}

	return &EntityTypeView{
		ID:   i.ID,
		Name: i.Name,
	}
}
