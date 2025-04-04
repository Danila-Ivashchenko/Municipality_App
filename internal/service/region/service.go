package region

import (
	"context"
	"errors"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/domain/service"
)

func (svc *regionService) Create(ctx context.Context, data *service.CreateRegionData) (*entity.Region, error) {
	err := svc.checkRegionExists(ctx, data.Name, data.Code)
	if err != nil {
		return nil, err
	}

	repoData := &repository.CreateRegionData{
		Name: data.Name,
		Code: data.Code,
	}

	err = svc.RegionRepository.Create(ctx, repoData)

	return svc.GetByCode(ctx, data.Code)
}

func (svc *regionService) checkRegionExists(ctx context.Context, name string, code string) error {
	regionExists, err := svc.RegionRepository.GetByCode(ctx, code)
	if err != nil {
		return err
	}

	if regionExists != nil {
		return errors.New("region with code " + code + " already exists")
	}

	regionExists, err = svc.RegionRepository.GetByName(ctx, name)

	if err != nil {
		return err
	}

	if regionExists != nil {
		return errors.New("region with name " + name + " already exists")
	}

	return nil
}

func (svc *regionService) Delete(ctx context.Context, id int64) error {
	return svc.RegionRepository.Delete(ctx, id)
}

func (svc *regionService) GetAll(ctx context.Context) ([]entity.Region, error) {
	return svc.RegionRepository.GetAll(ctx)
}

func (svc *regionService) GetByParams(ctx context.Context, params *service.GetRegionParams) ([]entity.Region, error) {
	repoParams := &repository.RegionParams{
		ID:   params.ID,
		Code: params.Code,
		Name: params.Name,
	}

	return svc.RegionRepository.GetByParams(ctx, repoParams)
}

func (svc *regionService) GetById(ctx context.Context, id int64) (*entity.Region, error) {
	return svc.RegionRepository.GetById(ctx, id)
}

func (svc *regionService) GetByName(ctx context.Context, name string) (*entity.Region, error) {
	return svc.RegionRepository.GetByName(ctx, name)
}

func (svc *regionService) GetByCode(ctx context.Context, code string) (*entity.Region, error) {
	return svc.RegionRepository.GetByCode(ctx, code)
}
