package entity_template

import (
	"context"
	"errors"
	"municipality_app/internal/domain/entity"
)

func (svc *objectTemplateService) GetByNameAndMunicipalityID(ctx context.Context, name string, municipalityID int64) (*entity.EntityTemplate, error) {
	return svc.EntityTemplateRepository.GetByNameAndMunicipalityID(ctx, name, municipalityID)
}

func (svc *objectTemplateService) GetByIDAndMunicipalityID(ctx context.Context, id, municipalityID int64) (*entity.EntityTemplate, error) {
	return svc.EntityTemplateRepository.GetByIDAndMunicipalityID(ctx, id, municipalityID)
}

func (svc *objectTemplateService) GetExByID(ctx context.Context, id int64) (*entity.EntityTemplateEx, error) {
	template, err := svc.EntityTemplateRepository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if template == nil {
		return nil, errors.New("template does not exist")
	}

	attributes, err := svc.EntityAttributeService.GetAttributesByEntityTemplateID(ctx, template.ID)
	if err != nil {
		return nil, err
	}

	objectsEx, err := svc.EntityService.GetExByTemplateID(ctx, template.ID)
	if err != nil {
		return nil, err
	}

	entityTyper, err := svc.EntityTypeService.GetByID(ctx, template.EntityTypeID)
	if err != nil {
		return nil, err
	}

	return entity.NewEntityTemplateEx(*template, objectsEx, attributes, entityTyper), nil
}

func (svc *objectTemplateService) GetByMunicipalityID(ctx context.Context, partitionID int64) ([]entity.EntityTemplate, error) {
	return svc.EntityTemplateRepository.GetByMunicipalityID(ctx, partitionID)
}

func (svc *objectTemplateService) GetExByMunicipalityID(ctx context.Context, municipalityID int64) ([]entity.EntityTemplateEx, error) {
	var (
		result []entity.EntityTemplateEx
	)

	templates, err := svc.EntityTemplateRepository.GetByMunicipalityID(ctx, municipalityID)
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

func (svc *objectTemplateService) GetExByIDs(ctx context.Context, ids []int64) ([]entity.EntityTemplateEx, error) {
	var (
		result []entity.EntityTemplateEx
	)

	templates, err := svc.EntityTemplateRepository.GetByIDs(ctx, ids)
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

func (svc *objectTemplateService) GetByID(ctx context.Context, id int64) (*entity.EntityTemplate, error) {
	return svc.EntityTemplateRepository.GetByID(ctx, id)
}
