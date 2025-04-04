package object_template

import (
	"context"
	"errors"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/service"
)

func (svc *objectTemplateService) Update(ctx context.Context, data *service.UpdateObjectTemplateData) (*entity.ObjectTemplateEx, error) {
	var (
		attributes []entity.ObjectAttribute
	)

	templateExist, err := svc.GetByIDAndMunicipalityID(ctx, data.ID, data.MunicipalityID)
	if err != nil {
		return nil, err
	}

	if templateExist == nil {
		return nil, errors.New("template not found")
	}

	if data.Name != nil {
		templateExist.Name = *data.Name
	}

	if data.ObjectType != nil {
		templateExist.ObjectTypeID = *data.ObjectType
	}

	template, err := svc.ObjectTemplateRepository.Update(ctx, templateExist)
	if err != nil {
		return nil, err
	}

	allObjectsByTemplate, err := svc.ObjectService.GetByTemplateID(ctx, templateExist.ID)
	if err != nil {
		return nil, err
	}

	for _, attributeData := range data.AttributesToCreate {
		createAttributeData := service.CreateObjectAttributeData{
			ObjectTemplateID: template.ID,
			Name:             attributeData.Name,
			DefaultValue:     attributeData.DefaultValue,
			ToShow:           attributeData.ToShow,
		}

		attribute, err := svc.ObjectAttributeService.CreateAttribute(ctx, createAttributeData)
		if err != nil {
			return nil, err
		}

		for _, object := range allObjectsByTemplate {
			value := &entity.ObjectAttributeValue{
				ObjectID:          object.ID,
				Value:             attribute.DefaultValue,
				ObjectAttributeID: attribute.ID,
			}

			_, err = svc.ObjectAttributeValueRepo.Create(ctx, value)
			if err != nil {
				return nil, err
			}
		}

		attributes = append(attributes, *attribute)
	}

	for _, attributeData := range data.AttributesToUpdate {
		updateAttributeData := service.UpdateObjectAttributeData{
			ID:               attributeData.ID,
			ObjectTemplateID: template.ID,
			Name:             attributeData.Name,
			DefaultValue:     attributeData.DefaultValue,
			ToShow:           attributeData.ToShow,
		}

		attribute, err := svc.ObjectAttributeService.UpdateAttribute(ctx, updateAttributeData)
		if err != nil {
			return nil, err
		}

		attributes = append(attributes, *attribute)
	}

	for _, attributeID := range data.AttributesToDelete {
		attribute, err := svc.ObjectAttributeService.GetAttributeByIDAndTemplateID(ctx, attributeID, template.ID)
		if err != nil {
			return nil, err
		}

		if attribute == nil {
			return nil, errors.New("attribute not found")
		}

		err = svc.ObjectAttributeService.DeleteAttribute(ctx, attributeID)
		if err != nil {
			return nil, err
		}
	}

	return svc.GetExByID(ctx, template.ID)
}
