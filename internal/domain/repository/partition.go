package repository

import (
	"context"
	"municipality_app/internal/domain/entity"
)

type PartitionRepository interface {
	Create(ctx context.Context, data *CreatePartitionData) (*entity.Partition, error)

	Update(ctx context.Context, data *entity.Partition) error
	Delete(ctx context.Context, id int64) error

	ChangeOrder(ctx context.Context, id int64, newOrder uint) error

	GetByChapterID(ctx context.Context, chapterID int64) ([]entity.Partition, error)
	GetByIDAndChapterID(ctx context.Context, id, chapterID int64) (*entity.Partition, error)
	GetByIDsAndChapterID(ctx context.Context, ids []int64, chapterID int64) ([]entity.Partition, error)

	GetByNameAndChapterID(ctx context.Context, name string, chapterID int64) (*entity.Partition, error)
	GetByNamesAndChapterID(ctx context.Context, names []string, chapterID int64) ([]entity.Partition, error)

	GetByID(ctx context.Context, id int64) (*entity.Partition, error)
	GetByIDs(ctx context.Context, ids []int64) ([]entity.Partition, error)
}

type CreatePartitionData struct {
	ChapterID   int64
	Name        string
	Description string
	Text        string
	OrderNumber uint
}

type UpdatePartitionData struct {
	ID          int64
	Name        string
	Description string
	Text        string
	OrderNumber uint
}
