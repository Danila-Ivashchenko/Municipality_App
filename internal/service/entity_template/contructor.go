package object_template

import (
	"go.uber.org/fx"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/domain/service"
)

type ServiceParams struct {
	fx.In

	EntityTemplateRepository repository.EntityTemplateRepository
	EntityAttributeValueRepo repository.EntityAttributeValueRepository

	EntityAttributeService service.EntityAttributeService
	EntityService          service.EntityService
}

type objectTemplateService struct {
	ServiceParams
}

func New(params ServiceParams) service.EntityTemplateService {
	return &objectTemplateService{
		ServiceParams: params,
	}
}
