package municipality

import (
	"context"
	"errors"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/domain/service"
)

func (svc *municipalityService) Create(ctx context.Context, data *service.CreateMunicipalityData) (*entity.MunicipalityEx, error) {
	region, err := svc.RegionRepository.GetById(ctx, data.RegionID)
	if err != nil {
		return nil, err
	}

	if region == nil {
		return nil, errors.New("region not found")
	}

	municipalityExists, err := svc.GetByName(ctx, data.Name)
	if err != nil {
		return nil, err
	}

	if municipalityExists != nil {
		return nil, errors.New("mun already exists")
	}

	repoData := &repository.CreateMunicipalityData{
		Name:     data.Name,
		RegionID: data.RegionID,
	}

	municipality, err := svc.MunicipalityRepository.Create(ctx, repoData)
	if err != nil {
		return nil, err
	}

	return entity.NewMunicipalityEx(municipality, region), nil

}

func (svc *municipalityService) Update(ctx context.Context, data *service.UpdateMunicipalityData) (*entity.MunicipalityEx, error) {
	municipalityExists, err := svc.GetById(ctx, data.ID)
	if err != nil {
		return nil, err
	}

	if municipalityExists == nil {
		return nil, errors.New("mun not found")
	}

	if data.Name != nil {
		municipalityExists.Name = *data.Name
	}

	if data.RegionID != nil {
		regionExists, err := svc.RegionRepository.GetById(ctx, *data.RegionID)
		if err != nil {
			return nil, err
		}

		if regionExists == nil {
			return nil, errors.New("region not found")
		}

		municipalityExists.RegionID = *data.RegionID
	}

	if data.IsHidden != nil {
		municipalityExists.IsHidden = *data.IsHidden
	}

	err = svc.MunicipalityRepository.Update(ctx, municipalityExists)
	if err != nil {
		return nil, err
	}

	municipalityRegion, err := svc.RegionRepository.GetById(ctx, municipalityExists.RegionID)
	if err != nil {
		return nil, err
	}

	if municipalityRegion == nil {
		return nil, errors.New("mun not found")
	}

	return entity.NewMunicipalityEx(municipalityExists, municipalityRegion), nil
}

func (svc *municipalityService) Delete(ctx context.Context, id int64) error {
	municipalityExists, err := svc.GetById(ctx, id)
	if err != nil {
		return err
	}

	if municipalityExists == nil {
		return errors.New("mun not found")
	}

	return svc.MunicipalityRepository.Delete(ctx, id)
}

func (svc *municipalityService) GetById(ctx context.Context, id int64) (*entity.Municipality, error) {
	return svc.MunicipalityRepository.GetById(ctx, id)
}

func (svc *municipalityService) GetExById(ctx context.Context, id int64) (*entity.MunicipalityEx, error) {
	municipality, err := svc.MunicipalityRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	if municipality == nil {
		return nil, errors.New("mun not found")
	}

	region, err := svc.RegionRepository.GetById(ctx, municipality.RegionID)
	if err != nil {
		return nil, err
	}

	if region == nil {
		return nil, errors.New("region not found")
	}

	return entity.NewMunicipalityEx(municipality, region), nil
}

func (svc *municipalityService) GetByName(ctx context.Context, name string) (*entity.Municipality, error) {
	return svc.MunicipalityRepository.GetByName(ctx, name)
}

func (svc *municipalityService) GetAll(ctx context.Context) ([]entity.Municipality, error) {
	return svc.MunicipalityRepository.GetAll(ctx)
}

func (svc *municipalityService) GetByParams(ctx context.Context, params *service.GetMunicipalityParams) ([]entity.Municipality, error) {
	repoParams := &repository.MunicipalityParams{
		ID:       params.ID,
		Name:     params.Name,
		RegionID: params.RegionID,
		IsHidden: params.IsHidden,
	}

	return svc.MunicipalityRepository.GetByParams(ctx, repoParams)
}

func (svc *municipalityService) GetExByParams(ctx context.Context, params *service.GetMunicipalityParams) ([]entity.MunicipalityEx, error) {
	var (
		regionIDsMap = make(map[int64]struct{})
		regionIDs    []int64
		regionByID   = make(map[int64]entity.Region)
		result       = make([]entity.MunicipalityEx, 0)
	)

	municipalities, err := svc.GetByParams(ctx, params)
	if err != nil {
		return nil, err
	}

	if len(municipalities) == 0 {
		return result, nil
	}

	for _, municipality := range municipalities {
		regionIDsMap[municipality.RegionID] = struct{}{}
	}

	for regionID := range regionIDsMap {
		regionIDs = append(regionIDs, regionID)
	}

	regions, err := svc.RegionRepository.GetByIds(ctx, regionIDs)
	if err != nil {
		return nil, err
	}

	for _, region := range regions {
		regionByID[region.ID] = region
	}

	for _, municipality := range municipalities {
		region, ok := regionByID[municipality.RegionID]
		if !ok {
			continue
		}

		result = append(result, *entity.NewMunicipalityEx(&municipality, &region))
	}

	return result, nil
}
