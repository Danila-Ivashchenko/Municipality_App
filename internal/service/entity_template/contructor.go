package entity_template

import (
	"go.uber.org/fx"
	"municipality_app/internal/common/transactor"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/domain/service"
)

type ServiceParams struct {
	fx.In

	EntityTemplateRepository repository.EntityTemplateRepository
	EntityAttributeValueRepo repository.EntityAttributeValueRepository

	EntityAttributeService service.EntityAttributeService
	EntityService          service.EntityService
	EntityTypeService      service.EntityTypeService

	Transactor transactor.Transactor
}

type objectTemplateService struct {
	ServiceParams
}

func New(params ServiceParams) service.EntityTemplateService {
	return &objectTemplateService{
		ServiceParams: params,
	}
}
