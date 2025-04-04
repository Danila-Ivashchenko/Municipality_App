package entity

import (
	"time"
)

type PassportEx struct {
	ID             int64
	Name           string
	MunicipalityID int64

	Description  string
	Year         string
	RevisionCode string

	CreatedAt time.Time
	UpdatedAt time.Time
	IsMain    bool
	IsHidden  bool

	Chapters []ChapterEx
}

func NewPassportEx(i *Passport, chapters []ChapterEx) *PassportEx {
	result := &PassportEx{
		ID:             i.ID,
		Name:           i.Name,
		MunicipalityID: i.MunicipalityID,
		Description:    i.Description,
		Year:           i.Year,
		RevisionCode:   i.RevisionCode,
		CreatedAt:      i.CreatedAt,
		UpdatedAt:      i.UpdatedAt,
		IsMain:         i.IsMain,
		IsHidden:       i.IsHidden,
	}

	result.Chapters = chapters

	return result
}

type ChapterEx struct {
	ID          int64
	PassportID  int64
	Name        string
	Description string
	Text        string
	OrderNumber uint
	Partitions  []PartitionEx
}

//func NewChapterEx(i Chapter, partitions []Partition) ChapterEx {
//	result := ChapterEx{
//		ID:          i.ID,
//		PassportID:  i.PassportID,
//		Name:        i.Name,
//		Description: i.Description,
//		Text:        i.Text,
//		OrderNumber: i.OrderNumber,
//	}
//
//	for _, partition := range partitions {
//		result.Partitions = append(result.Partitions, NewPartitionEx(partition))
//	}
//
//	return result
//}

func NewChapterExPtr(i *Chapter, partitions []PartitionEx) *ChapterEx {
	result := &ChapterEx{
		ID:          i.ID,
		PassportID:  i.PassportID,
		Name:        i.Name,
		Description: i.Description,
		Text:        i.Text,
		OrderNumber: i.OrderNumber,
	}

	for _, partition := range partitions {
		result.Partitions = append(result.Partitions, partition)
	}

	return result
}
