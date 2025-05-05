package entity_ex

import (
	"context"
	"municipality_app/internal/domain/entity"
)

func (svc *entityExService) GetByID(ctx context.Context, id int64) (*entity.EntityTemplateEx, error) {
	template, err := svc.EntityTemplateService.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if template == nil {
		return nil, nil
	}

	entitiesEx, err := svc.EntityService.GetExByTemplateID(ctx, template.ID)
	if err != nil {
		return nil, err
	}

	attributes, err := svc.EntityAttributeService.GetAttributesByEntityTemplateID(ctx, template.ID)
	if err != nil {
		return nil, err
	}

	entityType, err := svc.EntityTypeService.GetByID(ctx, template.EntityTypeID)
	if err != nil {
		return nil, err
	}

	templateEx := entity.NewEntityTemplateEx(*template, entitiesEx, attributes, entityType)

	return templateEx, nil
}

func (svc *entityExService) GetByMunicipalityID(ctx context.Context, id int64) ([]entity.EntityTemplateEx, error) {
	var (
		result []entity.EntityTemplateEx
	)

	templates, err := svc.EntityTemplateService.GetByMunicipalityID(ctx, id)
	if err != nil {
		return nil, err
	}

	for _, template := range templates {
		objectsEx, err := svc.EntityService.GetExByTemplateID(ctx, template.ID)
		if err != nil {
			return nil, err
		}

		attributes, err := svc.EntityAttributeService.GetAttributesByEntityTemplateID(ctx, template.ID)
		if err != nil {
			return nil, err
		}

		entityType, err := svc.EntityTypeService.GetByID(ctx, template.EntityTypeID)
		if err != nil {
			return nil, err
		}

		templateEx := entity.NewEntityTemplateEx(template, objectsEx, attributes, entityType)

		result = append(result, *templateEx)
	}

	return result, nil
}
