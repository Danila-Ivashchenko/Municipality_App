package object_attribute

import (
	"context"
	"errors"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/service"
)

func (svc *objectAttributeService) CreateValues(ctx context.Context, data service.CreateObjectAttributesData) ([]entity.ObjectAttributeValue, error) {
	var (
		result []entity.ObjectAttributeValue
	)

	for _, valueData := range data.ValuesData {
		valueExists, err := svc.ObjectAttributeValueRepository.GetByAttributeIDAndObjectID(ctx, valueData.AttributeID, data.ObjectID)
		if err != nil {
			return nil, err
		}

		if valueExists != nil {
			return nil, errors.New("value already exists")
		}
	}

	for _, valueData := range data.ValuesData {
		attribute, err := svc.ObjectAttributeRepository.GetByObjectTemplateIDAndID(ctx, valueData.AttributeID, data.ObjectTemplateID)
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

		attributeValue := &entity.ObjectAttributeValue{
			ObjectAttributeID: valueData.AttributeID,
			ObjectID:          data.ObjectID,
			Value:             value,
		}

		attributeValue, err = svc.ObjectAttributeValueRepository.Create(ctx, attributeValue)
		if err != nil {
			return nil, err
		}

		result = append(result, *attributeValue)
	}

	return result, nil
}

func (svc *objectAttributeService) GetValuesByObjectID(ctx context.Context, objectID int64) ([]entity.ObjectAttributeValue, error) {
	return svc.ObjectAttributeValueRepository.GetByObjectID(ctx, objectID)
}

func (svc *objectAttributeService) GetAttributesExByObjectID(ctx context.Context, objectID int64) ([]entity.ObjectAttributeValueEx, error) {
	var (
		ids []int64
	)

	attributeObjectValues, err := svc.GetValuesByObjectID(ctx, objectID)
	if err != nil {
		return nil, err
	}

	for _, attributeObjectValue := range attributeObjectValues {
		ids = append(ids, attributeObjectValue.ID)
	}

	return svc.GetAttributesExByIDs(ctx, ids)
}

func (svc *objectAttributeService) GetAttributesExByIDs(ctx context.Context, objectAttributeValueIDs []int64) ([]entity.ObjectAttributeValueEx, error) {
	var (
		result []entity.ObjectAttributeValueEx
	)

	attributeValues, err := svc.ObjectAttributeValueRepository.GetByIDs(ctx, objectAttributeValueIDs)
	if err != nil {
		return nil, err
	}

	for _, attributeValue := range attributeValues {
		attribute, err := svc.ObjectAttributeRepository.GetByID(ctx, attributeValue.ObjectAttributeID)
		if err != nil {
			return nil, err
		}

		if attribute == nil {
			continue
		}

		result = append(result, entity.NewObjectAttributeValueEx(*attribute, attributeValue))
	}

	return result, nil
}

func (svc *objectAttributeService) UpdateValues(ctx context.Context, data service.CreateObjectAttributesData) ([]entity.ObjectAttributeValue, error) {
	var (
		result                    []entity.ObjectAttributeValue
		attributeValue            *entity.ObjectAttributeValue
		defaultValueByAttributeID = make(map[int64]string)
		setValueByAttributeID     = make(map[int64]string)
		attributesIDs             = make(map[int64]struct{})
	)
	allAttributes, err := svc.ObjectAttributeRepository.GetByObjectTemplateID(ctx, data.ObjectTemplateID)
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

		valueExists, err := svc.ObjectAttributeValueRepository.GetByAttributeIDAndObjectID(ctx, attributeID, data.ObjectID)
		if err != nil {
			return nil, err
		}

		if valueExists != nil {
			valueExists.Value = value

			attributeValue, err = svc.ObjectAttributeValueRepository.Update(ctx, valueExists)
			if err != nil {
				return nil, err
			}
			result = append(result, *attributeValue)
		} else {
			attributeValue = &entity.ObjectAttributeValue{
				ObjectAttributeID: attributeID,
				ObjectID:          data.ObjectID,
				Value:             value,
			}

			attributeValue, err = svc.ObjectAttributeValueRepository.Create(ctx, attributeValue)
			if err != nil {
				return nil, err
			}

			result = append(result, *attributeValue)
		}
	}

	return result, nil
}
