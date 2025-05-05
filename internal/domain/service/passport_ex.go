package service

import (
	"context"
	"municipality_app/internal/domain/entity"
)

type PassportExService interface {
	GetByIDAndMunicipalityID(ctx context.Context, id, municipalityID int64) (*entity.PassportEx, error)
	UpdateChapterEx(ctx context.Context, data *PassportUpdateChaptersData) (*entity.ChapterEx, error)
	CreateChapterEx(ctx context.Context, data *PassportCreateChaptersData) (*entity.ChapterEx, error)
	DeleteChapterEx(ctx context.Context, passportID, chapterID int64) error

	GetChapterEx(ctx context.Context, chapterID, passportID int64) (*entity.ChapterEx, error)
	GetPartitionEx(ctx context.Context, partitionID, chapterID int64) (*entity.PartitionEx, error)
}

type PassportCreateChaptersData struct {
	PassportID     int64
	MunicipalityID int64
	ChapterData    CreateChapterExData
}

type PassportUpdateChaptersData struct {
	PassportID     int64
	MunicipalityID int64
	ChapterData    UpdateChapterExData
}

type CreateChapterExData struct {
	ChapterData CreateOneChapterData
}

type UpdateChapterExData struct {
	ChapterID      int64
	ChapterData    *UpdateChapterData
	PartitionsData *PartitionsData
}

type PartitionsData struct {
	Create []CreateOnePartitionData
	Update []UpdatePartitionData
	Delete []int64
}

type PassportCreatePartitionData struct {
	PassportID    int64
	ChapterID     int64
	PartitionData CreateOnePartitionData
}
