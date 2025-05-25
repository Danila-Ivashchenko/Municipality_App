package entity

import (
	"context"
	"municipality_app/internal/domain/core_errors"
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
				return nil, core_errors.EntityNameIsUsed
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
		return nil, core_errors.EntityNameIsUsed
	}

	err = svc.Transactor.Execute(ctx, func(tx context.Context) error {
		for _, entityData := range data.Entities {
			e, err := svc.EntityRepository.GetByID(tx, entityData.ID)
			if err != nil {
				return err
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

			e, err = svc.EntityRepository.Update(tx, e)
			if err != nil {
				return err
			}

			createAttributeValuesData := service.CreateEntityAttributesData{
				EntityID:         e.ID,
				EntityTemplateID: e.EntityTemplateID,
				ValuesData:       entityData.AttributeValues,
			}

			_, err = svc.EntityAttributeService.UpdateValues(tx, createAttributeValuesData)
			if err != nil {
				return err
			}

			attributeValues, err := svc.EntityAttributeService.GetAttributesExByEntityID(tx, e.ID)
			if err != nil {
				return err
			}

			result = append(result, *entity.NewEntityExPtr(e, attributeValues))
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}
