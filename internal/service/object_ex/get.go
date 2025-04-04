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

	templateEx := entity.NewObjectTemplateEx(*template, objectsEx, attributes)

	return templateEx, nil
}
