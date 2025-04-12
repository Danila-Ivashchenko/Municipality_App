package object_template

import (
	"context"
	"errors"
)

func (svc *objectTemplateService) DeleteByIDAndMunicipalityID(ctx context.Context, id, municipalityID int64) error {
	template, err := svc.EntityTemplateRepository.GetByIDAndMunicipalityID(ctx, id, municipalityID)
	if err != nil {
		return err
	}
	if template == nil {
		return errors.New("template does not exist")
	}

	return svc.Delete(ctx, template.ID)
}

func (svc *objectTemplateService) Delete(ctx context.Context, id int64) error {
	return svc.EntityTemplateRepository.Delete(ctx, id)
}
