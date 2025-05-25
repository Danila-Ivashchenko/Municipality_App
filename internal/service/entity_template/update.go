package entity_template

import (
	"context"
	"errors"
	"municipality_app/internal/domain/core_errors"
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
		return nil, core_errors.EntityTemplateNotFound
	}

	if data.Name != nil && templateExist.Name != *data.Name {
		templateWithNameExists, err := svc.GetByNameAndMunicipalityID(ctx, *data.Name, data.MunicipalityID)
		if err != nil {
			return nil, err
		}

		if templateWithNameExists != nil {
			return nil, core_errors.EntityTemplateNameIsUsed
		}

		templateExist.Name = *data.Name
	}

	if data.EntityType != nil {
		templateExist.EntityTypeID = *data.EntityType
	}

	err = svc.Transactor.Execute(ctx, func(tx context.Context) error {
		template, err := svc.EntityTemplateRepository.Update(tx, templateExist)
		if err != nil {
			return err
		}

		allEntitiesByTemplate, err := svc.EntityService.GetByTemplateID(tx, templateExist.ID)
		if err != nil {
			return err
		}

		for _, attributeData := range data.AttributesToCreate {
			createAttributeData := service.CreateEntityAttributeData{
				EntityTemplateID: template.ID,
				Name:             attributeData.Name,
				DefaultValue:     attributeData.DefaultValue,
				ToShow:           attributeData.ToShow,
			}

			attribute, err := svc.EntityAttributeService.CreateAttribute(tx, createAttributeData)
			if err != nil {
				return err
			}

			for _, object := range allEntitiesByTemplate {
				value := &entity.EntityAttributeValue{
					EntityID:          object.ID,
					Value:             attribute.DefaultValue,
					EntityAttributeID: attribute.ID,
				}

				_, err = svc.EntityAttributeValueRepo.Create(tx, value)
				if err != nil {
					return err
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

			attribute, err := svc.EntityAttributeService.UpdateAttribute(tx, updateAttributeData)
			if err != nil {
				return err
			}

			attributes = append(attributes, *attribute)
		}

		for _, attributeID := range data.AttributesToDelete {
			attribute, err := svc.EntityAttributeService.GetAttributeByIDAndTemplateID(tx, attributeID, template.ID)
			if err != nil {
				return err
			}

			if attribute == nil {
				return errors.New("attribute not found")
			}

			err = svc.EntityAttributeService.DeleteAttribute(tx, attributeID)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return svc.GetExByID(ctx, data.ID)
}
