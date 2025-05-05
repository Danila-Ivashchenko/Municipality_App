package entity_template

import (
	"context"
	"errors"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/service"
)

func (svc *objectTemplateService) Update(ctx context.Context, data *service.UpdateEntityTemplateData) (*entity.EntityTemplateEx, error) {
	var (
		attributes []entity.EntityAttribute
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

	if data.EntityType != nil {
		templateExist.EntityTypeID = *data.EntityType
	}

	template, err := svc.EntityTemplateRepository.Update(ctx, templateExist)
	if err != nil {
		return nil, err
	}

	allEntitiesByTemplate, err := svc.EntityService.GetByTemplateID(ctx, templateExist.ID)
	if err != nil {
		return nil, err
	}

	for _, attributeData := range data.AttributesToCreate {
		createAttributeData := service.CreateEntityAttributeData{
			EntityTemplateID: template.ID,
			Name:             attributeData.Name,
			DefaultValue:     attributeData.DefaultValue,
			ToShow:           attributeData.ToShow,
		}

		attribute, err := svc.EntityAttributeService.CreateAttribute(ctx, createAttributeData)
		if err != nil {
			return nil, err
		}

		for _, object := range allEntitiesByTemplate {
			value := &entity.EntityAttributeValue{
				EntityID:          object.ID,
				Value:             attribute.DefaultValue,
				EntityAttributeID: attribute.ID,
			}

			_, err = svc.EntityAttributeValueRepo.Create(ctx, value)
			if err != nil {
				return nil, err
			}
		}

		attributes = append(attributes, *attribute)
	}

	for _, attributeData := range data.AttributesToUpdate {
		updateAttributeData := service.UpdateEntityAttributeData{
			ID:               attributeData.ID,
			EntityTemplateID: template.ID,
			Name:             attributeData.Name,
			DefaultValue:     attributeData.DefaultValue,
			ToShow:           attributeData.ToShow,
		}

		attribute, err := svc.EntityAttributeService.UpdateAttribute(ctx, updateAttributeData)
		if err != nil {
			return nil, err
		}

		attributes = append(attributes, *attribute)
	}

	for _, attributeID := range data.AttributesToDelete {
		attribute, err := svc.EntityAttributeService.GetAttributeByIDAndTemplateID(ctx, attributeID, template.ID)
		if err != nil {
			return nil, err
		}

		if attribute == nil {
			return nil, errors.New("attribute not found")
		}

		err = svc.EntityAttributeService.DeleteAttribute(ctx, attributeID)
		if err != nil {
			return nil, err
		}
	}

	return svc.GetExByID(ctx, template.ID)
}
