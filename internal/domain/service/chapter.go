package service

import (
	"context"
	"municipality_app/internal/domain/entity"
)

type ChapterService interface {
	Create(ctx context.Context, data *CreateOneChapterData) (*entity.Chapter, error)

	Update(ctx context.Context, data *UpdateChapterData) (*entity.Chapter, error)
	DeleteToPassport(ctx context.Context, ids []int64, passportID int64) error

	GetByPassportID(ctx context.Context, passportID int64) ([]entity.Chapter, error)
	GetByIDAndPassportID(ctx context.Context, id, passportID int64) (*entity.Chapter, error)
	GetByIDsAndPassportID(ctx context.Context, ids []int64, passportID int64) ([]entity.Chapter, error)
}

type CreateChaptersData struct {
	PassportID   int64
	ChaptersData []CreateChapterData
}

type UpdateChaptersData struct {
	PassportID   int64
	ChaptersData []UpdateChapterData
}

type CreateChapterData struct {
	Name        string
	Description string
	Text        string
	OrderNumber uint
}

type CreateOneChapterData struct {
	Name        string
	PassportID  int64
	Description string
	Text        string
	OrderNumber uint
}

type UpdateChapterData struct {
	ID          int64
	Name        *string
	Description *string
	Text        *string
	OrderNumber *uint
}
