package repository

import (
	"context"
	"municipality_app/internal/domain/entity"
)

type ChapterRepository interface {
	Create(ctx context.Context, data *CreateChapterData) (*entity.Chapter, error)
	CreateMultiply(ctx context.Context, data []CreateChapterData) ([]entity.Chapter, error)

	Update(ctx context.Context, data *entity.Chapter) error
	Delete(ctx context.Context, id int64) error

	ChangeOrder(ctx context.Context, id int64, newOrder uint) error

	GetByPassportID(ctx context.Context, passportID int64) ([]entity.Chapter, error)
	GetByIDAndPassportID(ctx context.Context, id, passportID int64) (*entity.Chapter, error)
	GetByIDsAndPassportID(ctx context.Context, ids []int64, passportID int64) ([]entity.Chapter, error)

	GetByNameAndPassportID(ctx context.Context, name string, passportID int64) (*entity.Chapter, error)
	GetByNamesAndPassportID(ctx context.Context, names []string, passportID int64) ([]entity.Chapter, error)

	GetByID(ctx context.Context, id int64) (*entity.Chapter, error)
}

type CreateChapterData struct {
	PassportID  int64
	Name        string
	Description string
	Text        string
	OrderNumber uint
}

type UpdateChapterData struct {
	ID          int64
	Name        string
	Description string
	Text        string
	OrderNumber uint
}
