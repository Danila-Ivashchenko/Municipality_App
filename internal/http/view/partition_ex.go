package view

import "municipality_app/internal/domain/entity"

type PartitionExView struct {
	ID          int64                  `json:"id"`
	Name        string                 `json:"name"`
	ChapterID   int64                  `json:"chapter_id"`
	Description string                 `json:"description"`
	Text        string                 `json:"text"`
	OrderNumber uint                   `json:"order_number"`
	Objects     []ObjectTemplateExView `json:"objects,omitempty"`
	Entities    []EntityTemplateExView `json:"entities,omitempty"`
	Routes      []RouteView            `json:"routes,omitempty"`
}

func NewPartitionExView(i entity.PartitionEx) PartitionExView {
	result := PartitionExView{
		ID:          i.ID,
		Name:        i.Name,
		ChapterID:   i.ChapterID,
		Description: i.Description,
		Text:        i.Text,
		OrderNumber: i.OrderNumber,
	}

	for _, o := range i.Objects {
		result.Objects = append(result.Objects, *NewObjectTemplateExView(&o))
	}

	for _, e := range i.Entities {
		result.Entities = append(result.Entities, *NewEntityTemplateExView(&e))
	}

	for _, route := range i.Routes {
		result.Routes = append(result.Routes, *NewRouteView(&route))
	}

	return result
}
