package entity_attribute

import (
	"context"
	"errors"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/service"
)

func (svc *entityAttributeService) CreateAttribute(ctx context.Context, data service.CreateEntityAttributeData) (*entity.EntityAttribute, error) {
	attributeExists, err := svc.EntityAttributeRepository.GetByEntityTemplateIDAndName(ctx, data.Name, data.EntityTemplateID)
	if err != nil {
		return nil, err
	}

	if attributeExists != nil {
		return nil, errors.New("entity attribute exists")
	}

	attribute := &entity.EntityAttribute{
		EntityTemplateID: data.EntityTemplateID,
		Name:             data.Name,
		DefaultValue:     data.DefaultValue,
		ToShow:           data.ToShow,
	}

	return svc.EntityAttributeRepository.Create(ctx, attribute)
}

func (svc *entityAttributeService) DeleteAttribute(ctx context.Context, id int64) error {
	return svc.EntityAttributeRepository.Delete(ctx, id)
}

func (svc *entityAttributeService) UpdateAttribute(ctx context.Context, data service.UpdateEntityAttributeData) (*entity.EntityAttribute, error) {
	attributeExists, err := svc.EntityAttributeRepository.GetByEntityTemplateIDAndID(ctx, data.ID, data.EntityTemplateID)
	if err != nil {
		return nil, err
	}

	if attributeExists == nil {
		return nil, errors.New("entity attribute not found")
	}

	if data.Name != nil {
		attributeExists.Name = *data.Name
	}

	if data.DefaultValue != nil {
		attributeExists.DefaultValue = *data.DefaultValue
	}

	if data.ToShow != nil {
		attributeExists.ToShow = *data.ToShow
	}

	return svc.EntityAttributeRepository.Update(ctx, attributeExists)
}

func (svc *entityAttributeService) GetAttributesByEntityTemplateID(ctx context.Context, templateID int64) ([]entity.EntityAttribute, error) {
	return svc.EntityAttributeRepository.GetByEntityTemplateID(ctx, templateID)
}

func (svc *entityAttributeService) GetAttributeByIDAndTemplateID(ctx context.Context, attributeID, templateID int64) (*entity.EntityAttribute, error) {
	return svc.EntityAttributeRepository.GetByEntityTemplateIDAndID(ctx, attributeID, templateID)
}
