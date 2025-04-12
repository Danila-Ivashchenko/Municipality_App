package entity_attribute

import (
	"context"
	"errors"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/service"
)

func (svc *entityAttributeService) CreateValues(ctx context.Context, data service.CreateEntityAttributesData) ([]entity.EntityAttributeValue, error) {
	var (
		result []entity.EntityAttributeValue
	)

	for _, valueData := range data.ValuesData {
		valueExists, err := svc.EntityAttributeValueRepository.GetByAttributeIDAndEntityID(ctx, valueData.AttributeID, data.EntityID)
		if err != nil {
			return nil, err
		}

		if valueExists != nil {
			return nil, errors.New("value already exists")
		}
	}

	for _, valueData := range data.ValuesData {
		attribute, err := svc.EntityAttributeRepository.GetByEntityTemplateIDAndID(ctx, valueData.AttributeID, data.EntityTemplateID)
		if err != nil {
			return nil, err
		}

		if attribute == nil {
			return nil, errors.New("attribute not found")
		}

		value := attribute.DefaultValue

		if valueData.Value != nil {
			value = *valueData.Value
		}

		attributeValue := &entity.EntityAttributeValue{
			EntityAttributeID: valueData.AttributeID,
			EntityID:          data.EntityID,
			Value:             value,
		}

		attributeValue, err = svc.EntityAttributeValueRepository.Create(ctx, attributeValue)
		if err != nil {
			return nil, err
		}

		result = append(result, *attributeValue)
	}

	return result, nil
}

func (svc *entityAttributeService) GetValuesByEntityID(ctx context.Context, entityID int64) ([]entity.EntityAttributeValue, error) {
	return svc.EntityAttributeValueRepository.GetByEntityID(ctx, entityID)
}

func (svc *entityAttributeService) GetAttributesExByEntityID(ctx context.Context, entityID int64) ([]entity.EntityAttributeValueEx, error) {
	var (
		ids []int64
	)

	attributeEntityValues, err := svc.GetValuesByEntityID(ctx, entityID)
	if err != nil {
		return nil, err
	}

	for _, attributeEntityValue := range attributeEntityValues {
		ids = append(ids, attributeEntityValue.ID)
	}

	return svc.GetAttributesExByIDs(ctx, ids)
}

func (svc *entityAttributeService) GetAttributesExByIDs(ctx context.Context, entityAttributeValueIDs []int64) ([]entity.EntityAttributeValueEx, error) {
	var (
		result []entity.EntityAttributeValueEx
	)

	attributeValues, err := svc.EntityAttributeValueRepository.GetByIDs(ctx, entityAttributeValueIDs)
	if err != nil {
		return nil, err
	}

	for _, attributeValue := range attributeValues {
		attribute, err := svc.EntityAttributeRepository.GetByID(ctx, attributeValue.EntityAttributeID)
		if err != nil {
			return nil, err
		}

		if attribute == nil {
			continue
		}

		result = append(result, entity.NewEntityAttributeValueEx(*attribute, attributeValue))
	}

	return result, nil
}

func (svc *entityAttributeService) UpdateValues(ctx context.Context, data service.CreateEntityAttributesData) ([]entity.EntityAttributeValue, error) {
	var (
		result                    []entity.EntityAttributeValue
		attributeValue            *entity.EntityAttributeValue
		defaultValueByAttributeID = make(map[int64]string)
		setValueByAttributeID     = make(map[int64]string)
		attributesIDs             = make(map[int64]struct{})
	)
	allAttributes, err := svc.EntityAttributeRepository.GetByEntityTemplateID(ctx, data.EntityTemplateID)
	if err != nil {
		return nil, err
	}

	for _, attribute := range allAttributes {
		defaultValueByAttributeID[attribute.ID] = attribute.DefaultValue
		attributesIDs[attribute.ID] = struct{}{}
	}

	for _, value := range data.ValuesData {
		if value.Value != nil {
			setValueByAttributeID[value.AttributeID] = *value.Value
		}
	}

	for attributeID := range attributesIDs {
		value := defaultValueByAttributeID[attributeID]

		if setValue, exists := setValueByAttributeID[attributeID]; exists {
			value = setValue
		}

		valueExists, err := svc.EntityAttributeValueRepository.GetByAttributeIDAndEntityID(ctx, attributeID, data.EntityID)
		if err != nil {
			return nil, err
		}

		if valueExists != nil {
			valueExists.Value = value

			attributeValue, err = svc.EntityAttributeValueRepository.Update(ctx, valueExists)
			if err != nil {
				return nil, err
			}
			result = append(result, *attributeValue)
		} else {
			attributeValue = &entity.EntityAttributeValue{
				EntityAttributeID: attributeID,
				EntityID:          data.EntityID,
				Value:             value,
			}

			attributeValue, err = svc.EntityAttributeValueRepository.Create(ctx, attributeValue)
			if err != nil {
				return nil, err
			}

			result = append(result, *attributeValue)
		}
	}

	return result, nil
}
