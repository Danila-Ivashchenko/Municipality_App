package passport_ex

import (
	"context"
	"errors"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/service"
)

func (svc *passportExService) GetByIDAndMunicipalityID(ctx context.Context, id, municipalityID int64) (*entity.PassportEx, error) {
	passport, err := svc.PassportService.GetByIDAndMunicipalityID(ctx, id, municipalityID)
	if err != nil {
		return nil, err
	}

	if passport == nil {
		return nil, errors.New("passport not found")
	}

	chaptersEx, err := svc.GetChaptersEx(ctx, passport.ID)
	if err != nil {
		return nil, err
	}

	return entity.NewPassportEx(passport, chaptersEx), nil
}

func (svc *passportExService) GetChaptersEx(ctx context.Context, passportID int64) ([]entity.ChapterEx, error) {
	var (
		result    []entity.ChapterEx
		chapterEx *entity.ChapterEx
	)

	chapters, err := svc.ChapterService.GetByPassportID(ctx, passportID)
	if err != nil {
		return nil, err
	}

	for _, ch := range chapters {
		chapterEx, err = svc.GetChapterEx(ctx, ch.ID, ch.PassportID)
		if err != nil {
			return nil, err
		}

		if chapterEx != nil {
			result = append(result, *chapterEx)
		}
	}

	return result, nil
}

func (svc *passportExService) GetChapterEx(ctx context.Context, chapterID, passportID int64) (*entity.ChapterEx, error) {
	chapters, err := svc.ChapterService.GetByIDAndPassportID(ctx, chapterID, passportID)
	if err != nil {
		return nil, err
	}

	if chapters == nil {
		return nil, errors.New("chapter not found")
	}

	partitionsEx, err := svc.GetPartitionsExByChapterID(ctx, chapterID)
	if err != nil {
		return nil, err
	}

	return entity.NewChapterExPtr(chapters, partitionsEx), nil
}

func (svc *passportExService) GetPartitionEx(ctx context.Context, partitionID, chapterID int64) (*entity.PartitionEx, error) {
	partition, err := svc.PartitionService.GetByIDAndChapterID(ctx, partitionID, chapterID)
	if err != nil {
		return nil, err
	}

	if partition == nil {
		return nil, errors.New("partition not found")
	}

	return svc.PartitionService.GetExByID(ctx, partition.ID)
}

func (svc *passportExService) GetPartitionsExByChapterID(ctx context.Context, chapterID int64) ([]entity.PartitionEx, error) {
	var (
		partitionIDs []int64
	)
	partitions, err := svc.PartitionService.GetByChapterID(ctx, chapterID)
	if err != nil {
		return nil, err
	}

	for _, partition := range partitions {
		partitionIDs = append(partitionIDs, partition.ID)
	}

	return svc.GetPartitionsExByIDs(ctx, partitionIDs)
}

func (svc *passportExService) GetPartitionsExByIDs(ctx context.Context, ids []int64) ([]entity.PartitionEx, error) {
	var (
		result []entity.PartitionEx
	)

	partitions, err := svc.PartitionService.GetByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}

	for _, partition := range partitions {
		partitionEx, err := svc.PartitionService.GetExByID(ctx, partition.ID)
		if err != nil {
			return nil, err
		}

		result = append(result, *partitionEx)
	}

	return result, nil
}

func (svc *passportExService) CreateChapterEx(ctx context.Context, data *service.PassportCreateChaptersData) (*entity.ChapterEx, error) {
	passport, err := svc.PassportService.GetByIDAndMunicipalityID(ctx, data.PassportID, data.MunicipalityID)
	if err != nil {
		return nil, err
	}

	if passport == nil {
		return nil, errors.New("passport not found")
	}

	chapter, err := svc.ChapterService.Create(ctx, &data.ChapterData.ChapterData)
	if err != nil {
		return nil, err
	}

	return entity.NewChapterExPtr(chapter, nil), nil
}

func (svc *passportExService) UpdateChapterEx(ctx context.Context, data *service.PassportUpdateChaptersData) (*entity.ChapterEx, error) {
	var (
		chapter *entity.Chapter
	)

	passport, err := svc.PassportService.GetByIDAndMunicipalityID(ctx, data.PassportID, data.MunicipalityID)
	if err != nil {
		return nil, err
	}

	if passport == nil {
		return nil, errors.New("passport not found")
	}

	if data.ChapterData.ChapterData != nil {
		chapter, err = svc.ChapterService.Update(ctx, data.ChapterData.ChapterData)
		if err != nil {
			return nil, err
		}
	} else {
		chapter, err = svc.ChapterService.GetByIDAndPassportID(ctx, data.ChapterData.ChapterID, data.PassportID)
		if err != nil {
			return nil, err
		}
	}

	partitions, err := svc.GetPartitionsExByChapterID(ctx, chapter.ID)
	if err != nil {
		return nil, err
	}

	return entity.NewChapterExPtr(chapter, partitions), nil
}
