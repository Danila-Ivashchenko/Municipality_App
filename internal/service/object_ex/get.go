package object_ex

import (
	"context"
	"municipality_app/internal/domain/entity"
)

func (svc *objectExService) GetByID(ctx context.Context, id int64) (*entity.ObjectTemplateEx, error) {
	template, err := svc.ObjectTemplateService.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if template == nil {
		return nil, nil
	}

	objectsEx, err := svc.ObjectService.GetExByTemplateID(ctx, template.ID)
	if err != nil {
		return nil, err
	}

	attributes, err := svc.ObjectAttributeService.GetAttributesByObjectTemplateID(ctx, template.ID)
	if err != nil {
		return nil, err
	}

	objectType, err := svc.ObjectTypeService.GetByID(ctx, template.ObjectTypeID)
	if err != nil {
		return nil, err
	}

	templateEx := entity.NewObjectTemplateEx(*template, objectsEx, attributes, objectType)

	return templateEx, nil
}

func (svc *objectExService) GetByMunicipalityID(ctx context.Context, id int64) ([]entity.ObjectTemplateEx, error) {
	var (
		result []entity.ObjectTemplateEx
	)

	templates, err := svc.ObjectTemplateService.GetByMunicipalityID(ctx, id)
	if err != nil {
		return nil, err
	}

	for _, template := range templates {
		objectsEx, err := svc.ObjectService.GetExByTemplateID(ctx, template.ID)
		if err != nil {
			return nil, err
		}

		attributes, err := svc.ObjectAttributeService.GetAttributesByObjectTemplateID(ctx, template.ID)
		if err != nil {
			return nil, err
		}

		objectType, err := svc.ObjectTypeService.GetByID(ctx, template.ObjectTypeID)
		if err != nil {
			return nil, err
		}

		templateEx := entity.NewObjectTemplateEx(template, objectsEx, attributes, objectType)

		result = append(result, *templateEx)
	}

	return result, nil
}
