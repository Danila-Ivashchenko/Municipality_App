package object_attribute

import (
	"context"
	"errors"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/service"
)

func (svc *objectAttributeService) CreateAttribute(ctx context.Context, data service.CreateObjectAttributeData) (*entity.ObjectAttribute, error) {
	attributeExists, err := svc.ObjectAttributeRepository.GetByObjectTemplateIDAndName(ctx, data.Name, data.ObjectTemplateID)
	if err != nil {
		return nil, err
	}

	if attributeExists != nil {
		return nil, errors.New("object attribute exists")
	}

	attribute := &entity.ObjectAttribute{
		ObjectTemplateID: data.ObjectTemplateID,
		Name:             data.Name,
		DefaultValue:     data.DefaultValue,
		ToShow:           data.ToShow,
	}

	return svc.ObjectAttributeRepository.Create(ctx, attribute)
}

func (svc *objectAttributeService) DeleteAttribute(ctx context.Context, id int64) error {
	return svc.ObjectAttributeRepository.Delete(ctx, id)
}

func (svc *objectAttributeService) UpdateAttribute(ctx context.Context, data service.UpdateObjectAttributeData) (*entity.ObjectAttribute, error) {
	attributeExists, err := svc.ObjectAttributeRepository.GetByObjectTemplateIDAndID(ctx, data.ID, data.ObjectTemplateID)
	if err != nil {
		return nil, err
	}

	if attributeExists == nil {
		return nil, errors.New("object attribute not found")
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

	return svc.ObjectAttributeRepository.Update(ctx, attributeExists)
}

func (svc *objectAttributeService) GetAttributesByObjectTemplateID(ctx context.Context, templateID int64) ([]entity.ObjectAttribute, error) {
	return svc.ObjectAttributeRepository.GetByObjectTemplateID(ctx, templateID)
}

func (svc *objectAttributeService) GetAttributeByIDAndTemplateID(ctx context.Context, attributeID, templateID int64) (*entity.ObjectAttribute, error) {
	return svc.ObjectAttributeRepository.GetByObjectTemplateIDAndID(ctx, attributeID, templateID)
}
