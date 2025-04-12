package entity

import (
	"context"
	"fmt"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/service"
)

func (svc *entityService) UpdateMultiply(ctx context.Context, data *service.UpdateMultiplyEntitiesData) ([]entity.EntityEx, error) {
	var (
		uniqueNames = make(map[string]struct{})
		names       []string

		result []entity.EntityEx
	)

	for _, entityData := range data.Entities {
		if entityData.Name != nil {
			_, ok := uniqueNames[*entityData.Name]
			if ok {
				return nil, fmt.Errorf("duplicate entity name: %s", entityData.Name)
			}

			uniqueNames[*entityData.Name] = struct{}{}
			names = append(names, *entityData.Name)
		}

	}

	entityExists, err := svc.GetByNamesAndTemplateID(ctx, names, data.EntityTemplateID)
	if err != nil {
		return nil, err
	}

	if len(entityExists) > 0 {
		return nil, fmt.Errorf("duplicate entity name: %s", entityExists[0].Name)
	}

	for _, entityData := range data.Entities {
		e, err := svc.EntityRepository.GetByID(ctx, entityData.ID)
		if err != nil {
			return nil, err
		}

		if e == nil {
			continue
		}

		if entityData.Name != nil {
			e.Name = *entityData.Name
		}

		if entityData.Description != nil {
			e.Description = *entityData.Description
		}

		e, err = svc.EntityRepository.Update(ctx, e)
		if err != nil {
			return nil, err
		}

		createAttributeValuesData := service.CreateEntityAttributesData{
			EntityID:         e.ID,
			EntityTemplateID: e.EntityTemplateID,
			ValuesData:       entityData.AttributeValues,
		}

		_, err = svc.EntityAttributeService.UpdateValues(ctx, createAttributeValuesData)
		if err != nil {
			return nil, err
		}

		attributeValues, err := svc.EntityAttributeService.GetAttributesExByEntityID(ctx, e.ID)
		if err != nil {
			return nil, err
		}

		result = append(result, *entity.NewEntityExPtr(e, attributeValues))
	}

	return result, nil
}
