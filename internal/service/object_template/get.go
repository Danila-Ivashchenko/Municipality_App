package object_template

import (
	"context"
	"errors"
	"municipality_app/internal/domain/entity"
)

func (svc *objectTemplateService) GetByNameAndMunicipalityID(ctx context.Context, name string, municipalityID int64) (*entity.ObjectTemplate, error) {
	return svc.ObjectTemplateRepository.GetByNameAndMunicipalityID(ctx, name, municipalityID)
}

func (svc *objectTemplateService) GetByIDAndMunicipalityID(ctx context.Context, id, municipalityID int64) (*entity.ObjectTemplate, error) {
	return svc.ObjectTemplateRepository.GetByIDAndMunicipalityID(ctx, id, municipalityID)
}

func (svc *objectTemplateService) GetExByID(ctx context.Context, id int64) (*entity.ObjectTemplateEx, error) {
	template, err := svc.ObjectTemplateRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if template == nil {
		return nil, errors.New("template does not exist")
	}

	attributes, err := svc.ObjectAttributeService.GetAttributesByObjectTemplateID(ctx, template.ID)
	if err != nil {
		return nil, err
	}

	objectsEx, err := svc.ObjectService.GetExByTemplateID(ctx, template.ID)
	if err != nil {
		return nil, err
	}

	return entity.NewObjectTemplateEx(*template, objectsEx, attributes), nil
}

func (svc *objectTemplateService) GetByMunicipalityID(ctx context.Context, partitionID int64) ([]entity.ObjectTemplate, error) {
	return svc.ObjectTemplateRepository.GetByMunicipalityID(ctx, partitionID)
}

func (svc *objectTemplateService) GetExByMunicipalityID(ctx context.Context, municipalityID int64) ([]entity.ObjectTemplateEx, error) {
	var (
		result []entity.ObjectTemplateEx
	)

	templates, err := svc.ObjectTemplateRepository.GetByMunicipalityID(ctx, municipalityID)
	if err != nil {
		return nil, err
	}

	for _, template := range templates {
		templateEx, err := svc.GetExByID(ctx, template.ID)
		if err != nil {
			return nil, err
		}

		if templateEx != nil {
			result = append(result, *templateEx)
		}

	}

	return result, nil
}

func (svc *objectTemplateService) GetExByIDs(ctx context.Context, ids []int64) ([]entity.ObjectTemplateEx, error) {
	var (
		result []entity.ObjectTemplateEx
	)

	templates, err := svc.ObjectTemplateRepository.GetByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}

	for _, template := range templates {
		templateEx, err := svc.GetExByID(ctx, template.ID)
		if err != nil {
			return nil, err
		}

		if templateEx != nil {
			result = append(result, *templateEx)
		}

	}

	return result, nil
}

func (svc *objectTemplateService) GetByID(ctx context.Context, id int64) (*entity.ObjectTemplate, error) {
	return svc.ObjectTemplateRepository.GetByID(ctx, id)
}
