package object_ex

import (
	"go.uber.org/fx"
	"municipality_app/internal/domain/service"
)

type ServiceParams struct {
	fx.In

	ObjectService          service.ObjectService
	ObjectTemplateService  service.ObjectTemplateService
	ObjectAttributeService service.ObjectAttributeService
}

type objectExService struct {
	ServiceParams
}

func New(params ServiceParams) service.ObjectExService {
	return &objectExService{
		ServiceParams: params,
	}
}
