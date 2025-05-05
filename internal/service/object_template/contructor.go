package object_template

import (
	"go.uber.org/fx"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/domain/service"
)

type ServiceParams struct {
	fx.In

	ObjectTemplateRepository repository.ObjectTemplateRepository
	ObjectAttributeValueRepo repository.ObjectAttributeValueRepository

	ObjectAttributeService service.ObjectAttributeService
	ObjectService          service.ObjectService
	ObjectTypeService      service.ObjectTypeService
}

type objectTemplateService struct {
	ServiceParams
}

func New(params ServiceParams) service.ObjectTemplateService {
	return &objectTemplateService{
		ServiceParams: params,
	}
}
