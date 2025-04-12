package service

import (
	"context"
	"municipality_app/internal/domain/entity"
)

type PartitionService interface {
	Create(ctx context.Context, data *CreateOnePartitionData) (*entity.PartitionEx, error)

	Update(ctx context.Context, data *UpdatePartitionData) (*entity.PartitionEx, error)
	DeleteToChapter(ctx context.Context, ids []int64, chapterID int64) error

	GetByChapterID(ctx context.Context, chapterID int64) ([]entity.Partition, error)
	GetByIDAndChapterID(ctx context.Context, id, chapterID int64) (*entity.Partition, error)
	GetByIDsAndChapterID(ctx context.Context, ids []int64, chapterID int64) ([]entity.Partition, error)

	GetByIDs(ctx context.Context, ids []int64) ([]entity.Partition, error)

	GetExByID(ctx context.Context, id int64) (*entity.PartitionEx, error)
}

type CreatePartitionsData struct {
	ChapterID    int64
	ChaptersData []CreateChapterData
}

type UpdatePartitionsData struct {
	ChapterID    int64
	ChaptersData []UpdateChapterData
}

type CreateOnePartitionData struct {
	Name        string
	ChapterID   int64
	Description string
	Text        string
	OrderNumber uint
	ObjectIDs   []int64
	EntityIDs   []int64
}

type UpdatePartitionData struct {
	ID          int64
	Name        *string
	Description *string
	Text        *string
	OrderNumber *uint
	ObjectIDs   *[]int64
	EntityIDs   *[]int64
}
