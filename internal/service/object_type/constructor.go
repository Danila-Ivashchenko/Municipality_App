package object_type

import (
	"go.uber.org/fx"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/domain/service"
)

type ServiceParams struct {
	fx.In

	ObjectTypeRepository     repository.ObjectTypeRepository
	ObjectTemplateRepository repository.ObjectTemplateRepository
}

type objectTypeService struct {
	ServiceParams
}

func New(params ServiceParams) service.ObjectTypeService {
	return &objectTypeService{
		ServiceParams: params,
	}
}
