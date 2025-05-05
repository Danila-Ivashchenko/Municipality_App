package entity_ex

import (
	"go.uber.org/fx"
	"municipality_app/internal/domain/service"
)

type ServiceParams struct {
	fx.In

	EntityService          service.EntityService
	EntityTemplateService  service.EntityTemplateService
	EntityAttributeService service.EntityAttributeService
	EntityTypeService      service.EntityTypeService
}

type entityExService struct {
	ServiceParams
}

func New(params ServiceParams) service.EntityExService {
	return &entityExService{
		ServiceParams: params,
	}
}
