package object_template

import (
	"context"
	"errors"
	"municipality_app/internal/domain/core_errors"
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
		return nil, core_errors.ObjectTemplateNotFound
	}

	if data.Name != nil && templateExist.Name != *data.Name {
		templateWithNameExists, err := svc.GetByNameAndMunicipalityID(ctx, *data.Name, data.MunicipalityID)
		if err != nil {
			return nil, err
		}

		if templateWithNameExists != nil {
			return nil, core_errors.ObjectTemplateNameIsUsed
		}

		templateExist.Name = *data.Name
	}

	if data.ObjectType != nil {
		templateExist.ObjectTypeID = *data.ObjectType
	}

	err = svc.Transactor.Execute(ctx, func(tx context.Context) error {
		template, err := svc.ObjectTemplateRepository.Update(tx, templateExist)
		if err != nil {
			return err
		}

		allObjectsByTemplate, err := svc.ObjectService.GetByTemplateID(tx, templateExist.ID)
		if err != nil {
			return err
		}

		for _, attributeData := range data.AttributesToCreate {
			createAttributeData := service.CreateObjectAttributeData{
				ObjectTemplateID: template.ID,
				Name:             attributeData.Name,
				DefaultValue:     attributeData.DefaultValue,
				ToShow:           attributeData.ToShow,
			}

			attribute, err := svc.ObjectAttributeService.CreateAttribute(tx, createAttributeData)
			if err != nil {
				return err
			}

			for _, object := range allObjectsByTemplate {
				value := &entity.ObjectAttributeValue{
					ObjectID:          object.ID,
					Value:             attribute.DefaultValue,
					ObjectAttributeID: attribute.ID,
				}

				_, err = svc.ObjectAttributeValueRepo.Create(tx, value)
				if err != nil {
					return err
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

			attribute, err := svc.ObjectAttributeService.UpdateAttribute(tx, updateAttributeData)
			if err != nil {
				return err
			}

			attributes = append(attributes, *attribute)
		}

		for _, attributeID := range data.AttributesToDelete {
			attribute, err := svc.ObjectAttributeService.GetAttributeByIDAndTemplateID(tx, attributeID, template.ID)
			if err != nil {
				return err
			}

			if attribute == nil {
				return errors.New("attribute not found")
			}

			err = svc.ObjectAttributeService.DeleteAttribute(tx, attributeID)
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
