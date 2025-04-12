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

	templateEx := entity.NewEntityTemplateEx(*template, entitiesEx, attributes)

	return templateEx, nil
}
