package view

import (
	"municipality_app/internal/domain/entity"
	"sort"
)

type ChapterExView struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	PassportID  int64  `json:"passport_id"`
	Description string `json:"description"`
	Text        string `json:"text"`
	OrderNumber uint   `json:"order_number"`

	Partitions []PartitionExView `json:"partitions"`
}

func NewChapterExView(i entity.ChapterEx) ChapterExView {
	chapter := ChapterExView{
		ID:          i.ID,
		Name:        i.Name,
		PassportID:  i.PassportID,
		Description: i.Description,
		Text:        i.Text,
		OrderNumber: i.OrderNumber,
	}

	for _, p := range i.Partitions {
		chapter.Partitions = append(chapter.Partitions, NewPartitionExView(p))
	}

	sort.Slice(chapter.Partitions, func(i, j int) bool {
		return chapter.Partitions[i].OrderNumber < chapter.Partitions[j].OrderNumber
	})

	return chapter
}
