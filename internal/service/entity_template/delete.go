package entity_template

import (
	"context"
	"municipality_app/internal/domain/core_errors"
)

func (svc *objectTemplateService) DeleteByIDAndMunicipalityID(ctx context.Context, id, municipalityID int64) error {
	template, err := svc.EntityTemplateRepository.GetByIDAndMunicipalityID(ctx, id, municipalityID)
	if err != nil {
		return err
	}
	if template == nil {
		return core_errors.EntityTemplateNotFound
	}

	return svc.Delete(ctx, template.ID)
}

func (svc *objectTemplateService) Delete(ctx context.Context, id int64) error {
	return svc.EntityTemplateRepository.Delete(ctx, id)
}
