package entity

type Partition struct {
	ID          int64
	ChapterID   int64
	Name        string
	Description string
	Text        string
	OrderNumber uint
}

type PartitionEx struct {
	ID          int64
	ChapterID   int64
	Name        string
	Description string
	Text        string
	OrderNumber uint
	Objects     []ObjectTemplateEx
}

func NewPartitionEx(i Partition, objects []ObjectTemplateEx) PartitionEx {
	result := PartitionEx{
		ID:          i.ID,
		ChapterID:   i.ChapterID,
		Name:        i.Name,
		Description: i.Description,
		Text:        i.Text,
		OrderNumber: i.OrderNumber,
		Objects:     objects,
	}

	return result
}
